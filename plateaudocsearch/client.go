// Package plateaudocsearch provides full-text search for PLATEAU specification documents.
//
// The search index is a single Mokuji file (https://github.com/reearth/mokuji)
// of about 1.4 MB, downloaded once from GitHub Release, cached on disk, and
// held in memory while the client is open: a search makes no request.
package plateaudocsearch

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	mokuji "github.com/reearth/mokuji/go"
)

const (
	// DefaultReleaseURL is the default URL for the search index release
	DefaultReleaseURL = "https://github.com/eukarya-inc/plateau-spec/releases/download/search-index-latest/search-index.mokuji"
	// DefaultTimeout is the default HTTP client timeout for downloading
	DefaultTimeout = 5 * time.Minute
	// DefaultSearchLimit is the default number of search results
	DefaultSearchLimit = 10

	indexFileName = "search-index.mokuji"
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

	mu    sync.RWMutex
	index *mokuji.Index
}

// SearchResult represents a single search result
type SearchResult struct {
	// ID is the section's path in index.json, such as "toc4_01_01".
	ID      string  `json:"id"`
	DocType DocType `json:"docType"`
	// Path is the section's Markdown file, relative to its document type's
	// directory, such as "toc4_01_01.md".
	Path  string  `json:"path"`
	Title string  `json:"title"`
	Score float64 `json:"score"`
	// Snippets are fragments of the section's text around the matches, HTML
	// escaped, with each match wrapped in <mark>.
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

// Init downloads the search index (if not cached) and loads it into memory.
// This must be called before Search().
// Returns InitResult with download information.
func (c *Client) Init(ctx context.Context) (*InitResult, error) {
	if err := os.MkdirAll(c.cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	result := &InitResult{}
	path := filepath.Join(c.cacheDir, indexFileName)
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		start := time.Now()
		if data, err = c.download(ctx); err != nil {
			return nil, fmt.Errorf("download failed: %w", err)
		}
		result.Downloaded = true
		result.DownloadTime = time.Since(start)
	} else if err != nil {
		return nil, fmt.Errorf("failed to read cached index: %w", err)
	}

	index, err := open(ctx, data)
	if err != nil && !result.Downloaded {
		// A cached file this version cannot search is replaced, once.
		_ = os.Remove(path)
		return c.Init(ctx)
	}
	if err != nil {
		return nil, err
	}

	if result.Downloaded {
		// Cached only once it has opened, so that a bad download is not kept.
		if err := writeFileAtomic(path, data); err != nil {
			return nil, fmt.Errorf("failed to cache index: %w", err)
		}
		removeLegacyCache(c.cacheDir)
	}

	c.mu.Lock()
	c.index = index
	c.mu.Unlock()
	return result, nil
}

// open opens an index and checks it is laid out as this version reads.
func open(ctx context.Context, data []byte) (*mokuji.Index, error) {
	index, err := mokuji.Open(ctx, mokuji.BytesSource(data))
	if err != nil {
		return nil, fmt.Errorf("failed to open index: %w", err)
	}
	if layout := layoutOf(index); layout != indexLayout || index.TokenizerID != tokenizerID || index.Column(docTypeColumn) == nil {
		_ = index.Close()
		return nil, fmt.Errorf("not an index of layout %s (it records %q)", indexLayout, layout)
	}
	return index, nil
}

// Close releases the index held in memory.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.index == nil {
		return nil
	}
	err := c.index.Close()
	c.index = nil
	return err
}

// Search performs a full-text search on the specified document type.
// Use DocTypeAll (or "") to search all document types.
//
// A query matches the sections that contain every one of its terms: the
// character bigrams of its Japanese runs and its ASCII words whole, split
// where the two meet (Mokuji's `bigram`, as the index was built). Results are
// ranked by BM25.
func (c *Client) Search(ctx context.Context, docType DocType, query string, opts ...SearchOption) ([]SearchResult, error) {
	o := &searchOptions{
		limit: DefaultSearchLimit,
	}
	for _, opt := range opts {
		opt(o)
	}
	if docType != DocTypeAll && docType != DocTypeStandard && docType != DocTypeProcedure {
		return nil, fmt.Errorf("unknown document type: %s", docType)
	}

	c.mu.RLock()
	index := c.index
	c.mu.RUnlock()
	if index == nil {
		return nil, errors.New("index is not loaded (did you call Init?)")
	}

	options := []mokuji.SearchOption{mokuji.Limit(o.limit)}
	if docType != DocTypeAll {
		types, want := index.Column(docTypeColumn).Values, docTypeValue(docType)
		options = append(options, mokuji.Filter(func(doc int) bool { return types[doc] == want }))
	}
	found, err := index.Search(ctx, query, options...)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	pieces := queryPieces(query)
	results := []SearchResult{}
	for _, r := range found.Results {
		doc, err := decodePayload(r.Payload)
		if err != nil {
			return nil, fmt.Errorf("document %d: %w", r.Doc, err)
		}
		results = append(results, SearchResult{
			ID:       doc.ID,
			DocType:  doc.DocType,
			Path:     doc.ID + ".md",
			Title:    doc.Title,
			Score:    r.Score,
			Snippets: snippets(doc.Text, pieces),
		})
	}
	return results, nil
}

// Refresh forces re-download of the search index
func (c *Client) Refresh(ctx context.Context) (*InitResult, error) {
	_ = c.Close()
	if err := os.Remove(filepath.Join(c.cacheDir, indexFileName)); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to remove cached index: %w", err)
	}
	return c.Init(ctx)
}

// download fetches the search index from the release URL
func (c *Client) download(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.releaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func writeFileAtomic(path string, data []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".*")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tmp.Name()) }()
	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

// removeLegacyCache removes the Bleve indexes earlier versions cached.
func removeLegacyCache(cacheDir string) {
	for _, name := range []string{"standard.bleve", "procedure.bleve"} {
		_ = os.RemoveAll(filepath.Join(cacheDir, name))
	}
}
