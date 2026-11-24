// Command indexgen generates the search index of PLATEAU specification
// documents: one Mokuji file (https://github.com/reearth/mokuji).
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/eukarya-inc/plateau-spec/cmd/indexgen/search"
	"github.com/eukarya-inc/plateau-spec/plateaudocsearch"
)

func main() {
	var (
		docsDir = flag.String("docs", "./docs", "Base docs directory containing standard/ and procedure/")
		output  = flag.String("output", "./search-index.mokuji", "Output index file path")
		help    = flag.Bool("help", false, "Show help message")
	)

	flag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	var docs []plateaudocsearch.Document
	for _, docType := range []plateaudocsearch.DocType{plateaudocsearch.DocTypeStandard, plateaudocsearch.DocTypeProcedure} {
		dir := filepath.Join(*docsDir, string(docType))
		if _, err := os.Stat(filepath.Join(dir, "index.json")); err != nil {
			fmt.Printf("Skipping %s docs (no index.json found)\n", docType)
			continue
		}
		fmt.Printf("Processing %s docs from %s\n", docType, dir)
		collected, err := search.CollectDocuments(dir, docType)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s docs: %v\n", docType, err)
			os.Exit(1)
		}
		fmt.Printf("  Collected %d sections\n", len(collected))
		docs = append(docs, collected...)
	}
	if len(docs) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no documents found")
		os.Exit(1)
	}

	data, err := plateaudocsearch.Build(docs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building index: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*output, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing index: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully created %s (%d sections, %.2f MB)\n", *output, len(docs), float64(len(data))/(1024*1024))
}

func printUsage() {
	fmt.Println("PLATEAU Search Index Generator")
	fmt.Println()
	fmt.Println("Generates the full-text search index of PLATEAU specification documents as one Mokuji file.")
	fmt.Println("Every section listed in index.json is indexed with its own text, read from json/.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  indexgen -docs ./docs -output ./search-index.mokuji")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -docs string")
	fmt.Println("        Base docs directory containing standard/ and procedure/ (default: ./docs)")
	fmt.Println("  -output string")
	fmt.Println("        Output index file path (default: ./search-index.mokuji)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
}
