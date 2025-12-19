# 4.2.3.1.14 bldg:Door

## 4.2.3.1.14 bldg:Door

**表 4-41**

| 型の定義 | 採光、通風、換気、眺望、通行などの目的のため、建築物の屋根、天井、壁、床などに設けられた開口部のうち、人や物の出入りを目的とするもの。![Image](images/_bldgdoor_img001.png)bldg:Doorの例CityGMLでは、扉を面として表現し、一つの扉を外側と内側の二つのbldg:Doorのオブジェクトとして表現する。例えば、屋内と屋外をつなぐ窓があった場合、 外側となるbldg:Doorは、建築物の外壁面（bldg:WallSurface）等の境界面に含まれる。 内側となるbldg:Doorは、部屋の内壁面（bldg:InteriorWallSurface）等の境界面に含まれる。 このとき、屋外の境界面（bldg:WallSurface、bldg:GroundSurface、bldg:OuterFloorSurface、bldg:OuterCeilingSurface）に設けられた開口部は、常にその法線ベクトルが建築物の外側を向く。部屋の境界面（bldg:InteriorWallSurface、bldg:FloorSurface、bldg:CeilingSurface）に設けられた開口部は、常にその法線ベクトルが部屋の内側を向く。 | ← |
| --- | --- | --- |
| 上位の型 | bldg:_Opening | ← |
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
| bldg:lod3MultiSurface             [ _Opening ] | gml:MultiSurface [0..1] | 開口部の外周に囲まれた面。必須とする。 |
| bldg:lod4MultiSurface             [ _Opening ] | gml:MultiSurface [0..1] | 開口部の外周に囲まれた面。必須とする。 |
| uro:ifcOpeningAttribute             [ _Opening ] | uro:IfcAttribute [0..*] | DM・MVDで定義されるIFCのクラス及びプロパティセットに含まれる情報。bldg:Windowに付与可能なデータ型は、以下とする。uro:IfcOpeningElement uro:IfcWindow  uro:IfcPsetOpeningElementCommon  uro:IfcPsetWindowCommonbldg:Doorに付与可能なデータ型は、以下とする。uro:IfcOpeningElement uro:IfcDoor  uro:IfcPsetOpeningElementCommon  uro:IfcPsetDoorCommon |
| uro:indoorOpeningAttribute             [ _Opening ] | uro:IndoorAttribute [0..*] | 屋内ナビゲーションに必要な情報。bldg:Windowに付与可能なデータ型は以下とする。uro:IndoorZoneAttribute uro:IndoorUserDefinedAttributebldg:Doorに付与可能なデータ型は以下とする。uro:IndoorZoneAttribute uro:IndoorUserDefinedAttribute |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| bldg:address             [ Door ] | core:Address [0..*] | 扉に紐づけれらた住所。 CityGMLでは複数個の記述が可能（多重度[0..*]）であるが、標準製品仕様書では、最大1個とする。 |

