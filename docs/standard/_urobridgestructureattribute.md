# 4.11.3.2.1 uro:BridgeStructureAttribute

## 4.11.3.2.1 uro:BridgeStructureAttribute

**表 4-522**

| 型の定義 | 橋梁の構造に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:material     [ BridgeStructureAttribute ] | gml::CodeType [0..1] | 橋梁の主たる材質。コードリスト（BridgeStructureAttribute_material.xml）より選択する。 |
| uro:bridgeType     [ BridgeStructureAttribute ] | gml::CodeType [0..1] | 橋梁の種類。コードリスト（BridgeStructureAttribute_bridgeType.xml）より選択する。 |
| uro:length     [ BridgeStructureAttribute ] | gml::LengthType [0..1] | 橋梁の長さ。単位はmとする。 |
| uro:width     [ BridgeStructureAttribute ] | gml::LengthType [0..1] | 橋梁の幅員。単位はmとする。 |
| uro:area     [ BridgeStructureAttribute ] | gml::MeasureType [0..1] | 橋梁の上部工の面積。単位はm2とする。 |
| uro:weightRestriction     [ BridgeStructureAttribute ] | gml::MeasureType [0..1] | 橋梁の荷重制限。単位はtとする。 |
| uro:heightRestriction     [ BridgeStructureAttribute ] | gml::LengthType [0..1] | 橋梁の高さ制限。単位はmとする。 |
| uro:widthRestriction     [ BridgeStructureAttribute ] | gml::LengthType [0..1] | 橋梁の幅制限。単位はmとする。 |
| uro:underGirderHeight     [ BridgeStructureAttribute ] | gml::LengthType [0..1] | 橋梁の桁下の高さ制限。単位はmとする。 |
| uro:slopeType     [ BridgeStructureAttribute ] | gml::CodeType [0..1] | 橋梁が橋側歩道橋、横断歩道橋又はペデストリアンデッキの場合の、昇降形式。コードリスト（ConstructionStructureAttribute_slopeType.xml）より選択する。 |
| uro:escalator     [ BridgeStructureAttribute ] | xs::boolean [0..1] | 橋梁が橋側歩道橋、横断歩道橋又はペデストリアンデッキの場合の、エスカレータの有無。 1：有 0：無 |

