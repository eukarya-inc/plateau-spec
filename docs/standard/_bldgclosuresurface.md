# 4.2.3.1.9 bldg:ClosureSurface

## 4.2.3.1.9 bldg:ClosureSurface

**表 4-36**

| 型の定義 | 建築物の立体又は部屋の立体を構成するために仮想的に設ける閉鎖面。一棟の建築物を、主題属性の異なる複数の部分に分ける場合に、その境界面として使用する。![Image](images/_bldgclosuresurface_img001.png)LOD2又はLOD3でのbldg:ClosureSurfaceの例屋内においては、境界面となる内壁面や天井面、床面はないが、建築確認申請では部屋となっている空間を区切る場合に、部屋の境界面として便宜上設けられた仮想的な面をさす。![Image](images/_bldgclosuresurface_img002.png)LOD4でのbldg:ClosureSurfaceの例![Image](images/_bldgclosuresurface_img003.png)LOD4でのbldg:ClosureSurfaceを非表示にした例 | ← |
| --- | --- | --- |
| 上位の型 | bldg:_BoundarySurface | ← |
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
| bldg:lod2MultiSurface             [ _BoundarySurface ] | gml:MultiSurface [0..1] | 建築物モデル（LOD2）において境界面の形状・起伏を再現した面。取得基準及び取得方法は、建築物のLOD定義に従う。 |
| bldg:lod3MultiSurface             [ _BoundarySurface ] | gml:MultiSurface [0..1] | 建築物モデル（LOD3）において境界面の形状・起伏を再現した面。取得基準及び取得方法は、建築物のLOD定義に従う。 |
| bldg:lod4MultiSurface             [ _BoundarySurface ] | gml:MultiSurface [0..1] | 建築物モデル（LOD4）において境界面の形状・起伏を再現した面。取得基準及び取得方法は、建築物のLOD定義に従う。  bldg:ClosureSurfaceの場合は、BuildingPartと連続する他のBuildingPartとの境界線により囲まれた面又は内壁面、天井面若しくは床面が無いが建築確認申請上部屋として区分されている空間を区切る面。 |
| bldg:opening             [ _BoundarySurface ] | bldg:_Opening [0..*] | 境界面に設置される、窓や扉への参照。LOD3又はLOD4の空間属性。 LOD3又はLOD4の空間属性をもつ場合のみ開口部への参照を作成できる。  bldg:ClosureSurfaceの場合、本関連役割は使用しない。 |
| uro:ifcBoundarySurfaceAttribute             [ _BoundarySurface ] | uro:IfcAttribute [0..*] | IDM・MVDで定義されるIFCのクラス及びプロパティセットに含まれる情報。bldg:RoofSurfaceに付与可能なデータ型は以下とする。uro:IfcBuildingElement uro:IfcRoofこのとき、uro:IfcBuildingElement の属性uro:elementTypeの値はIfcSlab又はIfcRoofとなる。bldg:WallSurfaceに付与可能なデータ型は以下とする。uro:IfcWall uro:IfcWallStandardCase  uro:IfcCurtainWallこのとき、uro:IfcBuildingElement の属性uro:elementTypeの値はIfcWall、IfcWallStanadrdCase又はIfcCurtainWallとなる。bldg:GroundSurfaceに付与可能なデータ型は、以下とする。uro:IfcBuildingElementこのとき、uro:IfcBuildingElement の属性uro:elementTypeの値はIfcSlabとなる。bldg:OuterFloorSurface に付与可能なデータ型は以下とする。uro:IfcBuildingElementこのとき、uro:IfcBuildingElement の属性uro:elementType の値はIfcSlab となる。bldg:OuterCeilingSurface、bldg:InteriorWallSurface、bldg:CeilingSurface、bldg:FloorSurface及びbldg:ClosureSurfaceの場合、本関連役割は使用しない。 |
| uro:indoorBoundarySurfaceAttribute             [ _BoundarySurface ] | uro:IndoorAttribute [0..*] | 屋内ナビゲーションに必要な情報。以下の地物型で使用する。 bldg:InteriorWallSurface  bldg:CeilingSurface  bldg:FloorSurface なお、uro:IndoorAttributeは抽象クラスであるため、それぞれの地物型はこれを継承する具象クラスを使用すること。 |

