# 4.26.3.3.6 uro:CargoHandlingFacility

## 4.26.3.3.6 uro:CargoHandlingFacility

**表 4-879**

| 型の定義 | 港湾施設である荷さばき施設の属性を記述するためのデータ型。 | ← |
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
| uro:mainCargo     [ CargoHandlingFacility ] | gml::CodeType [0..1] | 主要取扱貨物名。コードリスト（PortAttribute_mainCargo.xml）より選択する。 |
| uro:mooringFacility     [ CargoHandlingFacility ] | xs::string [0..1] | 係留施設名。 |
| uro:liftableLoad     [ CargoHandlingFacility ] | gml::MeasureType [0..1] | 荷役能力_吊り上げ荷重。 |
| uro:ability     [ CargoHandlingFacility ] | xs::integer [0..1] | 荷役能力_1時間あたりの能力。 |
| uro:packingName     [ CargoHandlingFacility ] | gml::CodeType [0..1] | 荷姿名。コードリスト（CargoHandlingFacility_packingName.xml）より選択する。 |
| uro:acquisitionYear     [ CargoHandlingFacility ] | xs::gYear [0..1] | 取得年度。 |
| uro:innerTotalFloorArea     [ CargoHandlingFacility ] | gml::MeasureType [0..1] | 臨港地区内－総床面積。 |
| uro:innerOfSiteArea     [ CargoHandlingFacility ] | gml::MeasureType [0..1] | 臨港地区内－敷地面積。 |
| uro:outerOfTotalFloorArea     [ CargoHandlingFacility ] | gml::MeasureType [0..1] | 臨港地区外－総床面積。 |
| uro:outerSiteArea     [ CargoHandlingFacility ] | gml::MeasureType [0..1] | 臨港地区外－敷地面積。 |
| uro:mainMaterial     [ CargoHandlingFacility ] | gml::CodeType [0..1] | 主要用材。コードリスト（Common_mainMaterial.xml）より選択する。 |
| uro:totalCost     [ CargoHandlingFacility ] | xs::integer [0..1] | 事業費－総額。 |
| uro:note     [ CargoHandlingFacility ] | xs::string [0..1] | 備考。 |

