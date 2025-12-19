# 4.2.1.5.2 建築物モデル（LOD4.0）の定義

## 4.2.1.5.2 建築物モデル（LOD4.0）の定義

建築物モデル（LOD4.0）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-24**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD4.0 | ● | Building | Solid又はMultiSurface | 全てを対象とする。 | 屋根面（RoofSurface）、壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、底面（GroundSurface）、閉鎖面（ClosureSurface）、扉（Door）及び窓（Window）を境界面とする立体を作成する。 | 測量により取得する場合は、Solidとする。BIMモデルからの変換により取得する場合はMultiSurfaceとする。 |
| LOD4.0 | ■ | BuildingPart | Solid | 一棟の建築物を、属性の異なる複数の部分に分ける場合に必須とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、底面（GroundSurface）、閉鎖面（ClosureSurface）、扉（Door）及び窓（Window）を境界面とする立体を作成する。 | BIMモデルからの変換により取得する場合は使用しない。 |
| LOD4.0 | ● | RoofSurface | MultiSurface | 全てを対象とする。 | 屋根の上方からの正射影の外周を取得し、棟及び谷で区切る。区切った面の各頂点に屋根の高さを与える。 | 屋根の棟及び谷で区切ることにより、屋根の傾斜や向きを再現する。屋根の棟及び谷は、以下を指す。![Image](images/建築物モデル_lod4_0の定義__img001.png)曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.0 | ● | GroundSurface | MultiSurface | 全てを対象とする。 | 建築物の最下面の外周を取得する。 |  |
| LOD4.0 | ● | WallSurface | MultiSurface | 全てを対象とする。 | 外壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.0 | ■ | ClosureSurface | MultiSurface | 境界面となる内壁面や天井面、床面はないが、建築確認申請では部屋となっている空間を区切る場合に必須とする。 | 床面（FloorSurface）、天井面（CeilingSurface）及び内壁面（InteriorWallSurface）を区切る仮想的な境界線に囲まれた面を取得する。 |  |
| LOD4.0 | 〇 | OuterFloorSurface | MultiSurface | ユースケースで必要な場合 | 外壁のうち、上向きとなる面の外周を取得する。面の各頂点に、外壁の高さを与える。 | RoofSurfaceの代替として使用できる。 |
| LOD4.0 | 〇 | OuterCeilingSurface | MultiSurface | ユースケースで必要な場合 | 外壁のうち、下向きとなる面の外周を取得する。面の各頂点に、外壁の高さを与える。 | WallSurfaceの代替として利用できる。 |
| LOD4.0 | ● | Door | MultiSurface | 全てを対象とする。 | 扉（Door）の外周を取得する。 |  |
| LOD4.0 | ● | Window | MultiSurface | 全てを対象とする。 | 窓（Window）の外周を取得する。 |  |
| LOD4.0 | ● | BuildingInstallation | MultiSurface | 全てを対象とする。 | 屋外付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.0 | ● | Room | Solid | 全てを対象とする。 | 天井面（CeilingSurface）、内壁面（InteriorWallSurface）、閉鎖面（ClosureSurface）及び床面（FloorSurface）を境界面とする立体を作成する。 | 建築確認申請書に部屋として記載されている区画を対象とする。 |
| LOD4.0 | ● | CeilingSurface | MultiSurface | 全てを対象とする。 | 天井の外周を取得する。 |  |
| LOD4.0 | ● | InteriorWallSurface | MultiSurface | 全てを対象とする。 | 部屋（Room）を区切る内壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.0 | ● | FloorSurface | MultiSurface | 全てを対象とする。 | 床の外周を取得する。 |  |
| LOD4.0 |  | IntBuildingInstallation |  |  |  | 対象外 |
| LOD4.0 | ● | CeilingSurface | MultiSurface | 全てを対象とする。 | 天井の外周を取得する。 |  |
| LOD4.0 | ● | InteriorWallSurface | MultiSurface | 全てを対象とする。 | 部屋（Room）を区切る内壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.0 | ● | FloorSurface | MultiSurface | 全てを対象とする。 | 床の外周を取得する。 |  |
| LOD4.0 |  | BuildingFurniture |  |  |  | 対象外 |
| LOD4.0 | ● | CityObjectGroup | ー | 階 | ー | Roomの集まりとして表現する。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい）注記4.2.1.5.2の注記CityObjectGroupは空間属性をもたないため、「―」としている。 | ← | ← | ← | ← | ← | ← |

