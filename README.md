# plateau-spec

PLATEAU標準製品仕様書・標準作業手順書をMarkdown形式に変換したドキュメント・ライブラリ・ツールのリポジトリです。

## ドキュメント

[PLATEAU](https://www.mlit.go.jp/plateau/)の仕様書は可読性の低いHTMLで提供されているため、このリポジトリではそれらをMarkdown形式に変換して提供しています。

- [3D都市モデル標準製品仕様書](docs/standard/index.md) - [原本](https://www.mlit.go.jp/plateaudocument/)
- [3D都市モデル標準作業手順書](docs/procedure/index.md) - [原本](https://www.mlit.go.jp/plateaudocument02/)

> **注意**: このドキュメントは上記出典から自動変換されたものです。変換処理により原本と異なる場合があります。正確な情報は原本を参照してください。

## Goライブラリ

### ドキュメント取得ライブラリ (plateaudoc)

GitHub rawコンテンツからPLATEAU仕様書を直接取得するGoライブラリです。`fs.FS`インターフェースも実装しています。

#### インストール

```bash
go get github.com/eukarya-inc/plateau-spec/plateaudoc
```

#### 使い方

```go
package main

import (
    "context"
    "fmt"
    "io/fs"
    "log"

    "github.com/eukarya-inc/plateau-spec/plateaudoc"
)

func main() {
    ctx := context.Background()
    client := plateaudoc.New()

    // インデックス取得
    index, err := client.GetIndex(ctx, "standard")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Title: %s\n", index.Title)
    fmt.Printf("Chapters: %d\n", len(index.Chapters))

    // Markdown取得
    md, err := client.GetMarkdown(ctx, "standard", "toc1")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(md)

    // JSON取得
    jsonData, err := client.GetJSON(ctx, "standard", "toc1")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(jsonData)

    // 画像取得
    imgResult, err := client.GetImage(ctx, "standard", "toc4_img051.png")
    if err != nil {
        log.Fatal(err)
    }
    defer imgResult.Body.Close()
    // imgResult.Body から読み取り

    // fs.FS インターフェースを使う
    docsFS := plateaudoc.NewFS()
    data, err := fs.ReadFile(docsFS, "standard/toc1.md")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))
}
```

#### API

| メソッド | 説明 |
|---------|------|
| `GetIndex(ctx, docType)` | index.jsonを取得 |
| `GetIndexMarkdown(ctx, docType)` | index.mdを取得 |
| `GetMarkdown(ctx, docType, path)` | Markdownファイルを取得 |
| `GetJSON(ctx, docType, path)` | JSONファイルを取得 |
| `GetImage(ctx, docType, filename)` | 画像ファイルを取得 |

##### オプション

```go
// カスタムベースURL
client := plateaudoc.New(
    plateaudoc.WithBaseURL("https://custom-url/docs"),
)

// カスタムHTTPクライアント
client := plateaudoc.New(
    plateaudoc.WithHTTPClient(&http.Client{Timeout: time.Minute}),
)

// gzip圧縮を無効化
client := plateaudoc.New(
    plateaudoc.WithoutGzip(),
)
```

##### fs.FS インターフェース

```go
docsFS := plateaudoc.NewFS()

// 標準のfs関数が使える
data, _ := fs.ReadFile(docsFS, "standard/toc1.md")
file, _ := docsFS.Open("procedure/toc2.md")
```

### 全文検索ライブラリ (plateaudocearch)

GitHub Releaseから検索インデックスをダウンロードし、全文検索を提供するGoライブラリです。

#### インストール

```bash
go get github.com/eukarya-inc/plateau-spec/plateaudocearch
```

#### 使い方

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/eukarya-inc/plateau-spec/plateaudocearch"
)

func main() {
    ctx := context.Background()

    // クライアントを作成
    client := plateaudocearch.New()
    defer client.Close()

    // 初期化（インデックスをダウンロード）
    result, err := client.Init(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if result.Downloaded {
        fmt.Printf("インデックスをダウンロードしました (%v)\n", result.DownloadTime)
    } else {
        fmt.Println("キャッシュされたインデックスを使用します")
    }

    // 標準製品仕様書を検索
    results, err := client.Search(ctx, plateaudocearch.DocTypeStandard, "建築物")
    if err != nil {
        log.Fatal(err)
    }

    for _, r := range results {
        fmt.Printf("[%.2f] %s (%s)\n", r.Score, r.Title, r.Path)
    }
}
```

#### API

##### Client オプション

```go
// カスタムリリースURL
client := plateaudocearch.New(
    plateaudocearch.WithReleaseURL("https://custom-url/search-index.tar.gz"),
)

// カスタムキャッシュディレクトリ
client := plateaudocearch.New(
    plateaudocearch.WithCacheDir("/path/to/cache"),
)

// カスタムHTTPクライアント
client := plateaudocearch.New(
    plateaudocearch.WithHTTPClient(&http.Client{Timeout: 10 * time.Minute}),
)
```

##### 検索オプション

```go
// 検索結果の件数を指定
results, err := client.Search(ctx, plateaudocearch.DocTypeStandard, "建築物",
    plateaudocearch.WithLimit(20),
)

// 全ドキュメントタイプを検索
results, err := client.Search(ctx, "", "LOD")
// または
results, err := client.Search(ctx, plateaudocearch.DocTypeAll, "LOD")
```

##### ドキュメントタイプ

| 定数 | 値 | 説明 |
|-----|---|------|
| `DocTypeStandard` | `"standard"` | 標準製品仕様書 |
| `DocTypeProcedure` | `"procedure"` | 標準作業手順書 |
| `DocTypeAll` | `""` | 全ドキュメント |

##### SearchResult

```go
type SearchResult struct {
    ID       string   // ドキュメントID
    DocType  DocType  // ドキュメントタイプ
    Path     string   // ファイルパス
    Title    string   // タイトル
    Score    float64  // スコア
    Snippets []string // マッチした箇所のスニペット
}
```

## CLIツール

### ドキュメント生成ツール (specgen)

PLATEAU仕様書を自動でダウンロード・変換するCLIツールです。

#### ビルド

```bash
cd cmd
go build -o specgen .
```

#### 使い方

```bash
# 標準製品仕様書を生成
./specgen -type standard -output ./docs/standard

# 標準作業手順書を生成
./specgen -type procedure -output ./docs/procedure
```

#### オプション

| オプション | 説明 | デフォルト |
|-----------|------|-----------|
| `-type` | ドキュメントタイプ: `standard`（標準製品仕様書）または `procedure`（標準作業手順書） | `standard` |
| `-output` | 出力ディレクトリ（必須） | - |
| `-path` | 生成するページをカンマ区切りで指定（例: `toc1,toc2_03`）。指定しない場合は全ページを生成 | 全ページ |
| `-help` | ヘルプを表示 | - |

#### 特定ページのみ生成

`-path`オプションを使うと、特定のページだけを再生成できます：

```bash
# toc1とその子ページのみ生成
./specgen -type standard -output ./docs/standard -path toc1

# 複数ページを指定
./specgen -type standard -output ./docs/standard -path toc1,toc2,toc3
```

#### 出力構造

```
{output}/
├── index.json      # ドキュメントのアウトライン
├── index.md        # 目次（Markdown）
├── toc1.md         # 章ファイル
├── toc1_01.md      # 節ファイル
├── toc1_01_01.md   # 項ファイル
├── json/           # 各ページの生JSON
│   ├── toc1.json
│   └── ...
└── images/         # 抽出された画像
    ├── toc1_img001.png
    └── ...
```

### インデックス生成ツール (indexgen)

Markdown形式のドキュメントから全文検索用インデックスを生成するCLIツールです。

#### ビルド

```bash
cd cmd
go build -o indexgen ./indexgen
```

#### 使い方

```bash
# インデックスを生成
./indexgen -docs ./docs -output ./search-index.tar.gz
```

#### オプション

| オプション | 説明 | デフォルト |
|-----------|------|-----------|
| `-docs` | ドキュメントの基準ディレクトリ | `./docs` |
| `-output` | 出力ファイルパス (.tar.gz) | `./search-index.tar.gz` |

#### 出力

生成される `search-index.tar.gz` には以下が含まれます：

- `standard.bleve/` - 標準製品仕様書の検索インデックス
- `procedure.bleve/` - 標準作業手順書の検索インデックス

## GitHub Actions

検索インデックスは以下のワークフローで生成・公開されます：

- **generate-docs.yml**: ドキュメント生成
- **generate-search-index.yml**: インデックスを生成

生成されたインデックスは [GitHub Release](https://github.com/eukarya-inc/plateau-spec/releases/tag/search-index-latest) からダウンロードできます。

## ライセンス

変換ツールのソースコードはMITライセンスです。

ドキュメントの内容は国土交通省の著作物であり、原本の利用規約に従ってください。
