# 4.10.3.7.4 urf:UrbanRapidTransitRailroadAttribute

## 4.10.3.7.4 urf:UrbanRapidTransitRailroadAttribute

**表 4-383**

| 型の定義 | 都市計画法第11条第1項第1号に定める都市高速鉄道について定めるべき事項。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:structureType     [ UrbanRapidTransitRailroadAttribute ] | gml::CodeType [0..1] | 都市計画法施行令第6条第1項第4号及び都市計画法施行規則第7条第1項第6号で定められる構造(鉄道構造)。コードリスト（TrafficFacility_trafficFacilityStructureType.xml）より選択する。 |
| urf:crossType     [ UrbanRapidTransitRailroadAttribute ] | gml::CodeType [0..1] | 都市計画法施行令第6条第1項第4号及び都市計画法施行規則第7条第1項第6号で定められる構造（鉄道構造が地表式のみ）。コードリスト（ TrafficFacility_trafficFacilityCrossingType.xml）より選択する。 |
| urf:structuralDetails     [ UrbanRapidTransitRailroadAttribute ] | StructureDetails [0..*] | 鉄道の構造の内訳。 |

