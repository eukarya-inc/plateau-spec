// Command spec-generator generates PLATEAU specification documents as Markdown files.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/eukarya-inc/plateau-spec/cmd/specgen"
)

func main() {
	var (
		docType = flag.String("type", "standard", "Document type: 'standard' or 'procedure'")
		output  = flag.String("output", "", "Output directory (required)")
		paths   = flag.String("path", "", "Comma-separated list of paths to generate (e.g., 'toc1,toc2_03'). If empty, generates all.")
		help    = flag.Bool("help", false, "Show help message")
	)

	flag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	if *output == "" {
		fmt.Fprintln(os.Stderr, "Error: -output is required")
		printUsage()
		os.Exit(1)
	}

	if *docType != "standard" && *docType != "procedure" {
		fmt.Fprintf(os.Stderr, "Error: -type must be 'standard' or 'procedure', got '%s'\n", *docType)
		os.Exit(1)
	}

	// Parse paths filter
	var pathFilter []string
	if *paths != "" {
		for _, p := range strings.Split(*paths, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				pathFilter = append(pathFilter, p)
			}
		}
	}

	if len(pathFilter) > 0 {
		fmt.Printf("Generating %s specification documents to %s (filtered: %s)\n", *docType, *output, strings.Join(pathFilter, ", "))
	} else {
		fmt.Printf("Generating %s specification documents to %s\n", *docType, *output)
	}

	generator := specgen.NewGenerator(*docType, *output, pathFilter)
	if err := generator.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}

func printUsage() {
	fmt.Println("PLATEAU Specification Document Generator")
	fmt.Println()
	fmt.Println("Generates PLATEAU specification documents as Markdown files with images.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  spec-generator -type standard -output ./docs/spec/standard")
	fmt.Println("  spec-generator -type procedure -output ./docs/spec/procedure")
	fmt.Println("  spec-generator -type standard -output ./docs -path toc1,toc2")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -type string")
	fmt.Println("        Document type: 'standard' (3D都市モデル標準製品仕様書)")
	fmt.Println("                   or 'procedure' (3D都市モデル標準作業手順書)")
	fmt.Println("        (default: standard)")
	fmt.Println("  -output string")
	fmt.Println("        Output directory (required)")
	fmt.Println("  -path string")
	fmt.Println("        Comma-separated list of paths to generate (e.g., 'toc1,toc2_03')")
	fmt.Println("        If not specified, generates all pages")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Output structure:")
	fmt.Println("  {output}/")
	fmt.Println("    ├── index.json      # Document outline")
	fmt.Println("    ├── toc1.md         # Chapter files")
	fmt.Println("    ├── toc1_01.md      # Section files")
	fmt.Println("    ├── toc1_01_01.md   # Subsection files")
	fmt.Println("    ├── json/           # Raw JSON for each page")
	fmt.Println("    │   ├── toc1.json")
	fmt.Println("    │   └── ...")
	fmt.Println("    └── images/         # Extracted images")
	fmt.Println("        ├── toc1_img001.png")
	fmt.Println("        └── ...")
}
