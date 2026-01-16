// Package search provides indexing functionality for PLATEAU specification documents.
package search

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/mapping"

	// Import to register kagome tokenizer
	_ "github.com/eukarya-inc/plateau-spec/plateaudocsearch/search"
)

const kagomeTokenizerName = "kagome"

// Document represents a searchable document.
type Document struct {
	ID      string `json:"id"`
	Path    string `json:"path"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// IndexJSON represents the structure of index.json.
type IndexJSON struct {
	DocumentType string    `json:"documentType"`
	Title        string    `json:"title"`
	Chapters     []Chapter `json:"chapters"`
}

// Chapter represents a chapter in index.json.
type Chapter struct {
	Path     string    `json:"path"`
	Title    string    `json:"title"`
	Children []Chapter `json:"children,omitempty"`
}

// CreateIndex creates a new Bleve index with Japanese analyzer at the specified path.
func CreateIndex(indexPath string) (bleve.Index, error) {
	indexMapping := bleve.NewIndexMapping()

	if err := indexMapping.AddCustomTokenizer("kagome", map[string]interface{}{
		"type": kagomeTokenizerName,
	}); err != nil {
		return nil, err
	}

	if err := indexMapping.AddCustomAnalyzer("ja", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": "kagome",
		"token_filters": []string{
			lowercase.Name,
		},
	}); err != nil {
		return nil, err
	}

	indexMapping.DefaultAnalyzer = "ja"

	docMapping := bleve.NewDocumentMapping()
	textFieldMapping := mapping.NewTextFieldMapping()
	textFieldMapping.Analyzer = "ja"

	docMapping.AddFieldMappingsAt("title", textFieldMapping)
	docMapping.AddFieldMappingsAt("content", textFieldMapping)
	indexMapping.AddDocumentMapping("doc", docMapping)

	return bleve.New(indexPath, indexMapping)
}

// LoadIndexJSON loads and parses index.json from the docs directory.
func LoadIndexJSON(docsDir string) (*IndexJSON, error) {
	indexPath := filepath.Join(docsDir, "index.json")
	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, err
	}

	var indexJSON IndexJSON
	if err := json.Unmarshal(data, &indexJSON); err != nil {
		return nil, err
	}

	return &indexJSON, nil
}

// GetRootPaths extracts root-level paths from index.json chapters.
// Only includes paths without underscore (e.g., "toc1", "toc3" but not "toc0_01", "toc2_01").
func GetRootPaths(indexJSON *IndexJSON) []string {
	var paths []string
	for _, chapter := range indexJSON.Chapters {
		if !strings.Contains(chapter.Path, "_") {
			paths = append(paths, chapter.Path)
		}
	}
	return paths
}

// IndexFromJSON indexes only the root-level documents specified in index.json.
func IndexFromJSON(index bleve.Index, docsDir string) (int, error) {
	indexJSON, err := LoadIndexJSON(docsDir)
	if err != nil {
		return 0, err
	}

	rootPaths := GetRootPaths(indexJSON)
	count := 0

	for _, p := range rootPaths {
		mdPath := filepath.Join(docsDir, p+".md")
		content, err := os.ReadFile(mdPath)
		if err != nil {
			// Skip if file doesn't exist
			continue
		}

		doc := Document{
			ID:      p,
			Path:    p + ".md",
			Title:   extractTitleFromJSON(indexJSON, p),
			Content: string(content),
		}

		if err := index.Index(doc.ID, doc); err != nil {
			return count, err
		}
		count++
	}

	return count, nil
}

// extractTitleFromJSON finds the title for a path from index.json.
func extractTitleFromJSON(indexJSON *IndexJSON, path string) string {
	for _, chapter := range indexJSON.Chapters {
		if chapter.Path == path {
			return chapter.Title
		}
	}
	return ""
}

// CompressIndex compresses the index directory contents to a .tar.gz file.
// The contents are placed at the root level (e.g., standard.bleve/, procedure.bleve/).
func CompressIndex(srcDir, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	gzWriter := gzip.NewWriter(outFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == srcDir {
			return nil
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// Use path relative to srcDir (not its parent)
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tarWriter, file)
		return err
	})
}
