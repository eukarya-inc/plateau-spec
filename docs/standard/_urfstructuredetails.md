# 4.10.3.7.5 urf:StructureDetails

## 4.10.3.7.5 urf:StructureDetails

**表 4-384**

| 型の定義 | 道路及び都市高速鉄道の構造を区間ごとに記述するために使用する型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:startLocation     [ StructureDetails ] | xs::string [1..1] | 起点の位置。 |
| urf:endLocation     [ StructureDetails ] | xs::string [1..1] | 終点の位置。 |
| urf:viaLocations     [ StructureDetails ] | xs::string [0..1] | 起点から終点までに存在する主な地名。 |
| urf:length     [ StructureDetails ] | gml::LengthType [0..1] | 区間の長さ。単位はm（uom=”m”）とする。 |
| urf:structureType     [ StructureDetails ] | gml::CodeType [0..1] | 道路又は鉄道の構造の形式。コードリスト（TrafficFacility_trafficFacilityStructureType.xml）より選択する。 |
| urf:minimumWidth     [ StructureDetails ] | gml::LengthType [0..1] | 最小の幅員。単位はm（uom=”m”）とする。 |
| urf:maximumWidth     [ StructureDetails ] | gml::LengthType [0..1] | 最大の幅員。単位はm（uom=”m”）とする。 |
| urf:standardWidth     [ StructureDetails ] | gml::LengthType [0..1] | 標準的な幅員。単位はm（uom=”m”）とする。 |
| urf:crossType     [ StructureDetails ] | gml::CodeType [0..1] | 交差の種別。コードリスト（TrafficFacility_trafficFacilityCrossingType.xml）より選択する。 |

