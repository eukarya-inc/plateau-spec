package plateaudocsearch

import (
	"html"
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

const (
	// snippetContext is how many characters a snippet shows on each side of
	// the match it is cut around.
	snippetContext = 40
	// maxSnippets is how many snippets a result carries at most.
	maxSnippets = 3
)

// queryPieces is what a query is matched against text as: its runs of
// letters and digits, normalized as the text is. The bigram tokenizer splits
// at the same places, so a document matched by the index contains each piece
// as bigrams, and nearly always as a run.
func queryPieces(query string) [][]rune {
	var pieces [][]rune
	fields := strings.FieldsFunc(norm.NFKC.String(query), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for _, f := range fields {
		pieces = append(pieces, lowerRunes([]rune(f)))
	}
	return pieces
}

// lowerRunes lowercases rune by rune, so that an index into the result is an
// index into the input.
func lowerRunes(rs []rune) []rune {
	out := make([]rune, len(rs))
	for i, r := range rs {
		out[i] = unicode.ToLower(r)
	}
	return out
}

type span struct{ start, end int }

// snippets cuts up to maxSnippets fragments of text around where the query's
// pieces occur, with each occurrence wrapped in <mark>. The text is HTML
// escaped, as the fragments are meant to be shown.
func snippets(text string, pieces [][]rune) []string {
	if len(pieces) == 0 || text == "" {
		return nil
	}
	runes := []rune(text)
	lowered := lowerRunes(runes)

	var hits []span
	for _, p := range pieces {
		for i := 0; i+len(p) <= len(lowered); {
			if slices.Equal(lowered[i:i+len(p)], p) {
				hits = append(hits, span{i, i + len(p)})
				i += len(p)
			} else {
				i++
			}
		}
	}
	if len(hits) == 0 {
		return nil
	}
	slices.SortFunc(hits, func(a, b span) int { return a.start - b.start })
	hits = mergeSpans(hits)

	var out []string
	for i := 0; i < len(hits) && len(out) < maxSnippets; {
		start := max(hits[i].start-snippetContext, 0)
		end := min(hits[i].end+snippetContext, len(runes))
		// Every hit that fits in this window is marked in it.
		j := i
		for j < len(hits) && hits[j].end <= end {
			j++
		}

		var sb strings.Builder
		if start > 0 {
			sb.WriteString("…")
		}
		at := start
		for _, h := range hits[i:j] {
			sb.WriteString(html.EscapeString(string(runes[at:h.start])))
			sb.WriteString("<mark>")
			sb.WriteString(html.EscapeString(string(runes[h.start:h.end])))
			sb.WriteString("</mark>")
			at = h.end
		}
		sb.WriteString(html.EscapeString(string(runes[at:end])))
		if end < len(runes) {
			sb.WriteString("…")
		}
		out = append(out, sb.String())
		i = j
	}
	return out
}

// mergeSpans joins spans that overlap or touch; s is sorted by start.
func mergeSpans(s []span) []span {
	out := s[:1]
	for _, x := range s[1:] {
		last := &out[len(out)-1]
		if x.start <= last.end {
			last.end = max(last.end, x.end)
		} else {
			out = append(out, x)
		}
	}
	return out
}
