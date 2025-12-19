# 4.2.3.1.3 bldg:Room

## 4.2.3.1.3 bldg:Room

**表 4-30**

| 型の定義 | 壁、間仕切り、床、天井などで仕切られ、生活の場などに用いられる、建物内部の隔てられた空間の区画（部屋）。![Image](images/_bldgroom_img001.png)bldg:Roomの例bldg:Roomは、bldg:Buildingに含まれる地物として記述する。このとき、bldg:Roomは、複数の地物の集まりとして表現する。bldg:Roomに含まれる地物とは、以下である。 部屋を区切る境界面（bldg:_BoundarySurfaceの下位型） 部屋に付属する固定的な設備（bldg:InteriorBuildingInstallation） 部屋の中に設置された移動可能な家具（bldg:BuildingFurniture） さらに、部屋を区切る境界面及び部屋に付属する固定的な設備は、開口部（bldg:_Opening）の下位型を含むことができる。 | ← |
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
| bldg:class     [ Room ] | gml::CodeType [0..1] | 部屋の形態による区分。コードリスト（ Room_class.xml）より選択する。 |
| bldg:function     [ Room ] | gml::CodeType [0..*] | 部屋の主たる働き。コードリスト（ Room_function.xml）より選択する。 |
| bldg:usage     [ Room ] | gml::CodeType [0..*] | 部屋の主な使い道。標準製品仕様書では使用しない。 |
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
| bldg:boundedBy             [ Room ] | bldg:_BoundarySurface [0..*] | 部屋の外形を示す境界面。境界面は、内壁面（bldg:InteriorWallSurface）、天井面（bldg:CeilingSurface）、床面（bldg:FloorSurface）又は閉鎖面（bldg:ClosureSurface）のいずれかでなければならない。 |
| bldg:interiorFurniture             [ Room ] | bldg:BuildingFurniture [0..*] | 部屋に設置された移動可能な家具（bldg:BuildingFurniture）。 |
| bldg:lod4MultiSurface             [ Room ] | gml:MultiSurface [0..1] | 部屋の主要構造の外形を示す面の集まり。 gml:MultiSurfaceを構成するgml:Polygonは、以下のいずれかの地物のLOD4幾何オブジェクトに含まれなければならない。境界面（bldg:_BoundarySurface）及びその開口部（bldg:_Opening）：ただし、境界面は、このbldg:Roomが、関連役割bldg:boundedByにより参照する境界面であること。また、開口部は、その境界面に包含されていること。屋内付属物（bldg:IntBuildingInstallation）の境界面及びその開口部：ただし、屋内付属物は、このbldg:Roomが、関連役割bldg:roomInstallationにより参照する付属物であること。また、開口部はその付属物に包含されていること。bldg:lod4Solidを作成しない場合は、bldg:lod4MultiSurfaceを必ず作成する。 |
| bldg:lod4Solid             [ Room ] | gml:_Solid [0..1] | 部屋の外形を示す立体。 gml:Solidを構成するgml:Polygonは、以下のいずれかの地物のLOD4幾何オブジェクトに含まれなければならない。境界面（bldg:_BoundarySurface）及びその開口部（bldg:_Opening）ただし、境界面は、このbldg:Roomが、関連役割bldg:boundedByにより参照する境界面であること。また、開口部は、その境界面に包含されていること。屋内付属物（bldg:IntBuildingInstallation）の境界面及びその開口部ただし、屋内付属物は、このbldg:Roomが、関連役割bldg:roomInstallationにより参照する付属物であること。また、開口部はその付属物に包含されていること。bldg:lod4Solid 又はbldg:lod4MultiSurface のいずれかを必須とするが、bldg:lod4Solidにより記述することを基本とする。 |
| bldg:roomInstallation             [ Room ] | bldg:IntBuildingInstallation [0..*] | 部屋に設置された屋内付属物（bldg:IntBuildingInstallation）。 |
| uro:ifcRoomAttribute             [ Room ] | uro:IfcAttribute [0..*] | IDM・MVDで定義されるIFCのクラス及びプロパティセットに含まれる情報。bldg:Roomに付与可能なデータ型は以下とする。uro:IfcPsetSpaceCommon uro:IfcSpace  uro:IfcSpaceBaseQuantity  uro:IfcClassificationReference |
| uro:indoorRoomAttribute             [ Room ] | uro:IndoorAttribute [0..*] | 屋内ナビゲーションに必要な情報。bldg:Roomに付与可能なデータ型は以下とする。 uro:IndoorSpaceAttribute  uro:IndoorZoneAttribute  uro:IndoorUserDefinedAttribute |

