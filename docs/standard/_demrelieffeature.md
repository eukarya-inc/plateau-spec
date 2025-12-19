# 4.18.3.1.1 dem:ReliefFeature

## 4.18.3.1.1 dem:ReliefFeature

**表 4-776**

| 型の定義 | 地形の凹凸を表現した地物。基準地域メッシュ（第三次地域区画、一辺の長さ約1km）を地物の単位とする（基準地域メッシュの境界で区切る）。 | ← |
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
| dem:lod     [ ReliefFeature ] | core::integerBetween0and4 [1..1] | この地形に適用されるLOD0～3までの半角数字のいずれかとする。LOD1の場合は1となる。 dem:ReliefFeatureがもつdem:TINRelief、dem:MassPointRelief又はdem:BreaklineReliefのlodと一致させる。 |
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
| dem:reliefComponent             [ ReliefFeature ] | dem:_ReliefComponent [1..*] | 同じ基準地域メッシュ（第3次地域区画）に限り、dem:TINRelief又はdem:MassPointReliefのいずれか1つを複数もつことができる。 |
| uro:demDataQualityAttribute             [ ReliefFeature ] | uro:DataQualityAttribute [1..1] | 作成されたデータの品質に関する情報。必須とする。 |
| uro:demKeyValuePairAttribute             [ ReliefFeature ] | uro:KeyValuePairAttribute [0..*] | コード属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 |

