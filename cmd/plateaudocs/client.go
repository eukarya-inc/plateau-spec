package plateaudocs

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// GitHubRawURL is the base URL for GitHub raw content
	GitHubRawURL = "https://raw.githubusercontent.com"
	// GitHubOwner is the repository owner
	GitHubOwner = "eukarya-inc"
	// GitHubRepo is the repository name
	GitHubRepo = "plateau-spec"
	// GitHubBranch is the default branch
	GitHubBranch = "main"
	// DefaultBaseURL is the default GitHub raw content base URL
	DefaultBaseURL = GitHubRawURL + "/" + GitHubOwner + "/" + GitHubRepo + "/" + GitHubBranch + "/docs"
	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout = 30 * time.Second
)

// Client fetches PLATEAU specification documents from GitHub raw content
type Client struct {
	BaseURL     string
	HTTPClient  *http.Client
	disableGzip bool
}

// RawResult contains the response body and metadata.
// Note: When gzip compression is enabled (default), ContentLength represents
// the compressed size from the HTTP response, not the decompressed size.
// The actual decompressed size is unknown until fully read.
type RawResult struct {
	Body          io.ReadCloser
	ContentLength int64  // Compressed size when gzip enabled; -1 if unknown
	ContentType   string
}

// Option is a functional option for Client
type Option func(*Client)

// WithBaseURL sets a custom base URL
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// WithoutGzip disables gzip compression
func WithoutGzip() Option {
	return func(c *Client) {
		c.disableGzip = true
	}
}

// New creates a new Client
func New(opts ...Option) *Client {
	c := &Client{
		BaseURL: DefaultBaseURL,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		disableGzip: false,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// GetMarkdown fetches a markdown file for the given document type and path
// docType: "standard" or "procedure"
// path: "toc1", "toc1_01", etc.
func (c *Client) GetMarkdown(ctx context.Context, docType, path string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s.md", c.BaseURL, docType, path)
	return c.fetchText(ctx, url)
}

// GetMarkdownRaw fetches a markdown file with metadata
// docType: "standard" or "procedure"
// path: "toc1", "toc1_01", etc.
// Caller is responsible for closing the returned Body.
func (c *Client) GetMarkdownRaw(ctx context.Context, docType, path string) (*RawResult, error) {
	url := fmt.Sprintf("%s/%s/%s.md", c.BaseURL, docType, path)
	return c.fetchRaw(ctx, url)
}

// GetJSON fetches a JSON file for the given document type and path
// docType: "standard" or "procedure"
// path: "toc1", "toc1_01", etc.
func (c *Client) GetJSON(ctx context.Context, docType, path string) (map[string]any, error) {
	url := fmt.Sprintf("%s/%s/json/%s.json", c.BaseURL, docType, path)
	return c.fetchJSON(ctx, url)
}

// GetJSONRaw fetches a JSON file with metadata
// docType: "standard" or "procedure"
// path: "toc1", "toc1_01", etc.
// Caller is responsible for closing the returned Body.
func (c *Client) GetJSONRaw(ctx context.Context, docType, path string) (*RawResult, error) {
	url := fmt.Sprintf("%s/%s/json/%s.json", c.BaseURL, docType, path)
	return c.fetchRaw(ctx, url)
}

// GetIndex fetches the index.json for the given document type
// docType: "standard" or "procedure"
func (c *Client) GetIndex(ctx context.Context, docType string) (*Index, error) {
	url := fmt.Sprintf("%s/%s/index.json", c.BaseURL, docType)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if !c.disableGzip {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch index: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := c.wrapResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress response: %w", err)
	}
	defer func() { _ = body.Close() }()

	var index Index
	if err := json.NewDecoder(body).Decode(&index); err != nil {
		return nil, fmt.Errorf("failed to decode index: %w", err)
	}

	return &index, nil
}

// GetIndexMarkdown fetches the index.md for the given document type
// docType: "standard" or "procedure"
func (c *Client) GetIndexMarkdown(ctx context.Context, docType string) (string, error) {
	url := fmt.Sprintf("%s/%s/index.md", c.BaseURL, docType)
	return c.fetchText(ctx, url)
}

// GetIndexRaw fetches the index.json with metadata
// docType: "standard" or "procedure"
// Caller is responsible for closing the returned Body.
func (c *Client) GetIndexRaw(ctx context.Context, docType string) (*RawResult, error) {
	url := fmt.Sprintf("%s/%s/index.json", c.BaseURL, docType)
	return c.fetchRaw(ctx, url)
}

// GetIndexMarkdownRaw fetches the index.md with metadata
// docType: "standard" or "procedure"
// Caller is responsible for closing the returned Body.
func (c *Client) GetIndexMarkdownRaw(ctx context.Context, docType string) (*RawResult, error) {
	url := fmt.Sprintf("%s/%s/index.md", c.BaseURL, docType)
	return c.fetchRaw(ctx, url)
}

// GetImage fetches an image file with metadata
// docType: "standard" or "procedure"
// filename: "toc4_img051.png", etc.
// Caller is responsible for closing the returned Body.
func (c *Client) GetImage(ctx context.Context, docType, filename string) (*RawResult, error) {
	url := fmt.Sprintf("%s/%s/images/%s", c.BaseURL, docType, filename)
	return c.fetchRaw(ctx, url)
}

func (c *Client) fetchText(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	if !c.disableGzip {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := c.wrapResponseBody(resp)
	if err != nil {
		return "", fmt.Errorf("failed to decompress response: %w", err)
	}
	defer func() { _ = body.Close() }()

	data, err := io.ReadAll(body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(data), nil
}

func (c *Client) fetchJSON(ctx context.Context, url string) (map[string]any, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if !c.disableGzip {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := c.wrapResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress response: %w", err)
	}
	defer func() { _ = body.Close() }()

	var result map[string]any
	if err := json.NewDecoder(body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return result, nil
}

func (c *Client) fetchRaw(ctx context.Context, url string) (*RawResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if !c.disableGzip {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := c.wrapResponseBodyRaw(resp)
	if err != nil {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("failed to decompress response: %w", err)
	}

	return &RawResult{
		Body:          body,
		ContentLength: resp.ContentLength,
		ContentType:   resp.Header.Get("Content-Type"),
	}, nil
}

// wrapResponseBody wraps the response body with gzip decompression if needed.
// Returns a new reader that decompresses the content.
func (c *Client) wrapResponseBody(resp *http.Response) (io.ReadCloser, error) {
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		return gr, nil
	}
	return resp.Body, nil
}

// wrapResponseBodyRaw wraps the response body with gzip decompression if needed.
// Returns a combined closer that closes both gzip reader and underlying body.
func (c *Client) wrapResponseBodyRaw(resp *http.Response) (io.ReadCloser, error) {
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		return &gzipReadCloser{
			gzipReader: gr,
			underlying: resp.Body,
		}, nil
	}
	return resp.Body, nil
}

// gzipReadCloser wraps a gzip.Reader and its underlying reader
// to ensure both are closed properly.
type gzipReadCloser struct {
	gzipReader *gzip.Reader
	underlying io.ReadCloser
}

func (g *gzipReadCloser) Read(p []byte) (int, error) {
	return g.gzipReader.Read(p)
}

func (g *gzipReadCloser) Close() error {
	if err := g.gzipReader.Close(); err != nil {
		_ = g.underlying.Close()
		return err
	}
	return g.underlying.Close()
}
