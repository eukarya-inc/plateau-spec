# 4.15.3.1.13 uro:OffsetDepth

## 4.15.3.1.13 uro:OffsetDepth

**表 4-709**

| 型の定義 | 地下埋設物の道路縁等からの距離（オフセット）及び土被りの深さを表すデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:pos     [ OffsetDepth ] | gml::DirectPosition [0..1] | オフセット及び土被りを計測した位置。 |
| uro:offset     [ OffsetDepth ] | gml::LengthType [0..1] | オフセット量。単位はmとする。 |
| uro:depth     [ OffsetDepth ] | gml::LengthType [0..1] | 土被り量。単位は単位はmとする。 |
| uro:minDepth     [ OffsetDepth ] | gml::LengthType [0..1] | 最小の土被り量。土被り量が範囲で示されている場合に使用する。単位はmとする。 |
| uro:maxDepth     [ OffsetDepth ] | gml::LengthType [0..1] | 最大の土被り量。土被り量が範囲で示されている場合に使用する。単位はmとする。 |

