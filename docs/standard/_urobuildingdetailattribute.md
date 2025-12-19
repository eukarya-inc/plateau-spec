# 4.2.3.2.2 uro:BuildingDetailAttribute

## 4.2.3.2.2 uro:BuildingDetailAttribute

**表 4-46**

| 型の定義 | 都市計画法に基づき実施される都市計画基礎調査において収集された、建築物に関する基礎的な情報。 | ← |
| --- | --- | --- |
| 上位の型 | uro:BuildingAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:serialNumberOfBuildingCertification     [ BuildingDetailAttribute ] | xs::string [0..1] | 建築確認申請番号。 |
| uro:siteArea     [ BuildingDetailAttribute ] | gml::MeasureType [0..1] | 当該建築物が立地する敷地の面積。単位はm2（uom=”m2”）とする。 |
| uro:totalFloorArea     [ BuildingDetailAttribute ] | gml::MeasureType [0..1] | 当該建築物の各階の床面積の合計。単位はm2（uom=”m2”）とする。 |
| uro:buildingFootprintArea     [ BuildingDetailAttribute ] | gml::MeasureType [0..1] | 建築物の壁や柱の中心線で囲まれた部分の水平投影面積。単位はm2（uom=”m2”）とする。 |
| uro:buildingRoofEdgeArea     [ BuildingDetailAttribute ] | gml::MeasureType [0..1] | 屋根を含む建築物の水平投影面積。単位はm2（uom=”m2”）とする。 |
| uro:developmentArea     [ BuildingDetailAttribute ] | gml::MeasureType [0..1] | 開発された面積。単位はm2（uom=”m2”）とする。 |
| uro:buildingStructureType     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 構造種別。コードリスト（ BuildingDetailAttribute_buildingStructureType.xml）より選択する。 |
| uro:buildingStructureOrgType     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市ごとの独自の区分に基づく建築物の構造種別。コードリスト（BuildingDetailAttribute_buildingStructureOrgType.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:fireproofStructureType     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 耐火構造区分。コードリスト（ BuildingDetailAttribute_fireproofStructureType.xml）より選択する。 |
| uro:implementingBody     [ BuildingDetailAttribute ] | xs::string [0..1] | 建築物建築の実施主体の名称。 |
| uro:urbanPlanType     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 建築物が立地する土地が属する都市計画区域の区分。コードリスト（ Common_urbanPlanType.xml）より選択する。 |
| uro:areaClassificationType     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 建築物が立地する土地が属する区域区分。コードリスト（ Common_areaClassificationType.xml）より選択する。 |
| uro:districtsAndZonesType     [ BuildingDetailAttribute ] | gml::CodeType [0..*] | 建築物が立地する土地が属する地域地区の区分。コードリスト（ Common_districtsAndZonesType.xml）より選択する。建築物が複数の地域地区に含まれる場合は、複数を列挙する。 |
| uro:landUseType     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 建築物が立地する土地の土地利用区分。コードリスト（ Common_landUseType.xml）より選択する。 |
| uro:reference     [ BuildingDetailAttribute ] | xs::string [0..1] | 建築物の位置を示す図面上の番号。 |
| uro:majorUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | urf:orgUsageよりも粗い区分による都市独自の分類。コードリスト（BuildingDetailAttribute_majorUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:majorUsage2     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | uro:orgUsageよりも粗く、uro:majorUsageよりも細かい区分による都市独自の分類。コードリスト（BuildingDetailAttribute_majorUsage2.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:orgUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市計画基礎調査実施要領（国土交通省都市局）に示された建築物の「用途分類」に相当する都市独自の分類。コードリスト（BuildingDetailAttribute_orgUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:orgUsage2     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市計画基礎調査実施要領（国土交通省都市局）に示された建築物の「用途分類」のうち、商業施設、文教厚生施設、運輸倉庫施設、工場が詳細化された区分に相当する都市独自の分類。コードリスト（BuildingDetailAttribute_orgUsage2.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:detailedUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | uro:orgUsage2よりも細かい区分による都市独自の分類。コードリスト（ BuildingDetailAttribute_detailedUsage.xml）より選択する。本属性を使用する場合は、本製品仕様書に示すコードリストを必要に応じて加工すること。 |
| uro:detailedUsage2     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | uro:detailedUsageよりも細かい区分による都市独自の分類。コードリスト（BuildingDetailAttribute_detailedUsage2.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:detailedUsage3     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | uro:detailedUsage2よりも細かい区分による都市独自の分類。コードリスト（BuildingDetailAttribute_detailedUsage3.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:groundFloorUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市ごとの独自の区分に基づく建築物1階の用途。コードリスト（BuildingDetailAttribute_groundFloorUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:secondFloorUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市ごとの独自の区分に基づく建築物の2階又は2階以上の用途。コードリスト（BuildingDetailAttribute_secondFloorUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:thirdFloorUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市ごとの独自の区分に基づく建築物の3階又は3階以上の用途。コードリスト（BuildingDetailAttribute_thirdFloorUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:basementUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市ごとの独自の区分に基づく建築物の地下の用途。コードリスト（BuildingDetailAttribute_basementFloorUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:basementFirstUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市ごとの独自の区分に基づく建築物の地下1階の用途。コードリスト（BuildingDetailAttribute_basementFirstUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:basementSecondUsage     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 都市ごとの独自の区分に基づく建築物の地下2階の用途。コードリスト（BuildingDetailAttribute_basementSecondUsage.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:vacancy     [ BuildingDetailAttribute ] | gml::CodeType [0..1] | 空き家か否かの別。コードリスト（ BuildingDetailAttribute_vacancy.xml）より選択する。 |
| uro:buildingCoverageRate     [ BuildingDetailAttribute ] | xs::double [0..1] | 建蔽率（敷地面積に対する建築面積の割合）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:floorAreaRate     [ BuildingDetailAttribute ] | xs::double [0..1] | 容積率（敷地面積に対する延床面積の割合）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:specifiedBuildingCoverageRate     [ BuildingDetailAttribute ] | xs::double [0..1] | 指定建蔽率（用途地域別に定められている建蔽率）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:specifiedFloorAreaRate     [ BuildingDetailAttribute ] | xs::double [0..1] | 指定容積率（都市計画で定められる容積率の最高限度）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:standardFloorAreaRate     [ BuildingDetailAttribute ] | xs::double [0..1] | 基準容積率（前面道路の幅員が12ｍ未満の場合に、前面道路の幅員による限度により算出される容積率）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:buildingHeight     [ BuildingDetailAttribute ] | gml::LengthType [0..1] | 建築基準法施行令第2条に定義される地盤面からの建築物の高さ。単位はm（uom=”m”）とする。 |
| uro:eaveHeight     [ BuildingDetailAttribute ] | gml::LengthType [0..1] | 建築基準法施行令第2条に定義される建築物の地盤面から軒桁までの高さ。単位はm（uom=”m”）とする。 |
| uro:note     [ BuildingDetailAttribute ] | xs::string [0..1] | その他建築物に関して特筆すべき事項。 |
| uro:surveyYear     [ BuildingDetailAttribute ] | xs::gYear [1..1] | 建物利用現況調査の実施年（西暦）。 |

