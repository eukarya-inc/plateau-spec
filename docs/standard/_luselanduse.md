# 4.8.3.1.1 luse:LandUse

## 4.8.3.1.1 luse:LandUse

**表 4-299**

| 型の定義 | 都市計画基礎調査の土地利用現況。 | ← |
| --- | --- | --- |
| 上位の型 | core:_CityObject | ← |
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
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| luse:class     [ LandUse ] | gml::CodeType [0..1] | 土地利用用途の大まかな区分。土地利用用途の区分は、都市計画基礎調査実施要領（国土交通省都市局）による区分とする。コードリスト（ Common_landUseType.xml）より選択する。 |
| luse:function     [ LandUse ] | gml::CodeType [0..*] | 土地利用の機能。標準製品仕様書では使用しない。 |
| luse:usage     [ LandUse ] | gml::CodeType [0..*] | 土地利用の用途。標準製品仕様書では使用しない。 |
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
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| luse:lod1MultiSurface             [ LandUse ] | gml:MultiSurface [0..1] | 土地利用が変化する境界により囲われた同一の土地利用の範囲。 |
| uro:ifcLandUseAttribute             [ LandUse ] | uro:IfcAttribute [0..*] | 標準製品仕様書では使用しない。 |
| uro:landUseDetailAttribute             [ LandUse ] | uro:LandUseDetailAttribute [0..1] | 土地利用現況調査により得られた土地の詳細情報。 |
| uro:luseDataQualityAttribute             [ LandUse ] | uro:DataQualityAttribute [1..1] | 作成したデータの品質に関する情報。必須とする。 |
| uro:luseDmAttribute             [ LandUse ] | uro:DmAttribute [0..*] | 標準製品仕様書では使用しない。 |
| uro:luseFacilityAttribute             [ LandUse ] | uro:FacilityAttribute [0..*] | uro:luseFacilityTypeAttribute.classによって指定された分野における施設管理情報。 |
| uro:luseFacilityIdAttribute             [ LandUse ] | uro:FacilityIdAttribute [0..1] | uro:luseFacilityTypeAttribute.classによって指定された分野における施設の識別情報。 |
| uro:luseFacilityTypeAttribute             [ LandUse ] | uro:FacilityTypeAttribute [0..*] | 特定分野における施設の分類情報。 |
| uro:luseKeyValuePairAttribute             [ LandUse ] | uro:KeyValuePairAttribute [0..*] | 属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 |

