package plateaudocsearch

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	mokuji "github.com/reearth/mokuji/go"
	"golang.org/x/text/unicode/norm"
)

// tokenizerID is the Mokuji tokenizer the index is built and searched with.
// Character bigrams need no dictionary, so a query is analyzed by the reader
// alone, exactly as the documents were by the writer.
const tokenizerID = "bigram"

// indexLayout names how this package lays out an index: the payload's
// fields, the columns, and the tokenizer's rule. It is recorded in the
// metadata's generator.params, and a client opens only an index of the layout
// it reads. Change it whenever any of those change.
//
// Neither the tokenizer identifier nor the Mokuji version says this: `bigram`
// changed its rule in Mokuji 0.2 under the same identifier.
const (
	indexLayoutParam = "plateaudocsearch.layout"
	indexLayout      = "1"
)

// docTypeColumn is the attribute column saying which document a section is
// in: 0 for the standard specification, 1 for the procedure. A search for one
// of them filters on it before ranking, without reading any payload.
const docTypeColumn = "doctype"

// payloadSep separates the fields of a document's payload. It cannot occur in
// a title or in text, which are normalized before they are written.
const payloadSep = "\x1f"

// Document is one section of a specification, as it is indexed.
type Document struct {
	DocType DocType
	// ID is the section's path in index.json, such as "toc4_01_01".
	ID    string
	Title string
	// Text is the section's own text, without the text of its subsections,
	// so that a match is reported at the section where it occurs.
	Text string
}

// Build writes a search index for docs, in the order given.
func Build(docs []Document) ([]byte, error) {
	options := mokuji.DefaultWriteOptions()
	// BM25 needs every document's length in tokens.
	options.DocLengths = true
	// A result decompresses its whole payload block, and a payload carries the
	// section's text. Held in memory, a smaller block costs only compression:
	// 16 against the default 64 is 8% larger and searches about twice as fast.
	options.DocsPerPayloadBlock = 16

	builder, err := mokuji.NewBuilder(options)
	if err != nil {
		return nil, err
	}

	types := make([]int64, len(docs))
	for i, d := range docs {
		types[i] = docTypeValue(d.DocType)
	}
	if err := builder.AddColumn(mokuji.Column{Name: docTypeColumn, Values: types}); err != nil {
		return nil, err
	}

	postings := map[string]map[int]uint64{}
	source := sha256.New()
	for _, d := range docs {
		d.Title = cleanField(d.Title)
		d.Text = cleanField(d.Text)
		payload := encodePayload(d)
		_, _ = fmt.Fprintf(source, "%s\n", payload)

		doc, err := builder.AddDocument([]byte(payload))
		if err != nil {
			return nil, err
		}
		terms, _ := mokuji.Split(tokenizerID, d.Text)
		for _, term := range terms {
			if postings[term] == nil {
				postings[term] = map[int]uint64{}
			}
			postings[term][doc]++
		}
	}

	// The writer owns no sort: terms in byte order, and each term's documents
	// ascending.
	for _, term := range slices.Sorted(maps.Keys(postings)) {
		freqs := postings[term]
		for _, doc := range slices.Sorted(maps.Keys(freqs)) {
			if err := builder.AddPosting([]byte(term), doc, freqs[doc]); err != nil {
				return nil, err
			}
		}
	}

	meta, err := (&mokuji.Metadata{
		Generation: 1,
		Sources: []mokuji.Provenance{{
			ContentHash: fmt.Sprintf("sha2-256:%x", source.Sum(nil)),
			URI:         "https://github.com/eukarya-inc/plateau-spec/tree/main/docs",
		}},
		GeneratorName:    mokuji.GeneratorName,
		GeneratorVersion: mokuji.Version,
		GeneratorParams:  params(options),
		TokenizerID:      tokenizerID,
	}).Encode()
	if err != nil {
		return nil, err
	}
	return builder.Finish(meta)
}

// params are the writer's options as the metadata records them, with the
// layout.
func params(options mokuji.WriteOptions) map[string]any {
	p := options.Params()
	p[indexLayoutParam] = indexLayout
	return p
}

// layoutOf is the layout an index records, or "" if it records none.
func layoutOf(index *mokuji.Index) string {
	layout, _ := index.Generator.Params[indexLayoutParam].(string)
	return layout
}

// cleanField puts text in the form it is stored and searched in: NFKC, so
// that a snippet is cut from the same characters the tokenizer saw, with every
// run of white space one space.
func cleanField(s string) string {
	s = norm.NFKC.String(s)
	s = strings.ReplaceAll(s, payloadSep, " ")
	return strings.Join(strings.Fields(s), " ")
}

func docTypeValue(t DocType) int64 {
	if t == DocTypeProcedure {
		return 1
	}
	return 0
}

func encodePayload(d Document) string {
	return strings.Join([]string{string(d.DocType), d.ID, d.Title, d.Text}, payloadSep)
}

func decodePayload(payload string) (Document, error) {
	fields := strings.SplitN(payload, payloadSep, 4)
	if len(fields) != 4 {
		return Document{}, errors.New("malformed payload")
	}
	return Document{DocType: DocType(fields[0]), ID: fields[1], Title: fields[2], Text: fields[3]}, nil
}
