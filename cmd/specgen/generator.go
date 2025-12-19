// Package specgen generates PLATEAU specification documents as Markdown files.
package specgen

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/eukarya-inc/plateau-spec/cmd/metanorma2md"
)

// MaxConcurrentRequests limits concurrent HTTP requests
const MaxConcurrentRequests = 10

// Generator generates specification documents
type Generator struct {
	client     *PlateauClient
	outputDir  string
	docType    string
	pathFilter []string // if non-empty, only generate pages matching these IDs
}

// NewGenerator creates a new Generator
func NewGenerator(docType, outputDir string, pathFilter []string) *Generator {
	return &Generator{
		client:     NewClient(docType),
		outputDir:  outputDir,
		docType:    docType,
		pathFilter: pathFilter,
	}
}

// Generate generates all specification documents
func (g *Generator) Generate() error {
	// Create output directories
	if err := g.createDirectories(); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Get full outline (deep)
	fmt.Println("Fetching outline...")
	outline, err := g.client.GetOutlineWithDepth(10)
	if err != nil {
		return fmt.Errorf("failed to get outline: %w", err)
	}

	// Generate index.json (always generate full index)
	fmt.Println("Generating index.json...")
	if err := g.generateIndex(outline); err != nil {
		return fmt.Errorf("failed to generate index.json: %w", err)
	}

	// Filter outline if pathFilter is specified
	if len(g.pathFilter) > 0 {
		outline = g.filterOutline(outline)
		if len(outline) == 0 {
			return fmt.Errorf("no pages matched the filter: %v", g.pathFilter)
		}
	}

	// Count total pages
	totalPages := g.countOutlineItems(outline)
	fmt.Printf("Found %d pages to generate\n", totalPages)

	// Generate pages hierarchically
	var (
		mu         sync.Mutex
		completed  int
		errorsList []string
	)

	progress := &progressTracker{
		total:     totalPages,
		completed: &completed,
		errors:    &errorsList,
		mu:        &mu,
		sem:       make(chan struct{}, MaxConcurrentRequests),
	}

	// Process top-level chapters in parallel
	var wg sync.WaitGroup
	for _, item := range outline {
		wg.Add(1)
		go func(item OutlineItem) {
			defer wg.Done()
			g.generatePageHierarchical(item, progress)
		}(item)
	}
	wg.Wait()

	if len(errorsList) > 0 {
		fmt.Println("\nErrors occurred:")
		for _, e := range errorsList {
			fmt.Printf("  - %s\n", e)
		}
	}

	fmt.Printf("\nGeneration complete! %d/%d pages generated successfully.\n", completed-len(errorsList), totalPages)
	return nil
}

// progressTracker tracks generation progress
type progressTracker struct {
	total     int
	completed *int
	errors    *[]string
	mu        *sync.Mutex
	sem       chan struct{} // semaphore to limit concurrent HTTP requests
}

func (p *progressTracker) increment() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	*p.completed++
	return *p.completed
}

func (p *progressTracker) addError(err string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	*p.errors = append(*p.errors, err)
}

// countOutlineItems counts total items in outline recursively
func (g *Generator) countOutlineItems(items []OutlineItem) int {
	count := 0
	for _, item := range items {
		count++
		count += g.countOutlineItems(item.Children)
	}
	return count
}

// pageResult holds the result of generating a page
type pageResult struct {
	content []metanorma2md.Content
	err     error
}

// generatePageHierarchical generates a page and its children hierarchically
func (g *Generator) generatePageHierarchical(item OutlineItem, progress *progressTracker) *pageResult {
	itemID := item.ID
	prefix := fmt.Sprintf("[%s]", itemID)

	// First, spawn goroutines for all children
	childResults := make([]*pageResult, len(item.Children))
	var wg sync.WaitGroup

	for i, child := range item.Children {
		wg.Add(1)
		go func(idx int, childItem OutlineItem) {
			defer wg.Done()
			childResults[idx] = g.generatePageHierarchical(childItem, progress)
		}(i, child)
	}

	// Acquire semaphore for HTTP request
	progress.sem <- struct{}{}
	fmt.Printf("%s Fetching content from %s\n", prefix, item.Path)
	doc, err := g.client.GetContentByPath(item.Path)
	<-progress.sem // Release semaphore

	if err != nil {
		progress.addError(fmt.Sprintf("%s: %v", itemID, err))
		fmt.Printf("%s ERROR fetching: %v\n", prefix, err)
		wg.Wait()
		return &pageResult{err: err}
	}
	fmt.Printf("%s Downloaded content (%d blocks)\n", prefix, len(doc.Content))

	// Wait for all children to complete
	wg.Wait()

	// Convert this page's content
	mdDoc := g.convertToMetanormaDoc(doc)

	// Merge children's content
	for _, childResult := range childResults {
		if childResult != nil && childResult.err == nil {
			mdDoc.Content = append(mdDoc.Content, childResult.content...)
		}
	}

	// Convert to Markdown
	fmt.Printf("%s Converting to Markdown (%d total blocks)\n", prefix, len(mdDoc.Content))
	opts := &metanorma2md.Options{IncludeImages: true}
	markdown := metanorma2md.Convert(mdDoc, opts)
	fmt.Printf("%s Markdown generated (%d bytes)\n", prefix, len(markdown))

	// Extract and save images
	imagesDir := filepath.Join(g.outputDir, "images")
	imageCount := countBase64Images(markdown)
	if imageCount > 0 {
		fmt.Printf("%s Extracting %d images\n", prefix, imageCount)
	}
	markdown, err = ExtractAndSaveImages(markdown, itemID, imagesDir)
	if err != nil {
		progress.addError(fmt.Sprintf("%s: %v", itemID, err))
		fmt.Printf("%s ERROR processing images: %v\n", prefix, err)
		return &pageResult{content: mdDoc.Content, err: err}
	}

	// Save JSON
	jsonPath := filepath.Join(g.outputDir, "json", itemID+".json")
	fmt.Printf("%s Saving JSON to %s\n", prefix, jsonPath)
	if err := g.saveJSON(jsonPath, doc); err != nil {
		progress.addError(fmt.Sprintf("%s: %v", itemID, err))
		fmt.Printf("%s ERROR saving JSON: %v\n", prefix, err)
	}

	// Save Markdown
	mdPath := filepath.Join(g.outputDir, itemID+".md")
	fmt.Printf("%s Saving Markdown to %s\n", prefix, mdPath)
	if err := os.WriteFile(mdPath, []byte(markdown), 0644); err != nil {
		progress.addError(fmt.Sprintf("%s: %v", itemID, err))
		fmt.Printf("%s ERROR saving Markdown: %v\n", prefix, err)
		return &pageResult{content: mdDoc.Content, err: err}
	}

	completed := progress.increment()
	fmt.Printf("%s Done! (%d/%d)\n", prefix, completed, progress.total)

	return &pageResult{content: mdDoc.Content}
}

// createDirectories creates the output directory structure
func (g *Generator) createDirectories() error {
	dirs := []string{
		g.outputDir,
		filepath.Join(g.outputDir, "json"),
		filepath.Join(g.outputDir, "images"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}

// countBase64Images counts base64 images in markdown
func countBase64Images(markdown string) int {
	count := 0
	for i := 0; i < len(markdown); i++ {
		if i+20 < len(markdown) && markdown[i:i+12] == "data:image/" {
			count++
		}
	}
	return count
}

// convertToMetanormaDoc converts PlateauDocument to metanorma2md.Document
func (g *Generator) convertToMetanormaDoc(doc *PlateauDocument) *metanorma2md.Document {
	if doc == nil {
		return nil
	}

	mdDoc := &metanorma2md.Document{
		Title:    doc.Title,
		Path:     doc.Path,
		Metadata: doc.Metadata,
	}

	for _, c := range doc.Content {
		mdDoc.Content = append(mdDoc.Content, metanorma2md.Content{
			Type:    c.Type,
			Content: c.Content,
		})
	}

	return mdDoc
}

// saveJSON saves data as JSON file
func (g *Generator) saveJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// generateIndex generates index.json and index.md
func (g *Generator) generateIndex(outline []OutlineItem) error {
	index := &Index{
		DocumentType: g.docType,
		Title:        g.getDocumentTitle(),
		GeneratedAt:  time.Now().UTC().Format(time.RFC3339),
		Chapters:     g.convertOutlineToChapters(outline),
	}

	// Save index.json
	indexJSONPath := filepath.Join(g.outputDir, "index.json")
	if err := g.saveJSON(indexJSONPath, index); err != nil {
		return err
	}

	// Generate index.md with links
	indexMDPath := filepath.Join(g.outputDir, "index.md")
	return g.saveIndexMarkdown(indexMDPath, index)
}

// saveIndexMarkdown generates index.md with links to all pages
func (g *Generator) saveIndexMarkdown(path string, index *Index) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s\n\n", index.Title))

	// Add source URL
	sourceURL := "https://www.mlit.go.jp/plateaudocument/"
	if g.docType == "procedure" {
		sourceURL = "https://www.mlit.go.jp/plateaudocument02/"
	}
	sb.WriteString(fmt.Sprintf("出典: %s\n\n", sourceURL))

	// Add disclaimer
	sb.WriteString("> **注意**: このドキュメントは上記出典から自動変換されたものです。変換処理により原本と異なる場合があります。正確な情報は原本を参照してください。\n\n")

	sb.WriteString(fmt.Sprintf("生成日時: %s\n\n", index.GeneratedAt))
	sb.WriteString("## 目次\n\n")

	g.writeChapterLinks(&sb, index.Chapters, 0)

	return os.WriteFile(path, []byte(sb.String()), 0644)
}

// writeChapterLinks writes chapter links recursively with indentation
func (g *Generator) writeChapterLinks(sb *strings.Builder, chapters []Chapter, depth int) {
	indent := strings.Repeat("  ", depth)
	for _, ch := range chapters {
		_, _ = fmt.Fprintf(sb, "%s- [%s](%s.md)\n", indent, ch.Title, ch.Path)
		if len(ch.Children) > 0 {
			g.writeChapterLinks(sb, ch.Children, depth+1)
		}
	}
}

// getDocumentTitle returns the document title based on type
func (g *Generator) getDocumentTitle() string {
	if g.docType == "procedure" {
		return "3D都市モデル標準作業手順書"
	}
	return "3D都市モデル標準製品仕様書"
}

// convertOutlineToChapters converts OutlineItems to Chapter format
func (g *Generator) convertOutlineToChapters(items []OutlineItem) []Chapter {
	var chapters []Chapter
	for _, item := range items {
		// Extract just the filename part from path
		pathID := extractID(item.Path)
		chapter := Chapter{
			Path:     pathID,
			Title:    item.Title,
			Children: g.convertOutlineToChapters(item.Children),
		}
		chapters = append(chapters, chapter)
	}
	return chapters
}

// extractID extracts the ID from a path like /plateaudocument/toc1 -> toc1
func extractID(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return path
}

// filterOutline filters outline items based on pathFilter
// It matches items whose ID starts with any of the filter paths
func (g *Generator) filterOutline(items []OutlineItem) []OutlineItem {
	var result []OutlineItem
	for _, item := range items {
		if g.matchesFilter(item.ID) {
			result = append(result, item)
		} else if len(item.Children) > 0 {
			// Check if any children match
			filteredChildren := g.filterOutline(item.Children)
			if len(filteredChildren) > 0 {
				// Include parent with filtered children
				filtered := item
				filtered.Children = filteredChildren
				result = append(result, filtered)
			}
		}
	}
	return result
}

// matchesFilter checks if an ID matches any of the path filters
func (g *Generator) matchesFilter(id string) bool {
	for _, filter := range g.pathFilter {
		// Exact match or prefix match (e.g., "toc1" matches "toc1", "toc1_01", "toc1_01_02")
		if id == filter || strings.HasPrefix(id, filter+"_") {
			return true
		}
	}
	return false
}
