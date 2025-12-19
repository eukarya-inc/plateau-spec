# 4.26.3.3.15 uro:CyberportMarinaAndPBS

## 4.26.3.3.15 uro:CyberportMarinaAndPBS

**表 4-888**

| 型の定義 | 港湾施設であるマリーナ/PBSの属性を記述するためのデータ型。 | ← |
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
| uro:geologicalType     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 海底の地質名。コードリスト（PortAttribute_geologicalType.xml）より選択する。 |
| uro:obstructingStructures     [ CyberportMarinaAndPBS ] | xs::string [0..1] | 構造物による制限－構造物名。 |
| uro:mainPartLength     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 延長－取付部を除く延長。 |
| uro:totalLength     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 延長－取付部を含む延長。 |
| uro:waveDissipatorLength     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 消波工延長。 |
| uro:facilityWidth     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 施設の幅。 |
| uro:apronWidth     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | エプロン幅。 |
| uro:restrictionStructure     [ CyberportMarinaAndPBS ] | xs::string [0..1] | 構造物による制限。 |
| uro:plannedDepth     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 計画上の水深。 |
| uro:currentDepth     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 現在の水深。 |
| uro:innerTotalFloorArea     [ CyberportMarinaAndPBS ] | gml::MeasureType [0..1] | 臨港地区内－総床面積。 |
| uro:innerOfSiteArea     [ CyberportMarinaAndPBS ] | gml::MeasureType [0..1] | 臨港地区内－敷地面積。 |
| uro:outerOfTotalFloorArea     [ CyberportMarinaAndPBS ] | gml::MeasureType [0..1] | 臨港地区外－総床面積。 |
| uro:outerSiteArea     [ CyberportMarinaAndPBS ] | gml::MeasureType [0..1] | 臨港地区外－敷地面積。 |
| uro:ceilingHeight     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 天端高。 |
| uro:gravityResistant     [ CyberportMarinaAndPBS ] | gml::MeasureType [0..1] | 耐重力。 |
| uro:form     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 形態。コードリスト（PortAttribute_form.xml）より選択する。 |
| uro:areaType     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 防波堤等の内外の区分。コードリスト（CyberportMarinaAndPBS_areaType.xml）より選択する。 |
| uro:mainVessels     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 主要利用船舶の種類。コードリスト（CyberportMarinaAndPBS_mainVessels.xml）より選択する。 |
| uro:isDredged     [ CyberportMarinaAndPBS ] | xs::boolean [0..1] | 浚渫の有無 0：無、1：有 |
| uro:mooringPostWeight     [ CyberportMarinaAndPBS ] | gml::MeasureType [0..1] | 附帯設備－係船柱の重さ。単位は㎏とする。 |
| uro:numberOfMooringPosts     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 附帯設備－係船柱の個数。単位は個とする。 |
| uro:resistantMaterial     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 附帯設備－防げん材。 |
| uro:lighting     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 附帯設備－照明設備。 |
| uro:stairs     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 附帯設備－階段等。 |
| uro:lifesaving     [ CyberportMarinaAndPBS ] | xs::string [0..1] | 附帯設備－救命設備の名称。 |
| uro:lifesavingNumber     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 附帯設備－救命設備の数。 |
| uro:bumper     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 附帯設備－車止め。 |
| uro:numberOfVehicleBoardings     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 附帯設備－車両乗降設備－基数。 |
| uro:vehicleBoardingWidth     [ CyberportMarinaAndPBS ] | gml::LengthType [0..1] | 附帯設備－車両乗降設備－幅員。 |
| uro:shipType     [ CyberportMarinaAndPBS ] | xs::string [0..1] | 対象船舶－船型(D/W)。 |
| uro:numberOfSeats     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 対象船舶－船席数。 |
| uro:mainCargo     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 主要取扱貨物名。コードリスト（PortAttribute_mainCargo.xml）より選択する。 |
| uro:storageCapacity     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 保管容量－値。 |
| uro:storageCapacityUnit     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 保管容量－単位。コードリスト（PortAttribute_storageCapacityUnit.xml）より選択する。 |
| uro:structureType     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 構造形式。コードリスト（PortAttribute_structureType.xml）より選択する。 |
| uro:mainMaterial     [ CyberportMarinaAndPBS ] | gml::CodeType [0..1] | 主要用材。コードリスト（Common_mainMaterial.xml）より選択する。 |
| uro:totalCost     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 事業費－総額。 |
| uro:subsidy     [ CyberportMarinaAndPBS ] | xs::integer [0..1] | 事業費－補助金額。 |
| uro:note     [ CyberportMarinaAndPBS ] | xs::string [0..1] | 備考。 |

