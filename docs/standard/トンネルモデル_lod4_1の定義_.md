# 4.12.1.5.3 トンネルモデル（LOD4.1）の定義

## 4.12.1.5.3 トンネルモデル（LOD4.1）の定義

トンネルモデル（LOD4.1）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-559**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD4.1 | ● | Tunnel | Solid又はMultiSurface | 全てを対象とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び閉鎖面（ClosureSurface）の集まり、又は、これらを境界面とする立体を作成する。 | 測量により取得する場合は、Solidとする。BIMモデルからの変換により取得する場合はMultiSurfaceとする。 |
| LOD4.1 | ■ | RoofSurface | MultiSurface | トンネルの側面が垂直の場合に、その垂直面の上方に存在する面 | トンネルの外形を取得し、上向きの面を屋根面（RoofSurface）とする。面を構成する各頂点に、屋根の高さを与える。 | トンネル上部が湾曲しており、トンネルの屋根と外壁との区分が難しい場合は、外壁面（WallSurface）として取得する。 |
| LOD4.1 | ■ | GroundSurface | MultiSurface | トンネルの側面が垂直の場合に、その垂直面の下方に存在する面 | トンネルの外形を取得し、下向きの面を底面（GroundSurface）とする。面を構成する各頂点に、トンネル下面の高さを与える。 | トンネル下部が湾曲しており、トンネルの底と外壁との区分が難しい場合は、外壁面（WallSurface）として取得する。 |
| LOD4.1 | ● | WallSurface | MultiSurface |  | トンネルの外形を取得し、屋根面（RoofSurface）及び底面（GroundSurface）以外の面を外壁面（WallSurface）とする。面を構成する各頂点に、それぞれの高さを与える。 |  |
| LOD4.1 | ● | ClosureSurface | MultiSurface | 坑口の開口部 | トンネルの坑口の開口部の外周に囲まれた範囲を面として取得する。 |  |
| LOD4.1 | ■ | TunnelPart | Solid又はMultiSurface | 多連形のトンネルを一つのTunnelとして表現する場合 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 |  |
| LOD4.1 |  | OuterFloorSurface |  |  |  |  |
| LOD4.1 |  | OuterCeilingSurface |  |  |  |  |
| LOD4.1 | ● | Door | MultiSurface | 全てを対象とする。 | 扉（Door）の外周を取得する。 |  |
| LOD4.1 | ● | Window | MultiSurface | 全てを対象とする。 | 窓（Window）の外周を取得する。 |  |
| LOD4.1 | ● | TunnelInstallation | MultiSurface | 全てを対象とする。 | 屋外付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.1 | ● | HollowSpace | Solid | 全てを対象とする。 | 天井面（CeilingSurface）、内壁面（InteriorWallSurface）、閉鎖面（ClosureSurface）及び床面（FloorSurface）を境界面とする立体を作成する。 |  |
| LOD4.1 | ■ | CeilingSurface | MultiSurface | トンネル内部の側面が垂直の場合に、その垂直面の上方に存在する面 | トンネル内部に存在する下向きの面の外周を取得する。 |  |
| LOD4.1 | ● | InteriorWallSurface | MultiSurface |  | トンネル内部の壁面のうち、天井面（CeilingSurface）又は床面（FloorSuface）として取得する面を除いた面を取得する。 |  |
| LOD4.1 | ■ | FloorSurface | MultiSurface | トンネル内部の側面が垂直の場合に、その垂直面の下方に存在する面 | トンネル内部に存在する上向きの面の外周を取得する。 | 水路トンネルの場合は床面（FloorSurface）ではなく、内壁面（InteriorWallSurface）として取得する。 |
| LOD4.1 | ● | IntTunnelInstallation | MultiSurface | 短辺が実長3m以上又は短辺が実長1m以上かつ横断又は縦断の面積が3m2以上 | 屋内付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋内付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4.1 |  | TunnelFurniture |  |  |  |  |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

