# 4.2.1.5.4 建築物モデル（LOD4.2）の定義

## 4.2.1.5.4 建築物モデル（LOD4.2）の定義

建築物モデル（LOD4.2）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-26**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD4.2 | ● | Building | Solid | 全てを対象とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、底面（GroundSurface）、閉鎖面（ClosureSurface）、扉（Door）及び窓（Window）を境界面とする立体を作成する。 | 測量により取得する場合は立体（Solid）、BIMモデルからの変換により取得する場合は面の集まり（MultiSurface）とする。 |
| LOD4.2 | ■ | BuildingPart | Solid | 一棟の建築物を、属性の異なる複数の部分に分ける場合に必須とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、底面（GroundSurface）、閉鎖面（ClosureSurface）、扉（Door）及び窓（Window）を境界面とする立体を作成する。 | BIMからの変換により取得する場合は使用しない。 |
| LOD4.2 | ● | RoofSurface | MultiSurface | 全てを対象とする。 | 屋根の上方からの正射影の外周を取得し、棟及び谷で区切る。区切った面の各頂点に屋根の高さを与える。 | 屋根の棟及び谷で区切ることにより、屋根の傾斜や向きを再現する。曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | GroundSurface | MultiSurface | 全てを対象とする。 | 屋根の上方からの正射影の外周を取得し、建築物の最下面の高さが異なる箇所で区切る。区切った面の各頂点に建築物の最下面の高さを与える。 |  |
| LOD4.2 | ● | WallSurface | MultiSurface | 全てを対象とする。 | 外壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | 〇 | ClosureSurface | MultiSurface | 境界面となる内壁面や天井面、床面はないが、建築確認申請では部屋となっている空間を区切る場合に必須とする。 | 床面（FloorSurface）、天井面（CeilingSurface）及び内壁面（InteriorWallSurface）を区切る仮想的な境界線に囲まれた面を取得する。 |  |
| LOD4.2 | 〇 | OuterFloorSurface | MultiSurface | ユースケースで必要な場合 | 外壁のうち、上向きとなる面の外周を取得する。面の各頂点に、外壁の高さを与える。 | RoofSurfaceの代替として使用できる。 |
| LOD4.2 | 〇 | OuterCeilingSurface | MultiSurface | ユースケースで必要な場合 | 外壁のうち、下向きとなる面の外周を取得する。面の各頂点に、外壁の高さを与える。 | WallSurfaceの代替として利用できる |
| LOD4.2 | ● | BuildingInstallation | MultiSurface | 全てを対象とする。 | 屋外付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | Door | MultiSurface | 全てを対象とする。 | 扉（Door）の外周を取得する。 |  |
| LOD4.2 | ● | Window | MultiSurface | 全てを対象とする。 | 窓（Window）の外周を取得する。 |  |
| LOD4.2 | ● | Room | Solid | 全てを対象とする。 | 天井面（CeilingSurface）、内壁面（InteriorWallSurface）、閉鎖面（ClosureSurface）及び床面（FloorSurface）を境界面とする立体を作成する。 | 建築確認申請書に部屋として記載されている区画を対象とする。 |
| LOD4.2 | ● | CeilingSurface | MultiSurface | 全てを対象とする。 | 天井の外周を取得する。 |  |
| LOD4.2 | ● | InteriorWallSurface | MultiSurface | 全てを対象とする。 | 部屋（Room）を区切る内壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | FloorSurface | MultiSurface | 全てを対象とする。 | 床の外周を取得する。 |  |
| LOD4.2 | ● | IntBuildingInstallation | MultiSurface | 階段、スロープ、エスカレータ、輸送設備（エレベータ、エスカレータ、動く歩道）、柱、デッキ、ステージ、手すり、パネル、梁 | 屋内付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋内付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | BuildingFurniture | MultiSurface | 全てを対象とする。 | 家具の外形（外側から見える形）を構成する面を取得する。面の各頂点に家具の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | CityObjectGroup | ー | 階 | ー | Roomの集まりとして表現する。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい）注記4.2.1.5.4の注記CityObjectGroupは空間属性をもたないため、空間属性の型及び取得方法を「―」としている。 | ← | ← | ← | ← | ← | ← |

