# 4.26.3.3.3 uro:MooringFacility

## 4.26.3.3.3 uro:MooringFacility

**表 4-876**

| 型の定義 | 港湾施設である係留施設の属性を記述するためのデータ型。 | ← |
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
| uro:mainPartLength     [ MooringFacility ] | gml::LengthType [0..1] | 延長－取付部を除く延長。 |
| uro:totalLength     [ MooringFacility ] | gml::LengthType [0..1] | 延長－取付部を含む延長。 |
| uro:facilityWidth     [ MooringFacility ] | gml::LengthType [0..1] | 施設の幅。 |
| uro:apronWidth     [ MooringFacility ] | gml::LengthType [0..1] | エプロン幅。 |
| uro:plannedDepth     [ MooringFacility ] | gml::LengthType [0..1] | 水深－計画上の水深。 |
| uro:currentDepth     [ MooringFacility ] | gml::LengthType [0..1] | 水深－現在の水深。 |
| uro:area     [ MooringFacility ] | gml::MeasureType [0..1] | 面積。 |
| uro:ceilingHeight     [ MooringFacility ] | gml::LengthType [0..1] | 天端高。 |
| uro:gravityResistant     [ MooringFacility ] | gml::MeasureType [0..1] | 耐重力。 |
| uro:form     [ MooringFacility ] | gml::CodeType [0..1] | 形態。コードリスト（PortAttribute_form.xml）より選択する。 |
| uro:mainVessels     [ MooringFacility ] | gml::CodeType [0..1] | 主要利用船舶の種類。コードリスト（MooringFacility_mainVessels.xml）より選択する。 |
| uro:mooringPostWeight     [ MooringFacility ] | gml::MeasureType [0..1] | 附帯設備－係船柱の重さ。 |
| uro:numberOfMooringPosts     [ MooringFacility ] | xs::integer [0..1] | 附帯設備－係船柱の数。 |
| uro:resistantMaterial     [ MooringFacility ] | xs::integer [0..1] | 附帯設備－防げん材。 |
| uro:lighting     [ MooringFacility ] | xs::integer [0..1] | 附帯設備－照明設備。 |
| uro:stairs     [ MooringFacility ] | xs::integer [0..1] | 附帯設備－階段等。 |
| uro:lifesavingAppliances     [ MooringFacility ] | xs::string [0..1] | 附帯設備－救命設備の名称。 |
| uro:numberOfLifesavingAppliances     [ MooringFacility ] | xs::integer [0..1] | 附帯設備－救命設備の数。 |
| uro:bumper     [ MooringFacility ] | gml::LengthType [0..1] | 附帯設備－車止め。 |
| uro:numberOfVehicleBoardings     [ MooringFacility ] | xs::integer [0..1] | 附帯設備－車両乗降設備－基数。 |
| uro:vehicleBoardingWidth     [ MooringFacility ] | gml::LengthType [0..1] | 附帯設備－車両乗降設備－幅員。 |
| uro:shipType     [ MooringFacility ] | xs::string [0..1] | 対象船舶－船型（D／W）。 |
| uro:numberOfSeats     [ MooringFacility ] | xs::integer [0..1] | 対象船舶－船席数。 |
| uro:mainCargo     [ MooringFacility ] | gml::CodeType [0..1] | 主要取扱貨物名。コードリスト（PortAttribute_mainCargo.xml）より選択する。 |
| uro:structureType     [ MooringFacility ] | gml::CodeType [0..1] | 構造形式。コードリスト（PortAttribute_structureType.xml）より選択する。 |
| uro:mainMaterial     [ MooringFacility ] | gml::CodeType [0..1] | 主要用材。コードリスト（Common_mainMaterial.xml）より選択する。 |
| uro:totalCost     [ MooringFacility ] | xs::integer [0..1] | 事業費－総額。 |
| uro:subsidy     [ MooringFacility ] | xs::integer [0..1] | 事業費－補助金額。 |
| uro:note     [ MooringFacility ] | xs::string [0..1] | 備考。 |

