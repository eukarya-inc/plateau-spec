# 4.2.1.4.5 建築物モデル（LOD3.3）の定義

## 4.2.1.4.5 建築物モデル（LOD3.3）の定義

建築物モデル（LOD3.3）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-21**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD3.3 | ● | Building | Solid | 短辺が実長1m以上 | 屋根面（RoofSurface）、壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、底面（GroundSurface）、閉鎖面（ClosureSurface）、扉（Door）及び窓（Window）を境界面とする立体を作成する。 | 取得する建築物は、原則としてLOD0及びLOD1と同じである。 |
| LOD3.3 | ● | RoofSurface | MultiSurface | 全てを対象とする。 | 屋根の上方からの正射影の外周を取得し、棟（屋根の頂部であり、屋根の分水嶺となる箇所）及び谷（屋根と屋根のつなぎの谷状の部分）で区切る。区切った面の各頂点に屋根の高さを与える。 | 屋根の棟及び谷で区切ることにより、屋根の傾斜や向きを再現する。曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD3.3 | ● | GroundSurface | MultiSurface | 全てを対象とする。 | 地表面と外壁面との交線を取得し、各頂点に地表面の高さを与える。 | 地表面の高さは、上方からの正射影の外周に含まれる地表面の頂点の標高のうち、最も低い標高とする。 |
| LOD3.3 | ● | WallSurface | MultiSurface | 外壁 | 外壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| ↑ | ↑ | ↑ | ↑ | 軒裏 | 屋根の上方からの正射影の外周と、地表面と外壁面との交線により囲まれた面を取得する。高さは、各頂点の高さとする。 |  |
| LOD3.3 | ■ | BuildingPart | Solid | 一棟の建築物を、属性の異なる複数の部分に分ける場合に使用する。 | 屋根面（RoofSurface）、壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、底面（GroundSurface）、閉鎖面（ClosureSurface）、扉（Door）及び窓（Window）を境界面とする立体を作成する。 |  |
| LOD3.3 | ■ | ClosureSurface | MultiSurface | BuildingPartを使用する場合に必須とする。 | BuildingPartと連続するBuildingPartとの境界線により囲まれた面を取得する。 |  |
| LOD3.3 | 〇 | OuterFloorSurface | MultiSurface | ユースケースで必要な場合 | 外壁のうち、上向きとなる面の外周を取得する。面の各頂点に、外壁の高さを与える。 | RoofSurfaceの代替として使用できる。 |
| LOD3.3 | 〇 | OuterCeilingSurface | MultiSurface | ユースケースで必要な場合 | 外壁のうち、下向きとなる面の外周を取得する。面の各頂点に、外壁の高さを与える。 | WallSurfaceの代替として利用できる。 |
| LOD3.3 | ● | BuildingInstallation | MultiSurface | 全てを対象とする。 | 屋外付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD3.3 | ● | Door | MultiSurface | 全てを対象とする。 | 扉（Door）の外周を取得する。 |  |
| LOD3.3 | ● | Window | Window | 全てを対象とする。 | 窓（Window）の外周を取得する。 |  |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

