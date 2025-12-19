# 4.8.3.2.1 uro:LandUseDetailAttribute

## 4.8.3.2.1 uro:LandUseDetailAttribute

**表 4-300**

| 型の定義 | 都市計画に関する基礎調査の一つとして、土地利用の現況と変化の動向を把握することを目的とし都市計画法第 6 条の規定に基づき実施される調査の結果。 | ← |
| --- | --- | --- |
| 上位の型 | uro:LandUseAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:id     [ LandUseDetailAttribute ] | xs::string [0..1] | 土地利用現況図における識別子。 |
| uro:orgLandUse     [ LandUseDetailAttribute ] | gml::CodeType [0..1] | 都市独自の分類による土地利用用途。コードリスト（LandUseDetailAttribute_orgLandUse.xml）より選択する。本属性を使用する場合は、コードリストを作成すること。 |
| uro:nominalArea     [ LandUseDetailAttribute ] | gml::MeasureType [0..1] | 図上計測面積を調整した値。単位はm2（uom=”m2”）とする。 |
| uro:ownerType     [ LandUseDetailAttribute ] | gml::CodeType [0..1] | 土地所有者の区分。コードリスト（Common_ownerType.xml）より選択する。 |
| uro:owner     [ LandUseDetailAttribute ] | xs::string [0..1] | 土地所有者の名称。 |
| uro:areaInSquareMeter     [ LandUseDetailAttribute ] | gml::MeasureType [0..1] | 図上計測面積。単位はm2（uom=”m2”）とする。 |
| uro:areaInHa     [ LandUseDetailAttribute ] | gml::MeasureType [0..1] | 図上計測面積（ha換算数）。単位はha（uom=”ha”）とする。 |
| uro:buildingCoverageRate     [ LandUseDetailAttribute ] | xs::double [0..1] | 建蔽率（敷地面積に対する建築面積の割合）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:floorAreaRate     [ LandUseDetailAttribute ] | xs::double [0..1] | 容積率（敷地面積に対する延床面積の割合）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:specifiedBuildingCoverageRate     [ LandUseDetailAttribute ] | xs::double [0..1] | 指定建蔽率（用途地域別に定められている建蔽率）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:specifiedFloorAreaRate     [ LandUseDetailAttribute ] | xs::double [0..1] | 指定容積率（都市計画で定められる容積率の最高限度）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:standardFloorAreaRate     [ LandUseDetailAttribute ] | xs::double [0..1] | 基準容積率（前面道路の幅員が12ｍ未満の場合に、前面道路の幅員による限度により算出される容積率）。全体を「100」とする割合（百分率）で記述する。単位は％。 |
| uro:urbanPlanType     [ LandUseDetailAttribute ] | gml::CodeType [0..1] | 土地が属する都市計画区域の区分。コードリスト（Common_urbanPlanType.xml）より選択する。 |
| uro:areaClassificationType     [ LandUseDetailAttribute ] | gml::CodeType [0..1] | 土地が属する区域区分。コードリスト（Common_areaClassificationType.xml）より選択する。 |
| uro:districtsAndZonesType     [ LandUseDetailAttribute ] | gml::CodeType [0..*] | 土地が属する地域地区の区分。コードリスト（Common_districtsAndZonesType.xml）より選択する。土地利用が複数の地域地区に含まれる場合は、複数を列挙する。 |
| uro:prefecture     [ LandUseDetailAttribute ] | gml::CodeType [0..1] | 土地が所在する都道府県の都道府県コ－ド。JIS X0401に定義される2桁の半角数字。コードリスト（Common_localPublicAuthorities.xml）より選択する。 |
| uro:city     [ LandUseDetailAttribute ] | gml::CodeType [0..1] | 土地が所在する市区町村の市区町村コ－ド。 JIS X0401に定義される2桁の半角数字とJIS X0402に定義される3桁の半角数字とを組み合わせた5桁の半角数字。政令市の場合は、区の市区町村コードとする。コードリスト（Common_localPublicAuthorities.xml）より選択する。 運用上必須とする。 |
| uro:reference     [ LandUseDetailAttribute ] | xs::string [0..1] | 土地の位置を示す図面上の番号。 |
| uro:note     [ LandUseDetailAttribute ] | xs::string [0..1] | その他土地に関して特筆すべき事項。 |
| uro:surveyYear     [ LandUseDetailAttribute ] | xs::gYear [0..1] | 土地利用現況調査の実施年（西暦）。 |

