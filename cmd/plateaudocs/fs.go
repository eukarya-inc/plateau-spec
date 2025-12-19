package plateaudocs

import (
	"context"
	"io/fs"
	"strings"
	"time"
)

// FS implements fs.FS for reading PLATEAU documents from GitHub raw content.
// File paths should be in the format: {docType}/{path}
// Examples:
//   - standard/index.json
//   - standard/index.md
//   - standard/toc1.md
//   - standard/json/toc1.json
//   - standard/images/toc4_img051.png
//   - procedure/toc1.md
type FS struct {
	client *Client
	ctx    context.Context
}

// FSOption is a functional option for FS
type FSOption func(*FS)

// WithFSContext sets a custom context for FS
func WithFSContext(ctx context.Context) FSOption {
	return func(f *FS) {
		f.ctx = ctx
	}
}

// WithFSClient sets a custom Client for FS
func WithFSClient(client *Client) FSOption {
	return func(f *FS) {
		f.client = client
	}
}

// NewFS creates a new FS
func NewFS(opts ...FSOption) *FS {
	f := &FS{
		client: New(),
		ctx:    context.Background(),
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

// Open implements fs.FS
func (f *FS) Open(name string) (fs.File, error) {
	// Validate path
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrInvalid}
	}

	// Parse path: {docType}/{rest...}
	parts := strings.SplitN(name, "/", 2)
	if len(parts) < 2 {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
	}

	docType := parts[0]
	rest := parts[1]

	// Fetch the file using Client methods
	var result *RawResult
	var err error

	switch {
	case rest == "index.json":
		result, err = f.client.GetIndexRaw(f.ctx, docType)
	case rest == "index.md":
		result, err = f.client.GetIndexMarkdownRaw(f.ctx, docType)
	case strings.HasPrefix(rest, "json/") && strings.HasSuffix(rest, ".json"):
		// json/toc1.json -> toc1
		path := strings.TrimPrefix(rest, "json/")
		path = strings.TrimSuffix(path, ".json")
		result, err = f.client.GetJSONRaw(f.ctx, docType, path)
	case strings.HasPrefix(rest, "images/"):
		// images/toc4_img051.png -> toc4_img051.png
		filename := strings.TrimPrefix(rest, "images/")
		result, err = f.client.GetImage(f.ctx, docType, filename)
	case strings.HasSuffix(rest, ".md"):
		// toc1.md -> toc1
		path := strings.TrimSuffix(rest, ".md")
		result, err = f.client.GetMarkdownRaw(f.ctx, docType, path)
	default:
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
	}

	if err != nil {
		return nil, &fs.PathError{Op: "open", Path: name, Err: err}
	}

	return &fsFile{
		name:   name,
		result: result,
	}, nil
}

// fsFile implements fs.File
type fsFile struct {
	name   string
	result *RawResult
}

// Stat implements fs.File
func (f *fsFile) Stat() (fs.FileInfo, error) {
	return &fsFileInfo{
		name:        f.name,
		size:        f.result.ContentLength,
		contentType: f.result.ContentType,
	}, nil
}

// Read implements fs.File
func (f *fsFile) Read(p []byte) (int, error) {
	return f.result.Body.Read(p)
}

// Close implements fs.File
func (f *fsFile) Close() error {
	return f.result.Body.Close()
}

// fsFileInfo implements fs.FileInfo
// Note: Size() returns the compressed size when gzip is enabled,
// not the actual decompressed file size.
type fsFileInfo struct {
	name        string
	size        int64
	contentType string
}

// Name implements fs.FileInfo
func (fi *fsFileInfo) Name() string {
	parts := strings.Split(fi.name, "/")
	return parts[len(parts)-1]
}

// Size implements fs.FileInfo
func (fi *fsFileInfo) Size() int64 {
	return fi.size
}

// Mode implements fs.FileInfo
func (fi *fsFileInfo) Mode() fs.FileMode {
	return 0444 // read-only
}

// ModTime implements fs.FileInfo
func (fi *fsFileInfo) ModTime() time.Time {
	return time.Time{} // unknown
}

// IsDir implements fs.FileInfo
func (fi *fsFileInfo) IsDir() bool {
	return false
}

// Sys implements fs.FileInfo
func (fi *fsFileInfo) Sys() any {
	return nil
}
