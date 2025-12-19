# Q.2.2 使用可能な地物型とLOD

## Q.2.2 使用可能な地物型とLOD

地下街モデルは、LODごとに、使用すべき地物型やその空間属性が異なる。

**表 Q-2 — Q-1: 地下街記述のLOD対応条件**

| ID | `/req/ubld/1` |
| --- | --- |
| 主題 | 妥当な地下街データセットの要件 |
| 分類 | 要件分類Q-1: 妥当な地下街オブジェクト |
| 条文 | 地下街の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

地下街は、i-URのUrban Objectモジュールに定義されたuro:UndergroundBuildingを用いて
記述する。uro:UndergroundBuildingはCityGMLのBuildingモジュールに定義されたbldg:_AbstractBuildingを継承し、定義されている。そのため、uro:UndergroundBuildingは、建築物（bldg:Building）と同様に、境界面や開口部に区分することができる。

- 地下街モデル（LOD0）では、uro:UndergroundBuildingを使用し、地下街の形状を面で記述する。
- 地下街モデル（LOD1）では、uro:UndergroundBuildingを使用し、地下街の形状を立体で記述する。
- 地下街モデル（LOD2）では、uro:UndergroundBuildingを使用し、地下街の形状を立体で記述するとともに、uro:UndergroundBuildingの境界面を屋根面や外壁面等の地物型を用いて記述する。屋根面や外壁面はbldg:_BoundarySurfaceとして定義されている。bldg:_BoundarySurfaceは抽象地物であり、実装にはこれを継承する地物型（例：屋根面はbldg:RoofSurface）を使用する。
- 地下街モデル（LOD3）では、地下街モデル（LOD2）で使用可能な地物型に加え、境界面に存在する開口部を閉鎖面として区分する。さらに、屋根や壁に扉（bldg:Door）や窓（bldg:Window）を付けることができる。
また、地下街モデル（LOD2）及び地下街モデル（LOD3）では、bldg:BuildingPartを使用できる。bldg:BuildingPartは、一つの地下街を二つ以上の部分に分けた地下街の一部である。例えば、一つの地下街について一部の階数や用途が異なる場合や、複合的な施設から構成され施設ごとに都市計画基礎調査の情報が作成されている場合に、一つの地下街を複数に区切り、区切った部品ごとに階数や用途、都市計画基礎調査の情報といった属性を与えるために使用する。 
ただし、このとき、 uro:UndergroundBuildingを構成するbldg:BuildingPartとは、一体的な建築物（互いに接する）でなければならない。またLOD2及びLOD3では、地下街の外側に付属する階段やエレベーターといった設備を、bldg:BuildingInstallationを用いて記述できる。
- 地下街モデル（LOD4）では、地下街モデル（LOD3）で使用可能な地物型に加え、地下街の内部の空間を部屋（bldg:Room）に区分して表現する。各部屋は、天井面や床面といった境界面に区分するとともに、屋内に存在する固定的な設備（bldg:IntBuildingInstallation）や家具（bldg:BuildingFurniture）を表現することができる。

地下街の各LODにおいて使用可能な地物型とその空間属性を表Q-3に示す。

**表 Q-3 — 地下街モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | LOD4 | 適用 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| uro:UndergroundBuilding |  | ● | ● | ● | ● | ● |  |
| ↑ | lod0FootPrint |  |  |  |  |  |  |
| ↑ | lod0RoofEdge | ● |  |  |  |  |  |
| ↑ | lod1Solid |  | ● |  |  |  |  |
| ↑ | lod2Solid |  |  | ● |  |  |  |
| ↑ | lod3Solid |  |  |  | ● |  |  |
| ↑ | lod4Solid |  |  |  |  | ■ | Solid又はMultiSurfaceのいずれかとする。 |
| ↑ | lod4MultiSurface |  |  |  |  | ■ | ↑ |
| bldg:BuildingPart |  |  |  | ■ | ■ | ■ | 一棟の建築物を、属性の異なる複数の部分に分ける場合に必須とする。 |
| ↑ | lod1Solid |  |  |  |  |  |  |
| ↑ | lod2Solid |  |  | ■ |  |  |  |
| ↑ | lod3Solid |  |  |  | ■ |  |  |
| ↑ | lod4Solid |  |  |  |  | ■ | Solid又はMultiSurfaceのいずれかとする。 |
| ↑ | lod4MultiSurface |  |  |  |  | ■ | ↑ |
| bldg:Room |  |  |  |  |  | ● |  |
| ↑ | lod4Solid |  |  |  |  | ● |  |
| bldg:RoofSurface |  |  |  | ● | ● | ● |  |
| ↑ | lod2MultiSurface |  |  | ● |  |  |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:WallSurface |  |  |  | ● | ● | ● |  |
| ↑ | lod2MultiSurface |  |  | ● |  |  |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:GroundSurface |  |  |  | ● | ● | ● |  |
| ↑ | lod2MultiSurface |  |  | ● |  |  |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:OuterCeilingSurface |  |  |  |  |  |  | 対象外 |
| ↑ | lod2MultiSurface |  |  |  |  |  |  |
| ↑ | lod3MultiSurface |  |  |  |  |  |  |
| ↑ | lod4MultiSurface |  |  |  |  |  | ↑ |
| bldg:OuterFloorSurface |  |  |  |  |  |  |  |
| ↑ | lod2MultiSurface |  |  |  |  |  | 対象外 |
| ↑ | lod3MultiSurface |  |  |  |  |  |  |
| ↑ | lod4MultiSurface |  |  |  |  |  | ↑ |
| bldg:ClosureSurface |  |  |  | ■ | ■ | ■ | BuildingPartを作成する場合は必須とする。 LOD4において、内壁面等はないが、建築確認申請では部屋となっている空間を区切る場合は必須とする。 |
| ↑ | lod2MultiSurface |  |  | ■ |  |  | bldg:ClosureSurfaceを作る場合は必須とする。 |
| ↑ | lod3MultiSurface |  |  |  | ■ |  | ↑ |
| ↑ | lod4MultiSurface |  |  |  |  | ■ | ↑ |
| bldg:InteriorWallSurface |  |  |  |  |  | ● |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:CeilingSurface |  |  |  |  |  | ● |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:FloorSurface |  |  |  |  |  | ● |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:Door |  |  |  |  | ● | ● |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:Window |  |  |  |  | ● | ● |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |  |
| ↑ | lod4MultiSurface |  |  |  |  | ● |  |
| bldg:BuildingInstallation |  |  |  |  |  |  |  |
| ↑ | lod2Geometry |  |  |  |  |  |  |
| ↑ | lod3Geometry |  |  |  |  |  |  |
| ↑ | lod4Geometry |  |  |  |  |  |  |
| bldg:IntBuildingInstallation |  |  |  |  |  | ■ | LOD4.1及び 4.2では必須とする。 |
| ↑ | lod4Geometry |  |  |  |  | ■ | MultiSufaceを使用することを基本とする。 |
| bldg:BuildingFurniture |  |  |  |  |  | ■ | LOD4.2では必須とする。 |
| ↑ | lod4Geometry |  |  |  |  | ■ | bldg:BuildingFurnitureを作成する場合は必須とする。 MultiSufaceを使用することを基本とする。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← | ← |

補足

LOD4は、BIMモデルからの変換を前提とした区分である。測量を前提としたLOD2やLOD3とは異なり、地物の大きさによる取得の要否ではなく、変換対象となるIFCのクラスによりLODを細分している。LODが上がるにつれ、詳細な地物が含まれるモデルとなる。

