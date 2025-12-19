# 4.2.3.2.3 uro:LargeCustomerFacilityAttribute

## 4.2.3.2.3 uro:LargeCustomerFacilityAttribute

**表 4-47**

| 型の定義 | 都市計画法に基づき実施される都市計画基礎調査において収集された、大規模小売店舗や大規模集客施設に関する基礎的な情報。 | ← |
| --- | --- | --- |
| 上位の型 | uro:BuildingAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:class     [ LargeCustomerFacilityAttribute ] | gml::CodeType [0..1] | 集客施設の種類。コードリスト（ LargeCustomerFacilityAttribute_class.xml）より選択する。 |
| uro:name     [ LargeCustomerFacilityAttribute ] | xs::string [0..1] | 集客施設の名称。 |
| uro:capacity     [ LargeCustomerFacilityAttribute ] | xs::integer [0..1] | 集客施設の収容人数。（病院の場合は、病床数、大学等の場合は学生数とする。） |
| uro:owner     [ LargeCustomerFacilityAttribute ] | xs::string [0..1] | 施設の所有者の名称。 |
| uro:totalFloorArea     [ LargeCustomerFacilityAttribute ] | gml::MeasureType [0..1] | 集客施設各階の床面積を合計した面積。単位はm2（uom=”m2”）とする。 |
| uro:totalStoreFloorArea     [ LargeCustomerFacilityAttribute ] | gml::MeasureType [0..1] | 集客施設各階における店舗の床面積を合計した面積。単位はm2（uom=”m2”）とする。 |
| uro:inauguralDate     [ LargeCustomerFacilityAttribute ] | xs::date [0..1] | 集客施設が運営を開始した年月日。 |
| uro:yearOpened     [ LargeCustomerFacilityAttribute ] | xs::gYear [0..1] | 開設年。 |
| uro:yearClosed     [ LargeCustomerFacilityAttribute ] | xs::gYear [0..1] | 廃止年。 |
| uro:keyTenants     [ LargeCustomerFacilityAttribute ] | xs::string [0..1] | 集客施設が商業施設の場合の、主要なテナントの名称。 |
| uro:availability     [ LargeCustomerFacilityAttribute ] | xs::boolean [0..1] | 集客施設が医療施設の場合の、3次医療圏規模の有無。 |
| uro:urbanPlanType     [ LargeCustomerFacilityAttribute ] | gml::CodeType [0..1] | 集客施設が立地する土地が属する都市計画区域の区分。コードリスト（ Common_urbanPlanType.xml）より選択する。 |
| uro:areaClassificationType     [ LargeCustomerFacilityAttribute ] | gml::CodeType [0..1] | 集客施設が立地する土地が属する区域区分。コードリスト（ Common_areaClassificationType.xml）より選択する。 |
| uro:districtsAndZonesType     [ LargeCustomerFacilityAttribute ] | gml::CodeType [0..*] | 集客施設が立地する土地が属する地域地区の区分。コードリスト（ Common_districtsAndZonesType.xml）より選択する。建築物が複数の地域地区に含まれる場合は、複数を列挙する。 |
| uro:landUseType     [ LargeCustomerFacilityAttribute ] | gml::CodeType [0..1] | 集客施設が立地する土地の土地利用区分。コードリスト（ Common_landUseType.xml）より選択する。 |
| uro:reference     [ LargeCustomerFacilityAttribute ] | xs::string [0..1] | 図面対照番号。集客施設の位置を示す図面上の番号。 |
| uro:note     [ LargeCustomerFacilityAttribute ] | xs::string [0..1] | その他集客施設に関して特筆すべき事項。 |
| uro:surveyYear     [ LargeCustomerFacilityAttribute ] | xs::gYear [1..1] | 集客施設の立地状況調査の実施年（西暦）。 |

