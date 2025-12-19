# 4.26.3.3.4 uro:PortTransportationFacility

## 4.26.3.3.4 uro:PortTransportationFacility

**表 4-877**

| 型の定義 | 港湾施設である臨港交通施設の属性を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:PortAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:facilityId     [ FacilityAttribute ] | xs::string [0..1] | 施設の用途や区分ごとに付与された管理番号。 |
| uro:portFacilityDetailsType     [ PortAttribute ] | gml::CodeType [1..1] | 施設の種類。コードリスト（PortAttribute_facilityDetailType.xml）より選択する。 |
| uro:portName     [ PortAttribute ] | xs::string [1..1] | 港湾名。 |
| uro:portStatus     [ PortAttribute ] | gml::CodeType [0..1] | 港格。コードリスト（PortAttribute_portStatus.xml）より選択する。 |
| uro:district     [ PortAttribute ] | xs::string [0..1] | 地区名。 |
| uro:grantType     [ PortAttribute ] | gml::CodeType [0..1] | 施設区分名。コードリスト（PortAttribute_grantType.xml）より選択する。 |
| uro:isDesignated     [ PortAttribute ] | xs::boolean [0..1] | 特定技術基準対象施設 0：対象外、1：対象 「技術基準対象施設であつて、外郭施設その他の非常災害により損壊した場合において船舶の交通に支障を及ぼすおそれのあるものとして国土交通省令で定めるもの」 （港湾法第56条の2の21）。 |
| uro:degradationLevel     [ PortAttribute ] | xs::integer [0..1] | 性能低下度。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:structureType     [ PortTransportationFacility ] | gml::CodeType [0..1] | 構造形式。コードリスト（PortAttribute_structureType.xml）より選択する。 |
| uro:startingPoint     [ PortTransportationFacility ] | xs::string [0..1] | 起終点。 |
| uro:length     [ PortTransportationFacility ] | gml::LengthType [0..1] | 規模_延長。 |
| uro:area     [ PortTransportationFacility ] | gml::MeasureType [0..1] | 規模_面積。 |
| uro:beddingWidth     [ PortTransportationFacility ] | gml::LengthType [0..1] | 規模_道路敷幅。 |
| uro:numberOfLanes     [ PortTransportationFacility ] | xs::integer [0..1] | 規模_車線数。 |
| uro:parkingLotCapacityOfBus     [ PortTransportationFacility ] | xs::integer [0..1] | 規模_駐車場収容台数_バス。 |
| uro:parkingLotCapacityOfCars     [ PortTransportationFacility ] | xs::integer [0..1] | 規模_駐車場収容台数_乗用車。 |
| uro:routeType     [ PortTransportationFacility ] | gml::CodeType [0..1] | 規模_単線・複線区分。コードリスト（PortTransportationFacility_routeType.xml）より選択する。 |
| uro:heightToDigit     [ PortTransportationFacility ] | gml::LengthType [0..1] | 規模_桁下高。 |
| uro:heightLimit     [ PortTransportationFacility ] | gml::LengthType [0..1] | 規模_制限高。 |
| uro:minimumWidth     [ PortTransportationFacility ] | gml::LengthType [0..1] | 規模_車道幅員。 |
| uro:minimumDepth     [ PortTransportationFacility ] | gml::LengthType [0..1] | 規模_最小水深。 |
| uro:numberOfAircraftParkingSpaces     [ PortTransportationFacility ] | xs::integer [0..1] | 規模_駐機数。 |
| uro:pavementType     [ PortTransportationFacility ] | gml::CodeType [0..1] | 舗装形態/塗装形態。コードリスト（PortTransportationFacility_pavementType.xml）より選択する。 |
| uro:mainCargo     [ PortTransportationFacility ] | gml::CodeType [0..1] | 主要取扱貨物名。コードリスト（PortAttribute_mainCargo.xml）より選択する。 |
| uro:totalCost     [ PortTransportationFacility ] | xs::integer [0..1] | 事業費－総額。 |
| uro:subsidy     [ PortTransportationFacility ] | xs::integer [0..1] | 事業費－補助金額。 |
| uro:note     [ PortTransportationFacility ] | xs::string [0..1] | 備考。 |

