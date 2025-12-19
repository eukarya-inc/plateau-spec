package specgen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
