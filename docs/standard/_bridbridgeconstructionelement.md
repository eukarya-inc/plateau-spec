# 4.11.3.1.3 brid:BridgeConstructionElement

## 4.11.3.1.3 brid:BridgeConstructionElement

**表 4-506**

| 型の定義 | 橋梁の構造上重要な部材。橋脚、橋台、トラス、アーチ、吊材、パイロン、ケーブルをさす。![Image](images/_bridbridgeconstructionelement_img001.png)brid:BridgeConstructionElementの例 | ← |
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
| brid:class     [ BridgeConstructionElement ] | gml::CodeType [0..1] | 部材の形態による区分。標準製品仕様書では使用しない。 |
| brid:function     [ BridgeConstructionElement ] | gml::CodeType [0..*] | 部材の主たる働き。コードリスト （ BridgeConstructionElement_function.xml）より選択する。 |
| brid:usage     [ BridgeConstructionElement ] | gml::CodeType [0..*] | 部材の主な使い道。標準製品仕様書では使用しない。 |
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
| brid:boundedBy             [ BridgeConstructionElement ] | brid:_BoundarySurface [0..*] | 部材を構成する外壁、屋根等の境界面への参照。部材の境界面が橋梁（brid:Bridge又はbrid:BridgePart）の境界面となる場合にのみ作成する。 |
| brid:lod1Geometry             [ BridgeConstructionElement ] | gml:_Geometry [0..1] | 部材のLOD1 の形状。標準製品仕様書では使用しない。 |
| brid:lod2Geometry             [ BridgeConstructionElement ] | gml:_Geometry [0..1] | 部材のLOD2 の形状。 gml:MultiSurface により記述することを基本とする。 構造物の外形（外側から見える形）を構成する面を取得する。 面の各頂点に構造物の高さを与える。 容積の算出等ユースケースで必要な場合は、gml:Solid を使用する。 |
| brid:lod3Geometry             [ BridgeConstructionElement ] | gml:_Geometry [0..1] | 部材のLOD3 の形状。 gml:MultiSurfaceにより記述することを基本とする。 構造物の外形（外側から見える形）を構成する面を取得する。 面の各頂点に構造物の高さを与える。 容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。 |
| brid:lod4Geometry             [ BridgeConstructionElement ] | gml:_Geometry [0..1] | 部材のLOD4の形状。 gml:MultiSurfaceにより記述することを基本とする。 構造物の外形（外側から見える形）を構成する面を取得する。 面の各頂点に構造物の高さを与える。 容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。 |

