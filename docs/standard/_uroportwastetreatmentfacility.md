# 4.26.3.3.10 uro:PortWasteTreatmentFacility

## 4.26.3.3.10 uro:PortWasteTreatmentFacility

**表 4-883**

| 型の定義 | 港湾施設である廃棄物処理施設の属性を記述するためのデータ型。 | ← |
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
| uro:structureType     [ PortWasteTreatmentFacility ] | gml::CodeType [0..1] | 構造形式。コードリスト（PortAttribute_structureType.xml）より選択する。 |
| uro:perimeter     [ PortWasteTreatmentFacility ] | gml::LengthType [0..1] | 延長_外周建設延長。 |
| uro:mainPartLength     [ PortWasteTreatmentFacility ] | gml::LengthType [0..1] | 延長_機能保有延長。 |
| uro:innerShoreLength     [ PortWasteTreatmentFacility ] | gml::LengthType [0..1] | 延長_内護岸延長。 |
| uro:ceilingHeight     [ PortWasteTreatmentFacility ] | gml::LengthType [0..1] | 天端高。 |
| uro:waveDissipatorLength     [ PortWasteTreatmentFacility ] | gml::LengthType [0..1] | 消波工延長。 |
| uro:mainMaterial     [ PortWasteTreatmentFacility ] | gml::CodeType [0..1] | 主要用材。コードリスト（Common_mainMaterial.xml）より選択する。 |
| uro:wasteType     [ PortWasteTreatmentFacility ] | gml::CodeType [0..1] | 廃棄物の種類。コードリスト（PortWasteTreatmentFacilityAttribute_wasteType.xml）より選択する。 |
| uro:plannedDisposalArea     [ PortWasteTreatmentFacility ] | gml::MeasureType [0..1] | 計画処分面積。 |
| uro:plannedDisposalAmount     [ PortWasteTreatmentFacility ] | xs::integer [0..1] | 計画処分量。 |
| uro:receivingCapacity     [ PortWasteTreatmentFacility ] | xs::integer [0..1] | 受入容量。 |
| uro:shipType     [ PortWasteTreatmentFacility ] | xs::string [0..1] | 船型。 |
| uro:unitOfReceivingCapacity     [ PortWasteTreatmentFacility ] | gml::CodeType [0..1] | 受入容量単位。コードリスト（PortAttribute_storageCapacityUnit.xml）より選択する。 |
| uro:acquisitionYear     [ PortWasteTreatmentFacility ] | xs::gYear [0..1] | 取得年度。 |
| uro:totalCost     [ PortWasteTreatmentFacility ] | xs::integer [0..1] | 事業費－総額。 |
| uro:subsidy     [ PortWasteTreatmentFacility ] | xs::integer [0..1] | 事業費－補助金額。 |
| uro:note     [ PortWasteTreatmentFacility ] | xs::string [0..1] | 備考。 |

