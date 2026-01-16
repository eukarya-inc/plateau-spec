// Package plateaudocsearch provides full-text search for PLATEAU specification documents.
// It downloads the search index from GitHub Release and provides search functionality.
package plateaudocsearch

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve/v2"

	// Import search package to register kagome tokenizer
	_ "github.com/eukarya-inc/plateau-spec/plateaudocsearch/search"
)

const (
	// DefaultReleaseURL is the default URL for the search index release
	DefaultReleaseURL = "https://github.com/eukarya-inc/plateau-spec/releases/download/search-index-latest/search-index.tar.gz"
	// DefaultTimeout is the default HTTP client timeout for downloading
	DefaultTimeout = 5 * time.Minute
	// DefaultSearchLimit is the default number of search results
	DefaultSearchLimit = 10
)

// DocType represents the document type
type DocType string

const (
	// DocTypeStandard is the standard specification
	DocTypeStandard DocType = "standard"
	// DocTypeProcedure is the procedure specification
	DocTypeProcedure DocType = "procedure"
	// DocTypeAll searches all document types
	DocTypeAll DocType = ""
)

// Client provides full-text search for PLATEAU specification documents
type Client struct {
	releaseURL string
	cacheDir   string
	httpClient *http.Client
	indexes    map[DocType]bleve.Index
}

// SearchResult represents a single search result
type SearchResult struct {
	ID       string   `json:"id"`
	DocType  DocType  `json:"docType"`
	Path     string   `json:"path"`
	Title    string   `json:"title"`
	Score    float64  `json:"score"`
	Snippets []string `json:"snippets,omitempty"`
}

// InitResult contains information about the initialization
type InitResult struct {
	Downloaded   bool          // Whether the index was downloaded (false if cached)
	DownloadTime time.Duration // Time taken to download (0 if cached)
}

// searchOptions holds search configuration
type searchOptions struct {
	limit int
}

// SearchOption is a functional option for Search
type SearchOption func(*searchOptions)

// WithLimit sets the maximum number of search results
func WithLimit(limit int) SearchOption {
	return func(o *searchOptions) {
		o.limit = limit
	}
}

// Option is a functional option for Client
type Option func(*Client)

// WithReleaseURL sets a custom release URL
func WithReleaseURL(url string) Option {
	return func(c *Client) {
		c.releaseURL = url
	}
}

// WithCacheDir sets a custom cache directory
func WithCacheDir(dir string) Option {
	return func(c *Client) {
		c.cacheDir = dir
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// New creates a new Client.
// Call Init() to download the search index and initialize the client.
func New(opts ...Option) *Client {
	c := &Client{
		releaseURL: DefaultReleaseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		indexes: make(map[DocType]bleve.Index),
	}

	for _, opt := range opts {
		opt(c)
	}

	// Set default cache directory
	if c.cacheDir == "" {
		cacheDir, err := os.UserCacheDir()
		if err != nil {
			cacheDir = os.TempDir()
		}
		c.cacheDir = filepath.Join(cacheDir, "plateaudocsearch")
	}

	return c
}

// Init downloads the search index (if not cached) and opens the indexes.
// This must be called before Search().
// Returns InitResult with download information.
func (c *Client) Init(ctx context.Context) (*InitResult, error) {
	result := &InitResult{}

	// Ensure cache directory exists
	if err := os.MkdirAll(c.cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	// Download and extract index if needed
	downloaded, downloadTime, err := c.ensureIndex(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure index: %w", err)
	}
	result.Downloaded = downloaded
	result.DownloadTime = downloadTime

	// Open indexes
	if err := c.openIndexes(); err != nil {
		return nil, fmt.Errorf("failed to open indexes: %w", err)
	}

	return result, nil
}

// Close closes all opened indexes
func (c *Client) Close() error {
	var lastErr error
	for _, idx := range c.indexes {
		if err := idx.Close(); err != nil {
			lastErr = err
		}
	}
	c.indexes = make(map[DocType]bleve.Index)
	return lastErr
}

// Search performs a full-text search on the specified document type.
// Use DocTypeAll (or "") to search all document types.
func (c *Client) Search(ctx context.Context, docType DocType, query string, opts ...SearchOption) ([]SearchResult, error) {
	// Apply options
	o := &searchOptions{
		limit: DefaultSearchLimit,
	}
	for _, opt := range opts {
		opt(o)
	}

	// Search all document types if docType is empty
	if docType == "" || docType == DocTypeAll {
		var allResults []SearchResult
		for dt := range c.indexes {
			results, err := c.searchIndex(c.indexes[dt], query, o.limit, dt)
			if err != nil {
				return nil, fmt.Errorf("search failed for %s: %w", dt, err)
			}
			allResults = append(allResults, results...)
		}
		// Sort by score (descending) and limit
		sortByScore(allResults)
		if len(allResults) > o.limit {
			allResults = allResults[:o.limit]
		}
		return allResults, nil
	}

	idx, ok := c.indexes[docType]
	if !ok {
		return nil, fmt.Errorf("index not found for document type: %s (did you call Init?)", docType)
	}

	return c.searchIndex(idx, query, o.limit, docType)
}

func (c *Client) searchIndex(idx bleve.Index, query string, limit int, docType DocType) ([]SearchResult, error) {
	q := bleve.NewMatchQuery(query)
	req := bleve.NewSearchRequest(q)
	req.Size = limit
	req.Highlight = bleve.NewHighlight()
	req.Fields = []string{"title", "path"}

	result, err := idx.Search(req)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	var results []SearchResult
	for _, hit := range result.Hits {
		sr := SearchResult{
			ID:      hit.ID,
			DocType: docType,
			Score:   hit.Score,
		}

		if title, ok := hit.Fields["title"].(string); ok {
			sr.Title = title
		}
		if path, ok := hit.Fields["path"].(string); ok {
			sr.Path = path
		}

		if fragments, ok := hit.Fragments["content"]; ok {
			sr.Snippets = fragments
		}

		results = append(results, sr)
	}

	return results, nil
}

func sortByScore(results []SearchResult) {
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Score > results[i].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}

// ensureIndex downloads and extracts the index if not present
// Returns whether download occurred and the download time
func (c *Client) ensureIndex(ctx context.Context) (bool, time.Duration, error) {
	// Check if index already exists
	standardPath := filepath.Join(c.cacheDir, "standard.bleve")
	if _, err := os.Stat(standardPath); err == nil {
		return false, 0, nil // Already exists
	}

	// Download
	start := time.Now()
	tarPath := filepath.Join(c.cacheDir, "search-index.tar.gz")
	if err := c.download(ctx, tarPath); err != nil {
		return false, 0, fmt.Errorf("download failed: %w", err)
	}
	downloadTime := time.Since(start)
	defer func() { _ = os.Remove(tarPath) }()

	// Extract
	if err := c.extract(tarPath); err != nil {
		return false, 0, fmt.Errorf("extract failed: %w", err)
	}

	return true, downloadTime, nil
}

// download downloads the search index from the release URL
func (c *Client) download(ctx context.Context, destPath string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.releaseURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() { _ = out.Close() }()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// extract extracts the tar.gz file to the cache directory
func (c *Client) extract(tarPath string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return fmt.Errorf("failed to open tar file: %w", err)
	}
	defer func() { _ = file.Close() }()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer func() { _ = gzReader.Close() }()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar header: %w", err)
		}

		targetPath := filepath.Join(c.cacheDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %w", err)
			}

			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				_ = outFile.Close()
				return fmt.Errorf("failed to write file: %w", err)
			}
			_ = outFile.Close()
		}
	}

	return nil
}

// openIndexes opens the bleve indexes
func (c *Client) openIndexes() error {
	docTypes := []DocType{DocTypeStandard, DocTypeProcedure}

	for _, docType := range docTypes {
		indexPath := filepath.Join(c.cacheDir, string(docType)+".bleve")
		if _, err := os.Stat(indexPath); os.IsNotExist(err) {
			continue // Skip if not exists
		}

		idx, err := bleve.Open(indexPath)
		if err != nil {
			return fmt.Errorf("failed to open index %s: %w", docType, err)
		}
		c.indexes[docType] = idx
	}

	if len(c.indexes) == 0 {
		return fmt.Errorf("no indexes found")
	}

	return nil
}

// Refresh forces re-download of the search index
func (c *Client) Refresh(ctx context.Context) (*InitResult, error) {
	// Close existing indexes
	_ = c.Close()

	// Remove existing index files
	for _, docType := range []DocType{DocTypeStandard, DocTypeProcedure} {
		indexPath := filepath.Join(c.cacheDir, string(docType)+".bleve")
		_ = os.RemoveAll(indexPath)
	}

	// Re-download and open
	return c.Init(ctx)
}
