package specgen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// MaxConcurrentHTTPRequests limits concurrent HTTP requests
const MaxConcurrentHTTPRequests = 10

// PlateauClient handles communication with PLATEAU specification API
type PlateauClient struct {
	BaseURL      string
	DocumentType string // "standard" or "procedure"
	HTTPClient   *http.Client
}

// NewClient creates a new PLATEAU specification client
func NewClient(documentType string) *PlateauClient {
	baseURL := "https://www.mlit.go.jp/plateaudocument"
	if documentType == "procedure" {
		baseURL = "https://www.mlit.go.jp/plateaudocument02"
	}

	return &PlateauClient{
		BaseURL:      baseURL,
		DocumentType: documentType,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetOutlineWithDepth retrieves the outline with specified depth
func (c *PlateauClient) GetOutlineWithDepth(depth int) ([]OutlineItem, error) {
	url := fmt.Sprintf("%s/resource-nav.json", c.BaseURL)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch navigation: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var nav PlateauNavigation
	if err := json.NewDecoder(resp.Body).Decode(&nav); err != nil {
		return nil, fmt.Errorf("failed to decode navigation: %w", err)
	}

	// Convert navigation to outline items with specified depth
	return c.convertToOutlineRecursive(nav.Children, depth-1), nil
}

func (c *PlateauClient) getDocPath() string {
	if c.DocumentType == "procedure" {
		return "plateaudocument02"
	}
	return "plateaudocument"
}

// fetchNavigation fetches navigation for a given path
func (c *PlateauClient) fetchNavigation(path string) (*PlateauNavigation, error) {
	// Build URL based on whether path already contains the document base path
	var url string
	expectedBasePath := "/plateaudocument"
	if c.DocumentType == "procedure" {
		expectedBasePath = "/plateaudocument02"
	}

	if strings.HasPrefix(path, expectedBasePath) {
		url = fmt.Sprintf("https://www.mlit.go.jp%s/resource-nav.json", path)
	} else {
		url = fmt.Sprintf("%s%s/resource-nav.json", c.BaseURL, path)
	}

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	var nav PlateauNavigation
	if err := json.NewDecoder(resp.Body).Decode(&nav); err != nil {
		return nil, err
	}

	return &nav, nil
}

// convertToOutlineRecursive converts PlateauNavigation to OutlineItem and fetches children
func (c *PlateauClient) convertToOutlineRecursive(navItems []PlateauNavigation, maxDepth int) []OutlineItem {
	var items []OutlineItem
	for _, nav := range navItems {
		// Skip empty items
		if nav.Path == "" || nav.PlainTitle == "" {
			continue
		}

		pathParts := strings.Split(strings.Trim(nav.Path, "/"), "/")
		id := ""
		if len(pathParts) > 0 {
			id = pathParts[len(pathParts)-1]
		}

		item := OutlineItem{
			ID:    id,
			Title: nav.PlainTitle,
			Path:  nav.Path,
		}

		// Fetch children if depth allows
		if maxDepth > 0 {
			childNav, err := c.fetchNavigation(nav.Path)
			if err == nil && len(childNav.Children) > 0 {
				item.Children = c.convertToOutlineRecursive(childNav.Children, maxDepth-1)
			}
		}

		items = append(items, item)
	}
	return items
}

// GetContentByPath retrieves content for a specific path (single page only)
func (c *PlateauClient) GetContentByPath(path string) (*PlateauDocument, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	var url string
	expectedBasePath := "/plateaudocument"
	if c.DocumentType == "procedure" {
		expectedBasePath = "/plateaudocument02"
	}

	if strings.HasPrefix(path, expectedBasePath) {
		url = fmt.Sprintf("https://www.mlit.go.jp%s/resource-content.json", path)
	} else {
		url = fmt.Sprintf("%s%s/resource-content.json", c.BaseURL, path)
	}

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch content from path %s: %w", path, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code for path %s: %d", path, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse as raw JSON first to handle flexible content structure
	var rawDoc map[string]interface{}
	if err := json.Unmarshal(body, &rawDoc); err != nil {
		return nil, fmt.Errorf("failed to parse content JSON: %w", err)
	}

	// Convert to our document structure
	doc := &PlateauDocument{
		Path:     path,
		Metadata: make(map[string]interface{}),
	}

	// The PLATEAU spec API has a nested structure: content.contentDoc.content
	if contentWrapper, ok := rawDoc["content"].(map[string]interface{}); ok {
		// Extract title from labelInPlainText
		if label, ok := contentWrapper["labelInPlainText"].(string); ok {
			doc.Title = label
		}

		// Get contentDoc
		if contentDoc, ok := contentWrapper["contentDoc"].(map[string]interface{}); ok {
			// Process content array inside contentDoc
			if contents, ok := contentDoc["content"].([]interface{}); ok {
				for _, item := range contents {
					if contentMap, ok := item.(map[string]interface{}); ok {
						content := PlateauContent{
							Type:    "text",
							Content: contentMap,
						}

						// Determine content type
						if contentType, ok := contentMap["type"].(string); ok {
							content.Type = contentType
						}

						doc.Content = append(doc.Content, content)
					}
				}
			}
		}
	} else if title, ok := rawDoc["title"].(string); ok {
		// Fallback for simpler structure
		doc.Title = title

		if contents, ok := rawDoc["content"].([]interface{}); ok {
			for _, item := range contents {
				if contentMap, ok := item.(map[string]interface{}); ok {
					content := PlateauContent{
						Type:    "text",
						Content: contentMap,
					}
					if contentType, ok := contentMap["type"].(string); ok {
						content.Type = contentType
					}
					doc.Content = append(doc.Content, content)
				}
			}
		}
	}

	return doc, nil
}

// pathNode represents a node in the path tree
type pathNode struct {
	path     string
	children []*pathNode
}

// buildPathTree builds a tree structure of paths using parallel fetching
func (c *PlateauClient) buildPathTree(rootPath string) *pathNode {
	root := &pathNode{path: rootPath}
	c.buildPathTreeRecursive([]*pathNode{root})
	return root
}

// buildPathTreeRecursive fetches children for nodes in parallel, level by level
func (c *PlateauClient) buildPathTreeRecursive(nodes []*pathNode) {
	if len(nodes) == 0 {
		return
	}

	// Fetch children for all nodes at this level in parallel
	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, MaxConcurrentHTTPRequests)

	for _, node := range nodes {
		wg.Add(1)
		go func(n *pathNode) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			nav, err := c.fetchNavigation(n.path)
			if err != nil || nav == nil {
				return
			}

			mu.Lock()
			for _, child := range nav.Children {
				if child.Path != "" {
					n.children = append(n.children, &pathNode{path: child.Path})
				}
			}
			mu.Unlock()
		}(node)
	}

	wg.Wait()

	// Collect all children for the next level
	var nextLevel []*pathNode
	for _, node := range nodes {
		nextLevel = append(nextLevel, node.children...)
	}

	// Recursively fetch children for the next level
	if len(nextLevel) > 0 {
		c.buildPathTreeRecursive(nextLevel)
	}
}

// flattenPathTree flattens the tree in DFS order
func (c *PlateauClient) flattenPathTree(root *pathNode) []string {
	if root == nil {
		return nil
	}

	paths := []string{root.path}
	for _, child := range root.children {
		paths = append(paths, c.flattenPathTree(child)...)
	}
	return paths
}
