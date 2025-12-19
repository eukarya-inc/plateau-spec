# 4.7.3.2.3 uro:WaterwayDetailAttribute

## 4.7.3.2.3 uro:WaterwayDetailAttribute

**表 4-290**

| 型の定義 | 航路の詳細な情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:routeId     [ WaterwayDetailAttribute ] | xs::string [1..1] | 航路の番号及び航路に与えられた管理番号。 |
| uro:routeDirection     [ WaterwayDetailAttribute ] | gml::CodeType [0..1] | 進行方向。コードリスト（WaterwayDetailAttribute_routeDirection.xml）より選択する。 |
| uro:minimumWidth     [ WaterwayDetailAttribute ] | gml::LengthType [0..1] | 最小の幅員。単位はm（uom=”m”）とする。幅員が固定値である場合は、uro:minimumWidthを使用する。 |
| uro:maximumWidth     [ WaterwayDetailAttribute ] | gml::LengthType [0..1] | 最大の幅員。単位はm（uom=”m”）とする。幅員が固定値である場合は、uro:minimumWidthを使用し、本属性は使用しない。 |
| uro:length     [ WaterwayDetailAttribute ] | gml::LengthType [0..1] | 航路の延長。単位はkm（uom=”km”）とする。 |
| uro:navigation     [ WaterwayDetailAttribute ] | xs::string [0..1] | 航法。 |
| uro:plannedDepth     [ WaterwayDetailAttribute ] | gml::LengthType [0..1] | 計画水深。単位はm（uom=”m”）とする。 |
| uro:speedLimit     [ WaterwayDetailAttribute ] | gml::MeasureType [0..1] | 速力制限。単位はkt（uom=”kt”）とする。 |
| uro:targetShipType     [ WaterwayDetailAttribute ] | xs::string [0..*] | 対象船型。 |

