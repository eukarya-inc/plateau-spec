# 4.10.3.10.5 urf:DisasterPreventionBlockImprovementZonePlan

## 4.10.3.10.5 urf:DisasterPreventionBlockImprovementZonePlan

**表 4-433**

| 型の定義 | 密集市街地整備法第32条第1項の規定による防災街区整備地区計画。 （都市計画法第12条の4第1項第2号） 防災街区整備地区計画については、都市計画法第12条の4第2項に定める事項のほか、都市計画に、密集市街地法第32条第2項第1号及び第2号に掲げる事項を定めるものとするとともに、第3号に掲げる事項を定めるよう努めるものとする。特定建築物地区整備計画防災街区整備地区整備計画当該防災街区整備地区計画の目標その他当該区域の整備に関する方針関連役割「地区整備計画」により、「特定建築物地区整備計画」及び「防災街区整備地区整備計画」を保持する。属性gml:nameは、都市計画法第12条の4第2項で定める名称（地区計画を識別する名前）とする。多重度は任意となっているが、運用上必須とする。文字列とする。属性urf:functionは、都市計画法第12条の4第2項に定める地区計画の種類とする。促進区を定める場合、当該地区計画は、再開発等促進区又は開発整備促進区を定める地区計画となる。コードリスト（DistrictPlan_function.xml,DistrictPlan_function.xml,DistrictPlan_function.xml）より選択する。多重度は任意となっているが、運用上必須とする。属性urf:usageは使用しない。属性urf:locationは、都市計画法第12条の4第2項に定める位置とする。町丁目又は字まで記載する。多重度は任意となっているが、運用上必須とする。属性urf:objectivesは、密集市街地整備法第32条第2項第3号に定める当該地区計画の目標とする。属性urf:policyは、密集市街地整備法第32条第2項第3号に定める当該地区計画の方針とする。関連役割urf:districtDevelopmentPlanは、防災街区整備地区計画に定められた特定建築物地区整備計画及び防災街区整備地区整備計画とする。関連役割urf:promotionDistrictは、使用しない。 | ← |
| --- | --- | --- |
| 上位の型 | urf:_AbstractDistrictPlan | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |
| core:creationDate     [ _CityObject ] | xs::date [0..1] | データが作成された日。運用上必須とする。 |
| core:terminationDate     [ _CityObject ] | xs::date [0..1] | データが削除された日。 |
| core:relativeToTerrain     [ _CityObject ] | core::RelativeToTerrainType [0..1] | 地表面との相対的な位置関係。標準製品仕様書では使用しない。 |
| core:relativeToWater     [ _CityObject ] | core::RelativeToWaterType [0..1] | 水面との相対的な位置関係。標準製品仕様書では使用しない。 |
| urf:class     [ _UrbanFunction ] | gml::CodeType [0..1] | 分類。都市計画決定情報の場合は使用しない。 |
| urf:function     [ _UrbanFunction ] | gml::CodeType [0..*] | 機能。これを実装する下位の地物型に応じて使用するコードリストを選択する。 |
| urf:usage     [ _UrbanFunction ] | gml::CodeType [0..*] | 区域の用途。これを実装する下位の地物型に応じて使用するコードリストを選択する。 |
| urf:validFrom     [ _UrbanFunction ] | xs::date [0..1] | 効力を生じる日の年月日を西暦（YYYY-MM-DD）で記述する。都市計画決定情報の場合は、当初の決定日とする。多重度は任意となっているが、運用上必須とする。年月日が不明な場合は、0001-01-01とする。 |
| urf:validFromType     [ _UrbanFunction ] | gml::CodeType [0..1] | 効力を生じる日の種類。コードリスト（Common_validType.xml）より選択する。都市計画決定情報の場合は、多重度は任意となっているが、運用上必須とする。 |
| urf:enactmentFiscalYear     [ _UrbanFunction ] | xs::gYear [0..1] | 決定された年度（西暦）。 |
| urf:validTo     [ _UrbanFunction ] | xs::date [0..1] | 効力を失う日の年月日を西暦（YYYY-MM-DD）で記述する。 |
| urf:validToType     [ _UrbanFunction ] | gml::CodeType [0..1] | 効力を失う日の種類。コードリスト（Common_validType.xml）より選択する。 |
| urf:expirationFiscalYear     [ _UrbanFunction ] | xs::gYear [0..1] | 効力を失う日の年度（西暦）。 |
| urf:legalGrounds     [ _UrbanFunction ] | xs::string [0..1] | 法的背景。根拠法令を記述する。 |
| urf:custodian     [ _UrbanFunction ] | xs::string [0..1] | 決定者。都市計画決定情報の場合は、都市計画法第15条第1項で規定する都市計画を定める者の名称。多重度は任意となっているが、運用上必須とする。不明な場合は文字列で「Null」と入力する。 |
| urf:notificationNumber     [ _UrbanFunction ] | xs::string [0..1] | 告示番号（当初）。多重度は任意となっているが、都市計画決定情報の場合は、運用上必須とする。不明な場合は文字列で「Null」と入力する。 |
| urf:finalNotificationDate     [ _UrbanFunction ] | xs::date [0..1] | 告示（最終）の日付。多重度は任意となっているが、都市計画決定情報の場合は、運用上必須とする。 |
| urf:finalNotificationNumber     [ _UrbanFunction ] | xs::string [0..1] | 告示番号（最終）。多重度は任意となっているが、都市計画決定情報の場合は、運用上必須とする。 |
| urf:urbanPlanType     [ _UrbanFunction ] | gml::CodeType [0..1] | 指定された区域が属する都市計画区域の区分。コードリスト（Common_urbanPlanType.xml）より選択する。都市計画区域及び準都市計画区域の場合は使用しない。 |
| urf:areaClassificationType     [ _UrbanFunction ] | gml::CodeType [0..1] | 指定された区域 が属する区域区分。コードリスト（Common_areaClassificationType.xml）より選択する。都市計画区域、準都市計画区域及び区域区分の場合は使用しない。 |
| urf:nominalArea     [ _UrbanFunction ] | gml::MeasureType [0..1] | 公式の面積。単位はha（uom=”ha”）とする。区域の総面積とし、区域が複数の市区町村に跨っている場合は合計とする。 |
| urf:prefecture     [ _UrbanFunction ] | gml::CodeType [0..1] | 区域が所在する都道府県の都道府県コード。コードリスト（Common_localPublicAuthorities.xml）より選択する。 |
| urf:city     [ _UrbanFunction ] | gml::CodeType [0..1] | 区域が所在する市区町村の市区町村コード。コードリスト（Common_localPublicAuthorities.xml）より選択する。 |
| urf:reference     [ _UrbanFunction ] | xs::anyURI [0..1] | 外部の参照情報。 |
| urf:reason     [ _UrbanFunction ] | gml:StringOrRefType [0..1] | 区域が指定された理由。 |
| urf:note     [ _UrbanFunction ] | gml:StringOrRefType [0..1] | その他特筆すべき事項。 |
| urf:surveyYear     [ _UrbanFunction ] | xs::gYear [0..1] | 調査が実施された年（西暦）。都市計画決定情報、区域及び災害リスク（土砂災害）の場合は対象外とする。 |
| urf:location     [ Zone ] | xs::string [0..1] | 区域の位置を示す名称。 [記述例] 新潟県長岡市大字宮本東方町、大字高頭町、大字深沢町、大字親沢町及び大字大積町一丁目 |
| urf:objectives     [ _AbstractDistrictPlan ] | gml:StringOrRefType [0..1] | 当該地区計画等の目標。 |
| urf:policy     [ _AbstractDistrictPlan ] | gml:StringOrRefType [0..1] | 当該区域の整備、開発及び保全に関する方針。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:zonalDisasterPreventionFacilitiesAllocation     [ DisasterPreventionBlockImprovementZonePlan ] | gml:StringOrRefType [0..1] | 地区防災施設の区域。 |
| urf:specifiedZonalDisasterPreventionFacilitiesAllocation     [ DisasterPreventionBlockImprovementZonePlan ] | gml:StringOrRefType [0..1] | 特定地区防災施設の区域。 |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| core:externalReference             [ _CityObject ] | core:ExternalReference [0..*] | 外部への参照。標準製品仕様書では使用しない。 |
| gen:dateAttribute             [ _CityObject ] | gen:dateAttribute [0..*] | 日付型属性。属性を追加したい場合に使用する。 |
| gen:doubleAttribute             [ _CityObject ] | gen:doubleAttribute [0..*] | 実数型属性。属性を追加したい場合に使用する。 |
| gen:genericAttributeSet             [ _CityObject ] | gen:genericAttributeSet [0..*] | 汎用属性のセット。属性を追加したい場合に使用する。 |
| gen:intAttribute             [ _CityObject ] | gen:intAttribute [0..*] | 整数型属性。属性を追加したい場合に使用する。 |
| gen:measureAttribute             [ _CityObject ] | gen:measureAttribute [0..*] | 単位付き数値型属性。属性を追加したい場合に使用する。 |
| gen:stringAttribute             [ _CityObject ] | gen:stringAttribute [0..*] | 文字列型属性。属性を追加したい場合に使用する。 |
| gen:uriAttribute             [ _CityObject ] | gen:uriAttribute [0..*] | URI型属性。属性を追加したい場合に使用する。 |
| uro:pointCloud             [ _CityObject ] | uro:AbstractPointCloud [0..1] | ポイントクラウドへの参照。 |
| urf:lod-1MultiCurve             [ _UrbanFunction ] | gml:MultiCurve [0..1] | 地物の外周又は中心線（LOD-1）。標準製品仕様書では使用しない。 |
| urf:lod-1MultiPoint             [ _UrbanFunction ] | gml:MultiPoint [0..1] | 地物の位置（LOD-1）。標準製品仕様書では使用しない。 |
| urf:lod-1MultiSurface             [ _UrbanFunction ] | gml:MultiSurface [0..1] | 広がりを持った範囲（LOD-1）。都道府県レベルでの表現に使用する。標準製品仕様書では使用しない。 |
| urf:lod-2MultiCurve             [ _UrbanFunction ] | gml:MultiCurve [0..1] | 地物の外周又は中心線（LOD-2）。標準製品仕様書では使用しない。 |
| urf:lod-2MultiPoint             [ _UrbanFunction ] | gml:MultiPoint [0..1] | 地物の位置（LOD-2）。標準製品仕様書では使用しない。 |
| urf:lod-2MultiSurface             [ _UrbanFunction ] | gml:MultiSurface [0..1] | 広がりを持った範囲（LOD-2）。国レベルでの表現に使用する。標準製品仕様書では使用しない。 |
| urf:lod0MultiCurve             [ _UrbanFunction ] | gml:MultiCurve [0..1] | 地物の外周又は中心線（LOD0）。標準製品仕様書では使用しない。 |
| urf:lod0MultiPoint             [ _UrbanFunction ] | gml:MultiPoint [0..1] | 地物の位置（LOD0）。標準製品仕様書では使用しない。 |
| urf:lod0MultiSurface             [ _UrbanFunction ] | gml:MultiSurface [0..1] | 広がりを持った範囲（LOD0）。標準製品仕様書では使用しない。 |
| urf:lod1MultiCurve             [ _UrbanFunction ] | gml:MultiCurve [0..1] | 地物の外周又は中心線（LOD1）。標準製品仕様書では使用しない。 |
| urf:lod1MultiSurface             [ _UrbanFunction ] | gml:MultiSurface [0..1] | 面的な範囲。高さを0とする。区域の場合は、計画図に示す区域の境界線に囲まれた平面的な範囲を指す。 |
| urf:target             [ _UrbanFunction ] | core:_CityObject [0..*] | 都市オブジェクトへの参照。標準製品仕様書では使用しない。 |
| uro:dataQualityAttribute             [ _UrbanFunction ] | uro:DataQualityAttribute [0..1] | 作成したデータの品質に関する情報。必須とする。 |
| uro:keyValuePairAttribute             [ _UrbanFunction ] | uro:KeyValuePairAttribute [0..*] | コード属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 |
| urf:boundary             [ Zone ] | urf:Boundary [0..*] | 区域を構成する境界。区域の外周及び内周となる境界線。境界線の詳細な情報を記述したい場合にのみ作成する。 |
| urf:urbanParkAttribute             [ Zone ] | urf:UrbanParkAttribute [0..1] | 都市公園に関するデータ型を保持する。区域が都市公園の区域の場合に使用する。 |
| urf:districtDevelopmentPlan             [ _AbstractDistrictPlan ] | urf:DistrictDevelopmentPlan [0..*] | 都市計画法第12条の5第2項第1号に定める当該地区計画等に定められた地区整備計画。 |
| urf:promotionDistrict             [ _AbstractDistrictPlan ] | urf:PromotionDistrict [0..*] | 当該地区計画等に促進区を定める場合の促進区。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| urf:zonalDisasterPreventionFacilities             [ DisasterPreventionBlockImprovementZonePlan ] | urf:ZonalDisasterPreventionFacility [0..*] | 防災街区整備地区計画に計画された地区防災施設及び特定地区防災施設。 |

