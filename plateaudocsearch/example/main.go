package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eukarya-inc/plateau-spec/plateaudocsearch"
)

func main() {
	ctx := context.Background()

	fmt.Println("=== Creating Search Client ===")
	client := plateaudocsearch.New()
	defer func() { _ = client.Close() }()

	fmt.Println("=== Initializing (downloading index if needed) ===")
	result, err := client.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if result.Downloaded {
		fmt.Printf("Downloaded search index in %v\n", result.DownloadTime)
	} else {
		fmt.Println("Using cached search index")
	}

	// 標準仕様書を検索
	fmt.Println("\n=== Search (standard): 建築物 ===")
	results, err := client.Search(ctx, plateaudocsearch.DocTypeStandard, "建築物")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d results:\n", len(results))
	for i, r := range results {
		fmt.Printf("  %d. [%.2f] %s (%s)\n", i+1, r.Score, r.Title, r.Path)
		if len(r.Snippets) > 0 {
			fmt.Printf("     Snippet: %s\n", truncate(r.Snippets[0], 100))
		}
	}

	// 作業手順書を検索（件数指定）
	fmt.Println("\n=== Search (procedure): 品質 (limit=5) ===")
	results, err = client.Search(ctx, plateaudocsearch.DocTypeProcedure, "品質", plateaudocsearch.WithLimit(5))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d results:\n", len(results))
	for i, r := range results {
		fmt.Printf("  %d. [%.2f] %s (%s)\n", i+1, r.Score, r.Title, r.Path)
	}

	// 全文書タイプで検索（DocTypeAll または ""）
	fmt.Println("\n=== Search All: LOD ===")
	results, err = client.Search(ctx, "", "LOD", plateaudocsearch.WithLimit(5))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d results:\n", len(results))
	for i, r := range results {
		fmt.Printf("  %d. [%.2f] [%s] %s\n", i+1, r.Score, r.DocType, r.Title)
	}

	fmt.Println("\nAll tests passed!")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
