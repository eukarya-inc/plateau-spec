# 4.2.3.5.23 uro:IfcSite

## 4.2.3.5.23 uro:IfcSite

**表 4-80**

| 型の定義 | IFCで記述されたプロジェクトの敷地に適用される属性の集まり。 | ← |
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
| uro:refLongitude     [ IfcSite ] | xs::double [0..1] | 敷地の参照ポイントの経度。 |
| uro:refLatitude     [ IfcSite ] | xs::double [0..1] | 敷地の参照ポイントの緯度。 |
| uro:refElevation     [ IfcSite ] | gml::LengthType [0..1] | 敷地の参照ポイントの標高。 |
| uro:landTitleNumber     [ IfcSite ] | xs::string [0..1] | 土地登記に関連する識別情報。 |
| uro:siteAddress     [ IfcSite ] | Address [0..1] | 郵便住所。 |

