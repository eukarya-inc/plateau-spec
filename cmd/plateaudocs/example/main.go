package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"

	"github.com/eukarya-inc/plateau-spec/cmd/plateaudocs"
)

func main() {
	ctx := context.Background()
	client := plateaudocs.New()

	// インデックス取得
	fmt.Println("=== Index (standard) ===")
	index, err := client.GetIndex(ctx, "standard")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title: %s\n", index.Title)
	fmt.Printf("Generated: %s\n", index.GeneratedAt)
	fmt.Printf("Chapters: %d\n", len(index.Chapters))
	for _, ch := range index.Chapters {
		fmt.Printf("  - %s: %s\n", ch.Path, ch.Title)
	}

	// マークダウン取得
	fmt.Println("\n=== Markdown (standard/toc1) ===")
	md, err := client.GetMarkdown(ctx, "standard", "toc1")
	if err != nil {
		log.Fatal(err)
	}
	// 最初の500文字だけ表示
	if len(md) > 500 {
		fmt.Println(md[:500] + "...")
	} else {
		fmt.Println(md)
	}

	// JSON取得
	fmt.Println("\n=== JSON (standard/toc1) ===")
	jsonData, err := client.GetJSON(ctx, "standard", "toc1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Keys: ")
	for k := range jsonData {
		fmt.Printf("%s ", k)
	}
	fmt.Println()

	// 画像取得
	fmt.Println("\n=== Image (standard/toc4_img051.png) ===")
	imgResult, err := client.GetImage(ctx, "standard", "toc4_img051.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imgResult.Body.Close()

	fmt.Printf("Content-Length: %d bytes\n", imgResult.ContentLength)
	fmt.Printf("Content-Type: %s\n", imgResult.ContentType)

	// 画像サイズを確認
	imgData, err := io.ReadAll(imgResult.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actual size: %d bytes\n", len(imgData))

	// FS テスト
	fmt.Println("\n=== FS (fs.FS interface) ===")
	docsFS := plateaudocs.NewFS()

	// fs.Stat でサイズ取得
	file, err := docsFS.Open("standard/toc1.md")
	if err != nil {
		log.Fatal(err)
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File: %s, Size: %d bytes\n", info.Name(), info.Size())
	file.Close()

	// fs.ReadFile で読み取り
	data, err := fs.ReadFile(docsFS, "standard/toc1.md")
	if err != nil {
		log.Fatal(err)
	}
	if len(data) > 200 {
		fmt.Println(string(data[:200]) + "...")
	} else {
		fmt.Println(string(data))
	}

	fmt.Println("\n✅ All tests passed!")
}
