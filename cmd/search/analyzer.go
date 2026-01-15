// Package search provides full-text search functionality for PLATEAU specification documents.
package search

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

const kagomeTokenizerName = "kagome"

// KagomeTokenizer is a Bleve tokenizer that uses Kagome for Japanese morphological analysis.
type KagomeTokenizer struct {
	t *tokenizer.Tokenizer
}

// NewKagomeTokenizer creates a new KagomeTokenizer with the IPA dictionary.
func NewKagomeTokenizer() (*KagomeTokenizer, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}
	return &KagomeTokenizer{t: t}, nil
}

// Tokenize implements the Bleve Tokenizer interface.
func (kt *KagomeTokenizer) Tokenize(input []byte) analysis.TokenStream {
	tokens := kt.t.Tokenize(string(input))
	rv := make(analysis.TokenStream, 0, len(tokens))
	pos := 1
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		rv = append(rv, &analysis.Token{
			Term:     []byte(token.Surface),
			Start:    token.Start,
			End:      token.End,
			Position: pos,
			Type:     analysis.Ideographic,
		})
		pos++
	}
	return rv
}

func init() {
	registry.RegisterTokenizer(kagomeTokenizerName, func(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
		return NewKagomeTokenizer()
	})
}
