# 4.16.1.3.2 地下街モデル（LOD2）の定義

## 4.16.1.3.2 地下街モデル（LOD2）の定義

地下街モデル（LOD2）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-734**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD2 | ● | UndergroundBuilding | Solid |  | 屋根面（RoofSurface）、外壁面（WallSurface）及び底面（GroundSurface）を境界面とする立体を作成する。 |  |
| LOD2 | ● | RoofSurface | MultiSurface | 射影の短辺の実長3m以上 | 地下街の外形を取得し、上向き面を屋根面（RoofSurface）とする。面を構成する各頂点にそれぞれの高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD2 | ● | GroundSurface | MultiSurface | 全てを対象とする。 | 地下街の外形を取得し、下向き面を底面（GroundSurface）とする。面を構成する各頂点にそれぞれの高さを与える。 |  |
| LOD2 | ● | WallSurface | MultiSurface | 全てを対象とする。 | 地下街の外形を取得し、屋根面（RoofSuface）及び底面（GroundSurface）以外の面を外壁面（WallSurface）とする。面を構成する各頂点にそれぞれの高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD2 | ■ | BuildingPart | Solid | 1棟の地下街を、主題属性の異なる複数の部分に分けたい場合 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 | BuildingPartを使用する場合、1棟のBuildingには必ず2つ以上のBuildingPartが含まれていなければならず、それらは互いに接していなければならない。BuildingPartを使用する場合、Buildingの空間属性は空となる。 |
| LOD2 | ■ | ClosureSurface | MultiSurface | BuildingPartを作成する場合 | BuildingPartと連続する他のBuildingPartとの境界線により囲まれた面を取得する。 | ClosureSurfaceの境界線は、屋根面（RoofSurface）、外壁面（WallSurface）又は底面（GroundSurface）を区切る線分となる。 |
| LOD2 |  | OuterFloorSurface |  |  |  | 対象外 |
| LOD2 |  | OuterCeilingSurface |  |  |  | 対象外 |
| LOD2 |  | BuildingInstallation |  |  |  | 対象外 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

