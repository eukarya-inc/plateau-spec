# 4.2.1.3.2 建築物モデル（LOD2.0）の定義

## 4.2.1.3.2 建築物モデル（LOD2.0）の定義

建築物モデル（LOD2.0）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-13**

|  |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD2.0 | ● | Building | Solid | 射影の短辺の実長1m 以上 | 屋根面（RoofSurface）、外壁面（WallSurface）及び底面（GroundSurface）を境界面とする立体を作成する。 |  |
| LOD2.0 | ● | RoofSurface | MultiSurface | 射影の短辺の実長3m以上 | 屋根の上方からの正射影の外周を取得し、棟（屋根の頂部であり、屋根の分水嶺となる箇所）及び谷（屋根と屋根のつなぎの谷状の部分）で区切る。区切った面の各頂点に屋根の高さを与える。 | 屋根の棟及び谷で区切ることにより、屋根の傾斜や向きを再現する。屋根の棟及び谷は、以下を指す。![Image](images/建築物モデル_lod2_0の定義__img001.png)曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD2.0 | ● | GroundSurface | MultiSurface | 全て対象 | 建築物の上方からの正射影の外周を取得する。外周を構成する各頂点に、地表面の高さを与える。 | 地表面の高さは、建築物の上方からの正射影の外周に含まれる地表面の高さのうち、最も低い高さとする。 |
| LOD2.0 | ● | WallSurface | MultiSurface | 全て対象 | 屋根面（RoofSurface）と底面（GroundSurface）を垂直に結ぶ各辺をつないだ面を取得する。方位が変化する場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD2.0 | ■ | BuildingPart | Solid | 一棟の建築物を、主題属性の異なる複数の部分に分ける場合に必須とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 | BuildingPartを使用する場合、一棟のBuildingには必ず2つ以上のBuildingPartが含まれていなければならず、それらは互いに接していなければならない。BuildingPartを使用する場合、Buildingの空間属性は空となる。 |
| LOD2.0 | ■ | ClosureSurface | MultiSurface | BuildingPartを作成する場合に必須とする。 | BuildingPartと連続する他のBuildingPartとの境界線により囲まれた面を取得する。 | ClosureSurfaceの境界線は、屋根面（RoofSurface）、外壁面（WallSurface）又は底面（GroundSurface）を区切る線分となる。 |
| LOD2.0 |  | OuterFloorSurface |  |  |  | 対象外 |
| LOD2.0 |  | OuterCeilingSurface |  |  |  | 対象外 |
| LOD2.0 |  | BuildingInstallation |  |  |  | 対象外 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

