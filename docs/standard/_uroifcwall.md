# 4.2.3.5.27 uro:IfcWall

## 4.2.3.5.27 uro:IfcWall

**表 4-84**

| 型の定義 | IFCで記述された壁の属性。厚さが不均一な壁等、特殊な壁を表す。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcBuildingElement | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:globalId     [ IfcRoot ] | xs::string [0..1] | そのオブジェクトが単一に識別できる唯一な識別子。22桁の文字列により表現する。 |
| uro:name     [ IfcRoot ] | xs::string [0..1] | オブジェクトの名称。 |
| uro:description     [ IfcRoot ] | xs::string [0..1] | オブジェクトの文字情報による追加説明。 |
| uro:objectType     [ IfcObject ] | xs::string [0..1] | オブジェクトの特定の型を示す。 |
| uro:tag     [ IfcElement ] | xs::string [0..1] | オブジェクトのシリアルナンバー、ポジションナンバーなどの識別番号。 |
| uro:elementType     [ IfcBuildingElement ] | gml::CodeType [0..1] | 建築物の部材の区分。コードリスト（IfcBuildingElement_elementType.xml）から選択する。 uro:elementTypeの値により、uro:predefinedType以降の使用可能な属性が異なる。 |
| uro:predefinedType     [ IfcBuildingElement ] | gml::CodeType [0..1] | 定義済み型に基づく区分。コードリスト（IfcBuildingElement_predefinedType.xml）から選択する。 uro:elementTypeが、Covering, Railing, Slabに区分される場合にこの属性を使用できる。 |
| uro:shapeType     [ IfcBuildingElement ] | gml::CodeType [0..1] | 形状の区分。コードリスト（IfcBuildingElement_shapeType.xml）から選択する。 uro:elementTypeがRamp, Stairに区分される場合にこの属性を使用できる。 |
| uro:numberOfRiser     [ IfcBuildingElement ] | xs::integer [0..1] | 蹴上数。 uro:elementTypeがStairFlightに区分される場合にこの属性を使用できる。 |
| uro:numberOfTreads     [ IfcBuildingElement ] | xs::integer [0..1] | 踏面数。 uro:elementTypeがStairFlightに区分される場合にこの属性を使用できる。 |
| uro:riserHeight     [ IfcBuildingElement ] | gml::LengthType [0..1] | 蹴上の高さ。単位はmとする。 uro:elementTypeがStairFlightに区分される場合にこの属性を使用できる。 |
| uro:treadLength     [ IfcBuildingElement ] | gml::LengthType [0..1] | 踏面の奥行の長さ。単位はmとする。 uro:elementTypeがStairFlightに区分される場合にこの属性を使用できる。 |
| uro:operationType     [ IfcBuildingElement ] | IfcTransportElementTypeEnum [0..1] | 輸送設備の区分。 uro:elementTypeがTransportElementに区分される場合にこの属性を使用できる。 |
| uro:capacityByWeight     [ IfcBuildingElement ] | gml::MeasureType [0..1] | 許容積載量。単位はkg。 uro:elementTypeがTransportElementに区分される場合にこの属性を使用できる。 |
| uro:capacityByNumber     [ IfcBuildingElement ] | xs::integer [0..1] | 許容定員数。単位は人。 uro:elementTypeがTransportElementに区分される場合にこの属性を使用できる。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:nominalLength     [ IfcWall ] | gml::LengthType [0..1] | 壁の中心線に沿った長さ。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:nominalWidth     [ IfcWall ] | gml::LengthType [0..1] | 壁中心線に垂直に計測した壁の厚さ。壁中心線に沿って厚さが一定の場合のみ。単位はm。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:nominalHeight     [ IfcWall ] | gml::LengthType [0..1] | 壁の高さ。壁中心線に沿って高さが一定の場合のみ。単位はm。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:grossFootPrintArea     [ IfcWall ] | gml::MeasureType [0..1] | 平面図上に投影した壁の形状の面積。壁のへこみなどを考慮しない。単位はm。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:netFootPrintArea     [ IfcWall ] | gml::MeasureType [0..1] | 平面図上に投影した壁の形状の面積。壁のへこみなどを考慮する。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:grossSideArea     [ IfcWall ] | gml::MeasureType [0..1] | 立面ビューによるカーテンウォールの面積。カーテンウォールに対する変形を考慮しない。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:netSideArea     [ IfcWall ] | gml::MeasureType [0..1] | 立面ビューによるカーテンウォールの面積。カーテンウォールに対する変形を考慮する。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:grossSideAreaLeft     [ IfcWall ] | gml::MeasureType [0..1] | Wall path の方向から見て左側の側面面積。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:netSideAreaLeft     [ IfcWall ] | gml::MeasureType [0..1] | Wall path の方向から見て左側の正味側面面積。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:grossSideAreaRight     [ IfcWall ] | gml::MeasureType [0..1] | Wall path の方向から見て右側の側面面積。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:netSideAreaRight     [ IfcWall ] | gml::MeasureType [0..1] | Wall path の方向から見て右側の正味側面面積。単位はm2。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:grossVolume     [ IfcWall ] | gml::MeasureType [0..1] | スラブの体積。開口、へこみなどを考慮しない。単位はm3。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |
| uro:netVolume     [ IfcWall ] | gml::MeasureType [0..1] | スラブの体積。開口、へこみなどを考慮する。単位はm3。uro:elementTypeがWall又はWallStandardCaseに区分される場合にこの属性を使用できる。 |

