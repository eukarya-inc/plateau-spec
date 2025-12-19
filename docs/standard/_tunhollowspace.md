# 4.12.3.1.3 tun:HollowSpace

## 4.12.3.1.3 tun:HollowSpace

**表 4-564**

| 型の定義 | トンネルの内空。 tun:HollowSpaceはトンネルの内部空間を示す立体であり、輸送に使用する区画や管理点検に使用する区画等に分けることができる。各区画の立体の境界面（tun:_BoundarySurfaceの下位型）及びこの境界面の開口部（tun:_Openingの下位型）、tun:HollowSpaceに付属する固定的な設備（tun:IntTunnelInstallation）及び、tunHollowSpaceに設置された可動設備（tunTunnelFurniture）の集まりからなる。 | ← |
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
| tun:class     [ HollowSpace ] | gml::CodeType [0..1] | 区画の形態による区分。コードリストから選択する。この属性を使用する場合は、コードリスト（HollowSpace_class.xml）を作成すること。 |
| tun:function     [ HollowSpace ] | gml::CodeType [0..*] | 区画の主たる働き。コードリストから選択する。この属性を使用する場合は、コードリスト（HollowSpace_function.xml）を作成すること。 |
| tun:usage     [ HollowSpace ] | gml::CodeType [0..*] | 区画の主な使い道。標準製品仕様書では使用しない。 |
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
| tun:boundedBy             [ HollowSpace ] | tun:_BoundarySurface [0..*] | 区画の主要構造の外形を示す境界面。境界面は、壁面（tun:InteriorWallSurface）、天井面（tun:CeilingSurface）、床面（tun:FloorSurface）又は閉鎖面（tun:ClosureSurface）のいずれかでなければならない。 壁面と天井面との区分が構造上難しい場合は、内壁面（tun:InteriorWallSurface）として取得することを基本とする。 |
| tun:hollowSpaceInstallation             [ HollowSpace ] | tun:IntTunnelInstallation [0..*] | 区画に設置された固定設備。 |
| tun:interiorFurniture             [ HollowSpace ] | tun:TunnelFurniture [0..*] | 区画に設置された可動設備。 |
| tun:lod4MultiSurface             [ HollowSpace ] | gml:MultiSurface [0..1] | 区画の主要構造の外形を示す面の集まり。標準製品仕様書では使用しない。 |
| tun:lod4Solid             [ HollowSpace ] | gml:Solid [0..1] | gml:Solidを構成する境界面のgml:Polygonは、以下のいずれかの地物のLOD4幾何オブジェクトに含まれなければならない。tun:boundedByによりこのtunHollowSpaceが参照する境界面（tun:_BoundarySurface）及びその開口部（tun:_Opening）tun:hollowSpaceInstallationによりこのtun:HollowSpaceが参照する内部付属物（tunIntTunnelInstallation）の境界面及びその開口部 |

