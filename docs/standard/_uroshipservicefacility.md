# 4.26.3.3.9 uro:ShipServiceFacility

## 4.26.3.3.9 uro:ShipServiceFacility

**表 4-882**

| 型の定義 | 港湾施設である船舶役務用施設の属性を記述するためのデータ型。 | ← |
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
| uro:shipType     [ ShipServiceFacility ] | xs::string [0..1] | 対象船舶－船型（D／W）。 |
| uro:supplyAbility     [ ShipServiceFacility ] | xs::integer [0..1] | 供給能力容量。 |
| uro:supplyAbilityUnit     [ ShipServiceFacility ] | gml::CodeType [0..1] | 供給能力単位。コードリスト（ShipServiceFacility_supplyAbilityUnit.xml）より選択する。 |
| uro:mooringPlace     [ ShipServiceFacility ] | xs::string [0..1] | 補給を受ける船舶の係留場所。 |
| uro:length     [ ShipServiceFacility ] | gml::LengthType [0..1] | 長さ。 |
| uro:width     [ ShipServiceFacility ] | gml::LengthType [0..1] | 幅。 |
| uro:area     [ ShipServiceFacility ] | gml::MeasureType [0..1] | 面積。 |
| uro:acquisitionYear     [ ShipServiceFacility ] | xs::gYear [0..1] | 取得年度。 |
| uro:totalCost     [ ShipServiceFacility ] | xs::integer [0..1] | 事業費－総額。 |
| uro:note     [ ShipServiceFacility ] | xs::string [0..1] | 備考。 |

