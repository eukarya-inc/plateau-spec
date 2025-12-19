# 4.11.3.1.4 brid:BridgeRoom

## 4.11.3.1.4 brid:BridgeRoom

**表 4-507**

| 型の定義 | 壁、間仕切り、床、天井などで仕切られた、橋梁内部に設けられた区画。橋梁にこれと一体となった建屋が存在し、その内部空間を表現する場合に使用する。  brid:BridgeRoomは橋梁内部の区画を区切る境界面（brid:_BoundarySurfaceの下位型）及びこの境界面の開口部（brid:_Openingの下位型）、brid:BridgeRoomに付属する固定的な設備（brid:IntBridgeInstallation）及び、brid:BridgeRoomに設置された可動設備（brid:BridgeFurniture）の集まりからなる。 | ← |
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
| brid:class     [ BridgeRoom ] | gml::CodeType [0..1] | 区画の形態による区分。コードリスト（Room_class.xml）より選択する。 |
| brid:function     [ BridgeRoom ] | gml::CodeType [0..*] | 区画の主たる働き。コードリスト（Room_function.xml）より選択する。 |
| brid:usage     [ BridgeRoom ] | gml::CodeType [0..*] | 区画の主な使い道。標準製品仕様書では使用しない。 |
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
| brid:boundedBy             [ BridgeRoom ] | brid:_BoundarySurface [0..*] | 区画の主要構造の外形を示す境界面。境界面は、内壁面（brid:InteriorWallSurface）、天井面（brid:CeilingSurface）、床面（brid:FloorSurface）又は閉鎖面（brid:ClosureSurface）のいずれかでなければならない。 |
| brid:bridgeRoomInstallation             [ BridgeRoom ] | brid:IntBridgeInstallation [0..*] | 区画に設置された固定設備（brid:IntBridgeInstallation）。 |
| brid:interiorFurniture             [ BridgeRoom ] | brid:BridgeFurniture [0..*] | 区画に設置された可動設備（brid:BridgeFurniture）。 |
| brid:lod4MultiSurface             [ BridgeRoom ] | gml:MultiSurface [0..1] | 区画の主要構造の外形を示す面の集まり。 brid:BridgeRoomの形状は、brid:lod4Solidにより記述することを基本とするが、境界面により立体の境界を閉じることが出来ない場合には、brid:lod4MultiSurfaceにより記述することを可とする。  gml:MultiSurfaceを構成するgml:Polygonは、以下のいずれの地物のLOD4幾何オブジェクトに含まれなければならない。brid:boundedByによりこのbrid:BridgeRoomが参照する境界面（brid:_BoundarySurface）及びその開口部（brid:_Opening）brid:bridgeRoomInstallationによりこのbrid:BridgeRoomが参照する内部付属物（brid:IntBridgeInstallation）の境界面及びその開口部 |
| brid:lod4Solid             [ BridgeRoom ] | gml:Solid [0..1] | 区画の主要構造の外形を示す立体。 brid:lod4Solid 又はbrid:lod4MultiSurface のいずれかを必須とするが、brid:lod4Solidにより記述することを基本とする。  gml:Solidを構成する境界面のgml:Polygonは、以下のいずれの地物のLOD4幾何オブジェクトに含まれなければならない。brid:boundedByによりこのbrid:BridgeRoomが参照する境界面（brid:_BoundarySurface）及びその開口部（brid:_Opening）brid:bridgeRoomInstallationによりこのbrid:BridgeRoomが参照する内部付属物（brid:IntBridgeInstallation）の境界面及びその開口部 |

