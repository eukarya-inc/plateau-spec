package search

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/eukarya-inc/plateau-spec/plateaudocsearch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectDocuments(t *testing.T) {
	dir := t.TempDir()
	write := func(name, content string) {
		t.Helper()
		path := filepath.Join(dir, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0755))
		require.NoError(t, os.WriteFile(path, []byte(content), 0644))
	}
	write("index.json", `{"chapters": [{"path": "toc1", "title": "1 概要", "children": [{"path": "toc1_01", "title": "1.1 目的"}]}]}`)
	write("json/toc1.json", `{"title": "1 概要", "content": [
		{"type": "paragraph", "content": {"content": [{"type": "text", "text": "概要の本文。"}]}}
	]}`)
	write("json/toc1_01.json", `{"title": "1.1 目的", "content": [
		{"type": "paragraph", "content": {"content": [{"type": "text", "text": "目的の本文。"}]}}
	]}`)

	docs, err := CollectDocuments(dir, plateaudocsearch.DocTypeStandard)
	require.NoError(t, err)
	require.Len(t, docs, 2)

	assert.Equal(t, "toc1", docs[0].ID)
	assert.Equal(t, "1 概要", docs[0].Title)
	assert.Equal(t, plateaudocsearch.DocTypeStandard, docs[0].DocType)
	assert.Contains(t, docs[0].Text, "概要の本文。")
	// A section's text is its own, not its subsections'.
	assert.NotContains(t, docs[0].Text, "目的の本文。")

	assert.Equal(t, "toc1_01", docs[1].ID)
	assert.Contains(t, docs[1].Text, "目的の本文。")
}

func TestSectionTextStripsMarkup(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "page.json")
	require.NoError(t, os.WriteFile(path, []byte(`{"content": [
		{"type": "paragraph", "content": {"content": [{"type": "text", "text": "本文", "marks": [{"type": "strong"}]}]}}
	]}`), 0644))

	text, err := sectionText(path)
	require.NoError(t, err)
	assert.Contains(t, text, "本文")
	assert.NotContains(t, text, "*")
}
