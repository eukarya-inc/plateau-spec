# 4.26.3.3.13 uro:PortWelfareFacility

## 4.26.3.3.13 uro:PortWelfareFacility

**表 4-886**

| 型の定義 | 面積_防波堤等の外側。 | ← |
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
| uro:totalFloorArea     [ PortWelfareFacility ] | gml::MeasureType [0..1] | 面積_防波堤等の外側。 |
| uro:totalCost     [ PortWelfareFacility ] | xs::integer [0..1] | 事業費－総額。 |
| uro:note     [ PortWelfareFacility ] | xs::string [0..1] | 備考。 |

