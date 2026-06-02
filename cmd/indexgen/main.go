// Command indexgen generates Bleve search indexes from PLATEAU specification documents.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/eukarya-inc/plateau-spec/cmd/indexgen/search"
)

func main() {
	var (
		docsDir = flag.String("docs", "./docs", "Base docs directory containing standard/ and procedure/")
		output  = flag.String("output", "./search-index.tar.gz", "Output .tar.gz file path")
		help    = flag.Bool("help", false, "Show help message")
	)

	flag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	tmpDir, err := os.MkdirTemp("", "search-index-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temp directory: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Process standard docs
	standardDir := filepath.Join(*docsDir, "standard")
	if _, err := os.Stat(filepath.Join(standardDir, "index.json")); err == nil {
		fmt.Printf("Processing standard docs from %s\n", standardDir)
		if err := processDocType(standardDir, filepath.Join(tmpDir, "standard.bleve")); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing standard docs: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Skipping standard docs (no index.json found)\n")
	}

	// Process procedure docs
	procedureDir := filepath.Join(*docsDir, "procedure")
	if _, err := os.Stat(filepath.Join(procedureDir, "index.json")); err == nil {
		fmt.Printf("Processing procedure docs from %s\n", procedureDir)
		if err := processDocType(procedureDir, filepath.Join(tmpDir, "procedure.bleve")); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing procedure docs: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Skipping procedure docs (no index.json found)\n")
	}

	// Compress to tar.gz
	fmt.Printf("Compressing to %s\n", *output)
	if err := compressDir(tmpDir, *output); err != nil {
		fmt.Fprintf(os.Stderr, "Error compressing index: %v\n", err)
		os.Exit(1)
	}

	// Show result size
	info, err := os.Stat(*output)
	if err == nil {
		fmt.Printf("Successfully created %s (%.2f MB)\n", *output, float64(info.Size())/(1024*1024))
	}
}

func processDocType(docsDir, indexPath string) error {
	index, err := search.CreateIndex(indexPath)
	if err != nil {
		return fmt.Errorf("creating index: %w", err)
	}
	defer func() {
		if err := index.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing index: %v\n", err)
		}
	}()

	count, err := search.IndexFromJSON(index, docsDir)
	if err != nil {
		return fmt.Errorf("indexing: %w", err)
	}

	fmt.Printf("  Indexed %d documents\n", count)
	return nil
}

func compressDir(srcDir, destPath string) error {
	return search.CompressIndex(srcDir, destPath)
}

func printUsage() {
	fmt.Println("PLATEAU Search Index Generator")
	fmt.Println()
	fmt.Println("Generates Bleve full-text search indexes from PLATEAU specification documents.")
	fmt.Println("Reads index.json from each doc type directory and indexes only root-level chapters.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  indexgen -docs ./docs -output ./search-index.tar.gz")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -docs string")
	fmt.Println("        Base docs directory containing standard/ and procedure/ (default: ./docs)")
	fmt.Println("  -output string")
	fmt.Println("        Output .tar.gz file path (default: ./search-index.tar.gz)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Output structure (inside tar.gz):")
	fmt.Println("  standard.bleve/   # Search index for standard specification")
	fmt.Println("  procedure.bleve/  # Search index for procedure specification")
}
