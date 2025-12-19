# 4.12.1.4.2 トンネルモデル（LOD3）の定義

## 4.12.1.4.2 トンネルモデル（LOD3）の定義

トンネルモデル（LOD3）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-555**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD3 | ● | Tunnel | Solid又はMultiSurface |  | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）、及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 |  |
| LOD3 | ■ | RoofSurface | MultiSurface | トンネルの側面が垂直の場合に、その垂直面の上方に存在する面 | トンネルの外形を取得し、上向きの面を屋根面（RoofSurface）とする。面を構成する各頂点に、屋根の高さを与える。 | トンネル上部が湾曲しており、トンネルの屋根と外壁との区分が難しい場合は、外壁面（WallSurface）として取得する。 |
| LOD3 | ■ | GroundSurface | MultiSurface | トンネルの側面が垂直の場合に、その垂直面の下方に存在する面 | トンネルの外形を取得し、下向きの面を底面（GroundSurface）とする。面を構成する各頂点に、トンネル下面の高さを与える。 | トンネル下部が湾曲しており、トンネルの底と外壁との区分が難しい場合は、外壁面（WallSurface）として取得する。 |
| LOD3 | ● | WallSurface | MultiSurface |  | トンネルの外形を取得し、屋根面（RoofSurface）及び底面（GroundSurface）を除く面を外壁面（WallSurface）とする。面を構成する各頂点にそれぞれの高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD3 | ● | ClosureSurface | MultiSurface |  | 坑口の開口部の外周を取得する。 |  |
| LOD3 |  | OuterCeilingSurface |  |  |  | 対象外。 |
| LOD3 |  | OuterFloorSurface |  |  |  | 対象外。 |
| LOD3 | ■ | TunnelPart | Solid | 多連形のトンネルを一つのTunnelとして表現する場合に必須とする。一つのトンネルを、属性の異なる複数の部分に分ける場合に必須とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 |  |
| LOD3 | ● | TunnelInstallation | MultiSurface |  | 外部付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD3 | ● | Door | MultiSurface |  | 扉（Door）の外周を取得する。 |  |
| LOD3 | ● | Window | MultiSurface |  | 窓（Window）の外周を取得する。 |  |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

