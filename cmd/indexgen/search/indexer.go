// Package search collects the sections of PLATEAU specification documents to
// be indexed for full-text search.
package search

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/eukarya-inc/plateau-spec/cmd/metanorma2md"
	"github.com/eukarya-inc/plateau-spec/plateaudocsearch"
)

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

// page is a section as specgen saved it in json/: its own content, without
// its subsections'.
type page struct {
	Title   string                 `json:"title"`
	Content []metanorma2md.Content `json:"content"`
}

// LoadIndexJSON loads and parses index.json from the docs directory.
func LoadIndexJSON(docsDir string) (*IndexJSON, error) {
	data, err := os.ReadFile(filepath.Join(docsDir, "index.json"))
	if err != nil {
		return nil, err
	}

	var indexJSON IndexJSON
	if err := json.Unmarshal(data, &indexJSON); err != nil {
		return nil, err
	}
	return &indexJSON, nil
}

// CollectDocuments returns every section listed in docsDir/index.json, in
// outline order, each with its own text only.
//
// A section's Markdown file contains its subsections as well, so indexing the
// files would count a subsection once for every section above it. The JSON
// specgen saved beside them holds each section's own content, which is
// converted here instead.
func CollectDocuments(docsDir string, docType plateaudocsearch.DocType) ([]plateaudocsearch.Document, error) {
	indexJSON, err := LoadIndexJSON(docsDir)
	if err != nil {
		return nil, err
	}

	var docs []plateaudocsearch.Document
	var walk func([]Chapter) error
	walk = func(chapters []Chapter) error {
		for _, ch := range chapters {
			text, err := sectionText(filepath.Join(docsDir, "json", ch.Path+".json"))
			if err != nil {
				return fmt.Errorf("%s: %w", ch.Path, err)
			}
			docs = append(docs, plateaudocsearch.Document{
				DocType: docType,
				ID:      ch.Path,
				Title:   ch.Title,
				Text:    text,
			})
			if err := walk(ch.Children); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(indexJSON.Chapters); err != nil {
		return nil, err
	}
	return docs, nil
}

var (
	// An image is nothing to search; a link is its text.
	mdImage = regexp.MustCompile(`!\[[^\]]*\]\([^)]*\)|\[Image\]`)
	mdLink  = regexp.MustCompile(`\[([^\]]*)\]\([^)]*\)`)
	// Table rules, and the marks of headings, emphasis and table cells.
	mdRule   = regexp.MustCompile(`\|?(\s*:?-{3,}:?\s*\|)+`)
	mdMarkup = regexp.MustCompile(`[#*|]+`)
)

// sectionText is a section's own content as plain text: converted to
// Markdown as specgen converts it, then with the markup taken out, since a
// snippet is cut from it.
func sectionText(jsonPath string) (string, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return "", err
	}
	var p page
	if err := json.Unmarshal(data, &p); err != nil {
		return "", err
	}

	md := metanorma2md.Convert(&metanorma2md.Document{Content: p.Content}, nil)
	md = mdImage.ReplaceAllString(md, " ")
	md = mdLink.ReplaceAllString(md, "$1")
	md = mdRule.ReplaceAllString(md, " ")
	md = mdMarkup.ReplaceAllString(md, " ")
	return md, nil
}
