# 4.10.3.7.3 urf:UrbanRoadAttribute

## 4.10.3.7.3 urf:UrbanRoadAttribute

**表 4-382**

| 型の定義 | 都市計画法第11条第1項第1号に定める道路について定めるべき事項。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:routeTypeNumber     [ UrbanRoadAttribute ] | gml::CodeType [0..1] | 都市局長通達(昭和44年建設省都計発第102号)Ⅳ(2)③で定められる区分。コードリスト（UrbanRoadAttribute_routeTypeNumber.xml）より選択する。 |
| urf:routeSizeNumber     [ UrbanRoadAttribute ] | gml::CodeType [0..1] | 都市局長通達(昭和44年建設省都計発第102号)Ⅳ(2)③で定められる規模。コードリスト（UrbanRoadAttribute_routeSizeNumber.xml）より選択する。 |
| urf:routeSerialNumber     [ UrbanRoadAttribute ] | xs::string [0..1] | 都市局長通達(昭和44年建設省都計発第102号)Ⅳ(2)③で定められる一連番号。 |
| urf:roadType     [ UrbanRoadAttribute ] | gml::CodeType [0..1] | 都市計画法施行令第6条第1項第1号で定める道路の種別。コードリスト（UrbanRoadAttribute_roadType.xml）より選択する。 |
| urf:numberOfLanes     [ UrbanRoadAttribute ] | xs::integer [0..1] | 都市計画法施行令第6条第1項第1号で定める車線の数。 |
| urf:roadStructure     [ UrbanRoadAttribute ] | gml:StringOrRefType [0..1] | 都市計画法施行令第6条第1項第1号で定める道路の構造。 |
| urf:structureType     [ UrbanRoadAttribute ] | gml::CodeType [0..1] | 都市計画法施行令第6条第1項第1号及び都市計画法施行規則第7条第1項第2 号で定める構造種別。コードリスト（ TrafficFacility_trafficFacilityStructureType.xml）より選択する。 |
| urf:crossType     [ UrbanRoadAttribute ] | gml::CodeType [0..1] | 都市計画法施行令第6条第1項第1号及び都市計画法施行規則第7条第1項第2号で定める交差種別（道路構造が地表式のみ）。コードリスト（ TrafficFacility_trafficFacilityCrossingType.xml）より選択する。 |
| urf:trafficPlazas     [ UrbanRoadAttribute ] | gml::CodeType [0..1] | 都市計画法第11条第1項で定める施設における交通広場の有無。コードリスト（Common_availabilityType.xml）より選択する。 |
| urf:structuralDetails     [ UrbanRoadAttribute ] | StructureDetails [0..*] | 道路の構造の内訳。 |

