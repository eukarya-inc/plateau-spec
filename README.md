# plateau-spec

PLATEAU標準製品仕様書・標準作業手順書をMarkdown形式に変換したドキュメントリポジトリです。

## ドキュメント

- [3D都市モデル標準製品仕様書](docs/standard/index.md) - [原本](https://www.mlit.go.jp/plateaudocument/)
- [3D都市モデル標準作業手順書](docs/procedure/index.md) - [原本](https://www.mlit.go.jp/plateaudocument02/)

> **注意**: このドキュメントは上記出典から自動変換されたものです。変換処理により原本と異なる場合があります。正確な情報は原本を参照してください。

## 概要

[PLATEAU](https://www.mlit.go.jp/plateau/)の仕様書は読みづらいHTMLで提供されているため、このリポジトリではそれらをMarkdown形式に変換して提供しています。

## CLIツール

このリポジトリには、PLATEAU仕様書を自動でダウンロード・変換するCLIツールが含まれています。

### ビルド

```bash
cd cmd
go build -o specgen .
```

### 使い方

```bash
# 標準製品仕様書を生成
./specgen -type standard -output ./docs/standard

# 標準作業手順書を生成
./specgen -type procedure -output ./docs/procedure
```

### オプション

| オプション | 説明 | デフォルト |
|-----------|------|-----------|
| `-type` | ドキュメントタイプ: `standard`（標準製品仕様書）または `procedure`（標準作業手順書） | `standard` |
| `-output` | 出力ディレクトリ（必須） | - |
| `-path` | 生成するページをカンマ区切りで指定（例: `toc1,toc2_03`）。指定しない場合は全ページを生成 | 全ページ |
| `-help` | ヘルプを表示 | - |

### 特定ページのみ生成

`-path`オプションを使うと、特定のページだけを再生成できます：

```bash
# toc1とその子ページのみ生成
./specgen -type standard -output ./docs/standard -path toc1

# 複数ページを指定
./specgen -type standard -output ./docs/standard -path toc1,toc2,toc3
```

### 出力構造

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

## ライセンス

変換ツールのソースコードはMITライセンスです。

ドキュメントの内容は国土交通省の著作物であり、原本の利用規約に従ってください。
