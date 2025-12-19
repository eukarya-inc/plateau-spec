# 4.13.3.1.12 uro:DamAttribute

## 4.13.3.1.12 uro:DamAttribute

**表 4-629**

| 型の定義 | ダムの規模に関する基本的な情報を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:ConstructionStructureAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:structureType     [ ConstructionStructureAttribute ] | gml::CodeType [0..1] | 構造物の構造種別。コードリスト（ConstructionStructureAttribute_structureType.xml）から選択する。コードリストに無い場合は文字列により記述する。 |
| uro:length     [ ConstructionStructureAttribute ] | gml::LengthType [0..1] | 構造物の総延長。単位はmとする。ダムの場合は使用しない。 |
| uro:width     [ ConstructionStructureAttribute ] | gml::LengthType [0..1] | 構造物の幅。単位はmとする。堤防の場合は使用しない。 |
| uro:depth     [ ConstructionStructureAttribute ] | gml::LengthType [0..1] | 構造物が設置された位置の水深。単位はmとする。堤防の場合は使用しない。 |
| uro:volume     [ ConstructionStructureAttribute ] | gml::MeasureType [0..1] | 構造物の体積。単位は千m3とする。堤防の場合は使用しない。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:damCode     [ DamAttribute ] | gml::CodeType [0..1] | ダム年鑑の「全国ダム施設現況」の「水系別ダム一覧表」に示す番号。 |
| uro:totalWaterStorage     [ DamAttribute ] | gml::MeasureType [0..1] | 総貯水量。単位はm3とする。 |

