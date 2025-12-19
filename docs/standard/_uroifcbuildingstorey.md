# 4.24.3.2.1 uro:IfcBuildingStorey

## 4.24.3.2.1 uro:IfcBuildingStorey

**表 4-853**

| 型の定義 | IFCで記述された建築物の階数の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcSpatialStructureElement | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:globalId     [ IfcRoot ] | xs::string [0..1] | そのオブジェクトが単一に識別できる唯一な識別子。22桁の文字列により表現する。 |
| uro:name     [ IfcRoot ] | xs::string [0..1] | オブジェクトの名称。 |
| uro:description     [ IfcRoot ] | xs::string [0..1] | オブジェクトの文字情報による追加説明。 |
| uro:objectType     [ IfcObject ] | xs::string [0..1] | オブジェクトの特定の型を示す。 |
| uro:longName     [ IfcSpatialStructureElement ] | xs::string [0..1] | 人が識別可能な名称。uro:IfcBuildingの場合は、敷地を識別するための名称とする。 |
| uro:compositionType     [ IfcSpatialStructureElement ] | IfcElementCompositionEnum [0..1] | 空間構成の区分。単一であればELEMENT を設定。複数から構成される場合はCOMPLEX を設定。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:elevation     [ IfcBuildingStorey ] | gml::LengthType [0..1] | 建物階の建物の基準高さからの相対的高さ。単位はm。 |

