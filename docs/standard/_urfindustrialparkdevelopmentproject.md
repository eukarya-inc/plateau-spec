# 4.10.3.8.4 urf:IndustrialParkDevelopmentProject

## 4.10.3.8.4 urf:IndustrialParkDevelopmentProject

**表 4-416**

| 型の定義 | 都市計画法第12条第1項第3号に定める事業。首都圏の近郊整備地帯及び都市開発区域の整備に関する法律（昭和三十三年法律第九十八号）による工業団地造成事業又は近畿圏の近郊整備区域及び都市開発区域の整備及び開発に関する法律（昭和三十九年法律第百四十五号）による工業団地造成事業。属性urf:functionは、都市計画法第12条第2項に定める市街地開発事業の種類とする。コードリスト（UrbanDevelopmentProject_function.xml）より選択する。多重度は任意となっているが、運用上必須とする。属性urf:usageは使用しない。 | ← |
| --- | --- | --- |
| 上位の型 | urf:UrbanDevelopmentProject | ← |
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
| urf:scheduledExecutor     [ UrbanDevelopmentProject ] | xs::string [0..1] | 都市計画法第12条第5項に定める事業の実施予定機関の名称。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:publicFacilityAllocation     [ IndustrialParkDevelopmentProject ] | gml:StringOrRefType [1..1] | 首都圏の近郊整備地帯及び都市開発区域の整備に関する法律第5条第1項又は近畿圏の近郊整備区域及び都市開発区域の整備及び開発に関する法律第7条第1項に定める公共施設の配置及び規模。 |
| urf:residentialLandUsePlan     [ IndustrialParkDevelopmentProject ] | gml:StringOrRefType [1..1] | 首都圏の近郊整備地帯及び都市開発区域の整備に関する法律第5条第1項又は近畿圏の近郊整備区域及び都市開発区域の整備及び開発に関する法律第7条第1項に定める宅地の利用計画。 |
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

