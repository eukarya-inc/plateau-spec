# 4.13.3.1.18 uro:Height

## 4.13.3.1.18 uro:Height

**表 4-635**

| 型の定義 | 構造物の指定された2点間における高さを表すデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:highReference     [ Height ] | gml::CodeType [1..1] | 高い方の位置。コードリスト（Elevation_elevationReference.xml）より選択する。 |
| uro:lowReference     [ Height ] | gml::CodeType [1..1] | 低い方の位置。コードリスト（Elevation_elevationReference.xml）より選択する。 |
| uro:status     [ Height ] | HightStatusValue [1..1] | 高さは計測した値か推定した値かの別。 |
| uro:value     [ Height ] | gml::LengthType [1..1] | 指定された2点間の高さ。単位はmとする。 |

