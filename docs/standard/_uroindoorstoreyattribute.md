# 4.24.3.3.2 uro:IndoorStoreyAttribute

## 4.24.3.3.2 uro:IndoorStoreyAttribute

**表 4-856**

| 型の定義 | 階層に追加するナビゲーション用の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IndoorAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:source     [ IndoorAttribute ] | gml::CodeType [0..1] | 原典資料。コードリスト（Common_indoorSource.xml）から選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:category     [ IndoorStoreyAttribute ] | xs::boolean [0..1] | 屋内外区分。1：屋内0：屋外 |
| uro:ordinal     [ IndoorStoreyAttribute ] | xs::double [0..1] | 階層数。 |

