# 4.4.3.2.4 uro:RailwayRouteAttribute

## 4.4.3.2.4 uro:RailwayRouteAttribute

**表 4-199**

| 型の定義 | 鉄道路線の情報。路線ごとに作成する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:operatorType     [ RailwayRouteAttribute ] | gml::CodeType [1..1] | 鉄道事業者の区分。コードリスト（RailwayRouteAttribute_operatorType.xml）より選択する。 |
| uro:operator     [ RailwayRouteAttribute ] | xs::string [1..1] | 鉄道事業者名。 |
| uro:alternativeName     [ RailwayRouteAttribute ] | xs::string [0..*] | 路線名称（gml:name）以外に使用されている愛称等の別称。 |
| uro:railwayType     [ RailwayRouteAttribute ] | gml::CodeType [1..1] | 鉄道の区分。コードリスト（RailwayRouteAttribute_railwayType.xml）より選択する。 |
| uro:startStation     [ RailwayRouteAttribute ] | xs::string [1..1] | 鉄道路線の起点となる駅の名称。 |
| uro:endStation     [ RailwayRouteAttribute ] | xs::string [1..1] | 鉄道路線の終点となる駅の名称。 |

