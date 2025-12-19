# 4.12.3.1.12 tun:FloorSurface

## 4.12.3.1.12 tun:FloorSurface

**表 4-573**

| 型の定義 | トンネルの内部空間の下面に位置する水平で平らな板状の構造物（床面）。 tun:FloorSurfaceの法線ベクトルは上向きとなる。 | ← |
| --- | --- | --- |
| 上位の型 | tun:_BoundarySurface | ← |
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
| tun:lod2MultiSurface             [ _BoundarySurface ] | gml:MultiSurface [0..1] | トンネルでのLOD2における形状・起伏を再現した面。詳細な取得基準は、各地物型のLOD定義に従うこと。 |
| tun:lod3MultiSurface             [ _BoundarySurface ] | gml:MultiSurface [0..1] | トンネルでのLOD3における形状・起伏を再現した面。詳細な取得基準は、各地物型のLOD定義に従うこと。 |
| tun:lod4MultiSurface             [ _BoundarySurface ] | gml:MultiSurface [0..1] | トンネルでのLOD4における形状・起伏を再現した面。詳細な取得基準は、各地物型のLOD定義に従うこと。 |
| tun:opening             [ _BoundarySurface ] | tun:_Opening [0..*] | 境界面に設置される開口部（窓及び扉）への参照。 tun:ClosureSurfaceの場合、本関連役割は使用しない。 |

