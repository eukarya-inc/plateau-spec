package plateaudocs

// Index represents the structure of index.json
type Index struct {
	DocumentType string    `json:"documentType"`
	Title        string    `json:"title"`
	GeneratedAt  string    `json:"generatedAt"`
	Chapters     []Chapter `json:"chapters"`
}

// Chapter represents a chapter in the index
type Chapter struct {
	Path     string    `json:"path"`
	Title    string    `json:"title"`
	Children []Chapter `json:"children,omitempty"`
}
