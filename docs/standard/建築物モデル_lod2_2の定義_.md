# 4.2.1.3.4 建築物モデル（LOD2.2）の定義

## 4.2.1.3.4 建築物モデル（LOD2.2）の定義

建築物モデル（LOD2.2）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-15**

|  |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD2.2 | ● | Building | Solid | 射影の短辺の実長1m以上 | 屋根面（RoofSurface）、外壁面（WallSurface）、屋外床面（OuterFloorSurface）及び底面（GroundSurface）を境界面とする立体を作成する。 | 屋外床面（OuterFloorSurface）を使用する場合は、これも境界面となる。 |
| LOD2.2 | ● | RoofSurface | MultiSurface | 射影の短辺の実長1m以上又は上方からの正射影の面積1m2以上 | 屋根の上方からの正射影の外周を取得し、棟及び谷で区切る。区切った面の各頂点に屋根の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD2.2 | ● | GroundSurface | MultiSurface | 全て対象 | 建築物の上方からの正射影の外周を取得し、外周を構成する各頂点の水平座標に、地表面の高さを与える。 | 地表面の高さは、建築物の上方からの正射影の外周に含まれる地表面の高さのうち、最も低い高さとする。 |
| LOD2.2 | ● | WallSurface | MultiSurface | 全て対象 | 屋根面（RoofSurface）と底面（GroundSurface）を垂直に結ぶ各辺をつないだ面を取得する。方位が変化する場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD2.2 | ■ | BuildingPart | Solid | 一棟の建築物を、主題属性の異なる複数の部分に分ける場合に必須とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 | BuildingPartを使用する場合、一棟のBuildingには必ず2つ以上のBuildingPartが含まれていなければならず、それらは互いに接していなければならない。また、Buildingの空間属性は空でなければならない。 |
| LOD2.2 | ■ | ClosureSurface | MultiSurface | BuildingPartを作成する場合に必須とする。 | BuildingPartと連続する他のBuildingPartとの境界線により囲まれた面を取得する。 | ClosureSurfaceの境界線は、屋根面（RoofSurface）、外壁面（WallSurface）又は底面（GroundSurface）を区切る線分となる。 |
| LOD2.2 | 〇 | OuterFloorSurface | MultiSurface | ユースケースで必要な場合 | 屋外床面（OuterFloorSurface）の外周を取得し、外周の各頂点にその位置の屋根の高さを与える。 | RoofSurfaceの代替として使用できる。 |
| LOD2.2 |  | OuterCeilingSurface | MultiSurface |  |  | 対象外 |
| LOD2.2 | ● | BuildingInstallation | MultiSurface | 短辺の実長1m以上 | 屋外付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

