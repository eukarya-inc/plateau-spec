# 4.13.3.1.10 uro:ConstructionStructureAttribute

## 4.13.3.1.10 uro:ConstructionStructureAttribute

**表 4-627**

| 型の定義 | 構造物の規模に関する基本的な情報を記述するためのデータ型。構造物が護岸、床止め、堰、水門・閘門・陸閘、樋門・樋管、伏せ越し及び水制の場合に本データ型を用いて記述する。 構造物がダム及び堤防の場合は、これを継承するデータ型を用いて記述する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:structureType     [ ConstructionStructureAttribute ] | gml::CodeType [0..1] | 構造物の構造種別。コードリスト（ConstructionStructureAttribute_structureType.xml）から選択する。コードリストに無い場合は文字列により記述する。 |
| uro:length     [ ConstructionStructureAttribute ] | gml::LengthType [0..1] | 構造物の総延長。単位はmとする。ダムの場合は使用しない。 |
| uro:width     [ ConstructionStructureAttribute ] | gml::LengthType [0..1] | 構造物の幅。単位はmとする。堤防の場合は使用しない。 |
| uro:depth     [ ConstructionStructureAttribute ] | gml::LengthType [0..1] | 構造物が設置された位置の水深。単位はmとする。堤防の場合は使用しない。 |
| uro:volume     [ ConstructionStructureAttribute ] | gml::MeasureType [0..1] | 構造物の体積。単位は千m3とする。堤防の場合は使用しない。 |

