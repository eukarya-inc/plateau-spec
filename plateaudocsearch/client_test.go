package plateaudocsearch

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	mokuji "github.com/reearth/mokuji/go"
)

var testDocs = []Document{
	{DocType: DocTypeStandard, ID: "toc4_02", Title: "4.2 建築物モデル", Text: "4.2 建築物モデル 建築物のLOD1は立ち上げモデルとする。"},
	{DocType: DocTypeStandard, ID: "toc6", Title: "6 データ品質", Text: "6 データ品質 品質要求を定める。"},
	{DocType: DocTypeProcedure, ID: "tocC_02_01", Title: "C.2.1 建築物モデル", Text: "C.2.1 建築物モデル 建築物を作成する手順。建築物の外形を取得する。"},
	{DocType: DocTypeProcedure, ID: "tocA_03", Title: "A.3 品質評価", Text: "A.3 品質評価 ＬＯＤ２の品質を評価する。"},
}

func newTestClient(t *testing.T) (*Client, *atomic.Int32) {
	t.Helper()
	data, err := Build(testDocs)
	if err != nil {
		t.Fatal(err)
	}
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		_, _ = w.Write(data)
	}))
	t.Cleanup(server.Close)
	return New(WithReleaseURL(server.URL), WithCacheDir(t.TempDir())), &requests
}

func TestSearch(t *testing.T) {
	ctx := context.Background()
	client, _ := newTestClient(t)
	if _, err := client.Init(ctx); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	all, err := client.Search(ctx, DocTypeAll, "建築物")
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 2 {
		t.Fatalf("got %d results, want 2: %+v", len(all), all)
	}
	// The procedure section says 建築物 three times in a text of similar length.
	if all[0].ID != "tocC_02_01" || all[0].DocType != DocTypeProcedure || all[0].Path != "tocC_02_01.md" {
		t.Errorf("first result = %+v", all[0])
	}
	if all[0].Score <= all[1].Score {
		t.Errorf("scores not descending: %v, %v", all[0].Score, all[1].Score)
	}

	standard, err := client.Search(ctx, DocTypeStandard, "建築物")
	if err != nil {
		t.Fatal(err)
	}
	if len(standard) != 1 || standard[0].ID != "toc4_02" || standard[0].Title != "4.2 建築物モデル" {
		t.Errorf("standard results = %+v", standard)
	}

	limited, err := client.Search(ctx, DocTypeAll, "品質", WithLimit(1))
	if err != nil {
		t.Fatal(err)
	}
	if len(limited) != 1 {
		t.Errorf("got %d results with limit 1", len(limited))
	}

	none, err := client.Search(ctx, DocTypeAll, "存在しない語")
	if err != nil {
		t.Fatal(err)
	}
	if len(none) != 0 {
		t.Errorf("got %+v for an absent term", none)
	}

	if _, err := client.Search(ctx, "unknown", "建築物"); err == nil {
		t.Error("an unknown document type did not fail")
	}
}

func TestSearchNormalizesWidth(t *testing.T) {
	ctx := context.Background()
	client, _ := newTestClient(t)
	if _, err := client.Init(ctx); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	// The text says ＬＯＤ２ in full width; NFKC and lowercasing make it lod2.
	results, err := client.Search(ctx, DocTypeProcedure, "lod2")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 || results[0].ID != "tocA_03" {
		t.Fatalf("results = %+v", results)
	}
	want := "A.3 品質評価 <mark>LOD2</mark>の品質を評価する。"
	if len(results[0].Snippets) != 1 || results[0].Snippets[0] != want {
		t.Errorf("snippets = %q, want %q", results[0].Snippets, want)
	}
}

func TestInitCaches(t *testing.T) {
	ctx := context.Background()
	client, requests := newTestClient(t)

	first, err := client.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !first.Downloaded {
		t.Error("first Init did not download")
	}
	if _, err := os.Stat(filepath.Join(client.cacheDir, indexFileName)); err != nil {
		t.Errorf("index not cached: %v", err)
	}

	second, err := client.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if second.Downloaded || requests.Load() != 1 {
		t.Errorf("second Init downloaded again (%d requests)", requests.Load())
	}

	refreshed, err := client.Refresh(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !refreshed.Downloaded || requests.Load() != 2 {
		t.Errorf("Refresh did not download (%d requests)", requests.Load())
	}
	_ = client.Close()

	if _, err := client.Search(ctx, DocTypeAll, "建築物"); err == nil {
		t.Error("Search after Close did not fail")
	}
}

func TestInitRejectsBadDownload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("not an index"))
	}))
	defer server.Close()
	dir := t.TempDir()
	client := New(WithReleaseURL(server.URL), WithCacheDir(dir))
	if _, err := client.Init(context.Background()); err == nil {
		t.Fatal("Init accepted a file that is not an index")
	}
	if _, err := os.Stat(filepath.Join(dir, indexFileName)); err == nil {
		t.Error("a bad download was cached")
	}
}

func TestSnippets(t *testing.T) {
	text := "前置き建築物と<道路>。" + strings.Repeat("あ", 100) + "建築物"
	got := snippets(text, queryPieces("建築物"))
	if len(got) != 2 {
		t.Fatalf("got %d snippets: %q", len(got), got)
	}
	if want := "前置き<mark>建築物</mark>と&lt;道路&gt;。" + strings.Repeat("あ", 34) + "…"; got[0] != want {
		t.Errorf("snippet 0 = %q, want %q", got[0], want)
	}
	if want := "…" + strings.Repeat("あ", 40) + "<mark>建築物</mark>"; got[1] != want {
		t.Errorf("snippet 1 = %q, want %q", got[1], want)
	}
}

func TestSearchFiltersBeforeLimit(t *testing.T) {
	// Five procedure sections rank above the one standard section: a filter
	// applied after the limit would leave nothing.
	var docs []Document
	for i := range 5 {
		docs = append(docs, Document{DocType: DocTypeProcedure, ID: strings.Repeat("p", i+1), Title: "p", Text: "建築物 建築物"})
	}
	docs = append(docs, Document{DocType: DocTypeStandard, ID: "s", Title: "s", Text: "建築物 " + strings.Repeat("他の語 ", 50)})
	data, err := Build(docs)
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write(data) }))
	defer server.Close()
	client := New(WithReleaseURL(server.URL), WithCacheDir(t.TempDir()))
	ctx := context.Background()
	if _, err := client.Init(ctx); err != nil {
		t.Fatal(err)
	}

	results, err := client.Search(ctx, DocTypeStandard, "建築物", WithLimit(1))
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 || results[0].ID != "s" {
		t.Errorf("results = %+v", results)
	}
}

func TestInitReplacesIncompatibleCache(t *testing.T) {
	ctx := context.Background()
	client, requests := newTestClient(t)
	path := filepath.Join(client.cacheDir, indexFileName)
	if err := os.WriteFile(path, []byte("an index of an earlier version"), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := client.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Downloaded || requests.Load() != 1 {
		t.Errorf("the incompatible cache was not replaced (%d requests)", requests.Load())
	}
	if _, err := client.Search(ctx, DocTypeStandard, "建築物"); err != nil {
		t.Error(err)
	}
}

func TestOpenChecksLayout(t *testing.T) {
	ctx := context.Background()
	data, err := Build(testDocs)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := open(ctx, data); err != nil {
		t.Errorf("an index this package built did not open: %v", err)
	}

	// The same tokenizer and column, but no layout: an index some other
	// version of this package built.
	options := mokuji.DefaultWriteOptions()
	options.DocLengths = true
	builder, err := mokuji.NewBuilder(options)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := builder.AddDocument([]byte(encodePayload(testDocs[0]))); err != nil {
		t.Fatal(err)
	}
	if err := builder.AddColumn(mokuji.Column{Name: docTypeColumn, Values: []int64{0}}); err != nil {
		t.Fatal(err)
	}
	if err := builder.AddPosting([]byte("建築"), 0, 1); err != nil {
		t.Fatal(err)
	}
	meta, err := (&mokuji.Metadata{Generation: 1, GeneratorName: mokuji.GeneratorName, GeneratorVersion: mokuji.Version, GeneratorParams: options.Params(), TokenizerID: tokenizerID}).Encode()
	if err != nil {
		t.Fatal(err)
	}
	other, err := builder.Finish(meta)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := open(ctx, other); err == nil {
		t.Error("an index without the layout opened")
	}
}
