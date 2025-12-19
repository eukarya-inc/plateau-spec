# 4.16.1.5.4 地下街モデル（LOD4.2）の定義

## 4.16.1.5.4 地下街モデル（LOD4.2）の定義

地下街モデル（LOD4.2）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-741**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD4.2 | ● | UndergroundBuilding | Solid又はMultiSurface | 全てを対象とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）及び底面（GroundSurface）を境界面とする立体又は面の集まりを作成する。 | 測量により取得する場合は、Solidとする。BIMモデルからの変換により取得する場合はMultiSurfaceとする。 |
| LOD4.2 | ● | RoofSurface | MultiSurface | 全てを対象とする。 | 地下街の外形を取得し、上向き面を屋根面（RoofSurface）とする。面を構成する各頂点にそれぞれの高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | GroundSurface | MultiSurface | 全てを対象とする。 | 地下街の外形を取得し、下向き面を底面（GroundSurface）とする。面を構成する各頂点にそれぞれの高さを与える。 |  |
| LOD4.2 | ● | WallSurface | MultiSurface | 全てを対象とする。 | 地下街の外形を取得し、屋根面（RoofSuface）及び底面（GroundSurface）以外の面を外壁面（WallSurface）とする。面を構成する各頂点にそれぞれの高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ■ | BuildingPart | Solid | 1棟の地下街を、主題属性の異なる複数の部分に分けたい場合 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 | BuildingPartを使用する場合、1棟のBuildingには必ず2つ以上のBuildingPartが含まれていなければならず、それらは互いに接していなければならない。BuildingPartを使用する場合、Buildingの空間属性は空となる。 |
| LOD4.2 | ● | ClosureSurface | MultiSurface | 地下街への出入り口となる開口部 | 地下街の外壁面と地表面との交線により囲まれた面を取得する。 | 地上に設置された地下街出入口の建屋は都市設備（CityFurniture）として取得する。 |
| ↑ | ■ | ClosureSurface | MultiSurface | BuildingPartを作成する場合 | BuildingPartと連続する他のBuildingPartとの境界線により囲まれた面を取得する。 | ClosureSurfaceの境界線は、屋根面（RoofSurface）、外壁面（WallSurface）又は底面（GroundSurface）を区切る線分となる。 |
| LOD4.2 |  | OuterFloorSurface |  |  |  | 対象外 |
| LOD4.2 |  | OuterCeilingSurface |  |  |  | 対象外 |
| LOD4.2 | ● | BuildingInstallation | MultiSurface | 全てを対象とする。 | 屋外付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | Door | MultiSurface | 全てを対象とする。 | 扉（Door）の外周を取得する。 |  |
| LOD4.2 | ● | Window | MultiSurface | 全てを対象とする。 | 窓（Window）の外周を取得する。 |  |
| LOD4.2 | ● | Room | Solid | 全てを対象とする。 | 天井面（CeilingSurface）、内壁面（InteriorWallSurface）、閉鎖面（ClosureSurface）及び床面（FloorSurface）を境界面とする立体を作成する。 |  |
| LOD4.2 | ● | CeilingSurface | MultiSurface | 全てを対象とする。 | 天井の外周を取得する。 |  |
| LOD4.2 | ● | InteriorWallSurface | MultiSurface | 全てを対象とする。 | 部屋（Room）を区切る内壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | FloorSurface | MultiSurface | 全てを対象とする。 | 床の外周を取得する。 |  |
| LOD4.2 | ● | IntBuildingInstallation | MultiSurface | 階段、スロープ、エスカレータ、輸送設備（エレベータ、エスカレータ、動く歩道）、柱、デッキ、ステージ、手すり、パネル、梁 | 屋内付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋内付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | BuildingFurniture | MultiSurface | 全てを対象とする。 | 家具の外形（外側から見える形）を構成する面を取得する。面の各頂点に家具の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.2 | ● | CityObjectGroup | ― | 階 | ― | Roomの集まりとして表現する。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい）注記4.16.1.5.4の注記CityObjectGroupは空間属性をもたないため、「―」としている。 | ← | ← | ← | ← | ← | ← |

