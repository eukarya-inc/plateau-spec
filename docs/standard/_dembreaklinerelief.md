# 4.18.3.1.4 dem:BreaklineRelief

## 4.18.3.1.4 dem:BreaklineRelief

**表 4-779**

| 型の定義 | 線の集まりにより地形の起伏を表現した地物。基準地域メッシュ（第三次地域区画、一辺の長さ約1km）を地物の単位とする。![Image](images/_dembreaklinerelief_img001.png)dem:BreaklineReliefの例 | ← |
| --- | --- | --- |
| 上位の型 | dem:_ReliefComponent | ← |
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
| dem:lod     [ _ReliefComponent ] | core::integerBetween0and4 [1..1] | この地形に適用されるLOD0～3までの半角数字のいずれかとする。LOD1の場合は1となる。 dem:ReliefFeatureがもつdem:TINRelief、dem:MassPointRelief又はdem:BreaklineReliefのlodと一致させる。 |
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
| dem:extent             [ _ReliefComponent ] | gml:Polygon [0..1] | 空間範囲。extentのexteriorとして、地形の外形を多角形で記述し、extentのinteriorは地形の内空を記述する。 |
| uro:demDmAttribute             [ _ReliefComponent ] | uro:DmAttribute [0..*] | 公共測量標準図式による表現に必要な情報。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| dem:breaklines             [ BreaklineRelief ] | gml:MultiCurve [0..1] | 抑止線を表現する線の集まり。標準製品仕様書では使用しない。 |
| dem:ridgeOrValleyLines             [ BreaklineRelief ] | gml:MultiCurve [0..1] | 谷や尾根を表現する線の集まり。 |

