# 4.26.3.4.1 uro:FishingPortFacility

## 4.26.3.4.1 uro:FishingPortFacility

**表 4-889**

| 型の定義 | 漁港施設の内容を表すデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:FishingPortAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:facilityId     [ FacilityAttribute ] | xs::string [0..1] | 施設の用途や区分ごとに付与された管理番号。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:facilityDetailsType     [ FishingPortFacility ] | gml::CodeType [1..1] | 宿泊施設、休憩所 等の施設の種類。コードリスト（FishingPortFacilityAttribute_facilityDetailsType.xml）より選択する。 |
| uro:portName     [ FishingPortFacility ] | xs::string [1..1] | 漁港の名称。 |
| uro:portType     [ FishingPortFacility ] | gml::CodeType [1..1] | 漁港漁場整備法施行規則 第九条に基づく漁港の種類。 コードリスト（ FishingPortFacilityAttribute_portType.xml）より選択する。 |
| uro:address     [ FishingPortFacility ] | xs::string [1..1] | 所在地。 |
| uro:designatedArea     [ FishingPortFacility ] | xs::string [1..1] | 区域。 |
| uro:designation     [ FishingPortFacility ] | xs::string [0..*] | 漁港の指定。 |
| uro:designatedAdministrator     [ FishingPortFacility ] | xs::string [0..*] | 漁港管理者の指定。 |
| uro:referenceNumber     [ FishingPortFacility ] | xs::string [0..*] | 漁港の平面図対象番号。 |
| uro:grantType     [ FishingPortFacility ] | gml::CodeType [0..1] | 施設区分名。コードリスト（PortAttribute_grantType.xml）より選択する。 |
| uro:administrator     [ FishingPortFacility ] | xs::string [0..1] | 所有者の名称。 |
| uro:facilityManager     [ FishingPortFacility ] | xs::string [0..1] | 管理者の名称。 |
| uro:structureType     [ FishingPortFacility ] | gml::CodeType [0..1] | 構造_様式又は形式。コードリスト（FishingPortFacilityAttribute_structureType.xml）より選択する。 |
| uro:mainMaterial     [ FishingPortFacility ] | gml::CodeType [0..1] | 構造_主要用材。コードリスト（Common_mainMaterial.xml）より選択する。 |
| uro:otherStructure     [ FishingPortFacility ] | xs::string [0..1] | 構造_その他の構造。 |
| uro:length     [ FishingPortFacility ] | gml::LengthType [0..1] | 規模_延長。 |
| uro:width     [ FishingPortFacility ] | gml::LengthType [0..1] | 規模_幅員。 |
| uro:ceilingHeight     [ FishingPortFacility ] | gml::LengthType [0..1] | 規模_天端高。 |
| uro:depth     [ FishingPortFacility ] | gml::LengthType [0..1] | 規模_水深。 |
| uro:area     [ FishingPortFacility ] | gml::MeasureType [0..1] | 規模_面積。 |
| uro:otherSizeDescription     [ FishingPortFacility ] | xs::string [0..1] | 規模_その他の規模数量。 |
| uro:dateOfConstructionOrAcquisition     [ FishingPortFacility ] | xs::date [0..1] | 建設又は取得の年月日。 |
| uro:cost     [ FishingPortFacility ] | xs::integer [0..1] | 建設又は取得の価格。 |
| uro:note     [ FishingPortFacility ] | xs::string [0..1] | 備考。 |

