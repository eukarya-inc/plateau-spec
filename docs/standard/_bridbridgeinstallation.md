# 4.11.3.1.16 brid:BridgeInstallation

## 4.11.3.1.16 brid:BridgeInstallation

**表 4-519**

| 型の定義 | 橋梁の外部付属物。橋梁の外側に設置され、橋梁の外観を特徴づける設備。橋梁の付帯的な設備であり、主要な部分であってはならない。また、橋梁（brid:Bridge又はbrid:BridgePart）と接していなければならない。 橋梁の外部付属物には以下を含む。ただし、全て外部に設置され、橋梁と接するもののみを対象とする。 支承、落橋防止装置、伸縮装置、排水施設、高欄、防護柵、遮音壁、遮光壁、航空障害灯など。 ただし、ユースケースの要求に応じて、取得対象とする付属物を限定してもよく、また、付属物として取得せず橋梁の一部として取得してもよい。 | ← |
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
| brid:class     [ BridgeInstallation ] | gml::CodeType [0..1] | 外部付属物の形態による区分。 |
| brid:function     [ BridgeInstallation ] | gml::CodeType [0..*] | 外部付属物の主たる働き。コードリスト（BridgeInstallation_function.xml）より選択する。 |
| brid:usage     [ BridgeInstallation ] | gml::CodeType [0..*] | 付属物の主な使い道。コードリスト（BridgeInstallation_usage.xml）より選択する。この属性を使用する場合は、コードリストを作成する。 |
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
| brid:boundedBy             [ BridgeInstallation ] | brid:_BoundarySurface [0..*] | 外部付属物を構成する外壁、屋根等の境界面への参照。付属物の境界面が橋梁（brid:Bridge又はbrid:BridgePart）の境界面となる場合にのみ作成する。橋梁の空間に付属物を含まない場合は、付属物を構成する面を、境界面（brid:_BoundarySurface）に区別する必要はない。 |
| brid:lod2Geometry             [ BridgeInstallation ] | gml:_Geometry [0..1] | 外部付属物のLOD2の形状。 gml:MultiSurfaceにより記述することを基本とする。 付属物の外形（外側から見える形）を構成する面を取得し、面の各頂点に屋外付属物の高さを与える。 容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。 |
| brid:lod3Geometry             [ BridgeInstallation ] | gml:_Geometry [0..1] | 外部付属物のLOD3の形状。 gml:MultiSurfaceにより記述することを基本とする。 付属物の外形（外側から見える形）を構成する面を取得し、面の各頂点に屋外付属物の高さを与える。 容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。 |
| brid:lod4Geometry             [ BridgeInstallation ] | gml:_Geometry [0..1] | 外部付属物のLOD4の形状。 gml:MultiSurfaceにより記述することを基本とする。 付属物の外形（外側から見える形）を構成する面を取得し、面の各頂点に屋外付属物の高さを与える。 容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。 |

