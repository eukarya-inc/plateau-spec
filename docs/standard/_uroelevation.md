# 4.13.3.1.17 uro:Elevation

## 4.13.3.1.17 uro:Elevation

**表 4-634**

| 型の定義 | 構造物の指定された位置における標高を表すデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:elevationReference     [ Elevation ] | gml::CodeType [1..1] | 標高を計測する位置。コードリスト（Elevation_elevationReference.xml）より選択する。 |
| uro:elevationValue     [ Elevation ] | gml::DirectPosition [1..1] | 指定された位置での標高。東京湾平均海面からの高さとする。単位はmとする。 |

