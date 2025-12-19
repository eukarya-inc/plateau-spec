# M.8 tun:_BoundarySurface

## M.8 tun:_BoundarySurface

tun:_BoundarySurfaceは、トンネル（tun:Tunnel）やトンネルの部分（tun:TunnelPart）の屋根面及び壁面、底面等の境界面を構成する地物型である。トンネルモデル（LOD2）、トンネルモデル（LOD3）及びトンネルモデル（LOD4）を作成する際に使用される。

この地物型は抽象地物であり、実装においては、この下位型の具象となる地物型を使用する。

**表 M-21 — M-15: tun:_BoundarySurfaceの種類**

| ID | `/req/tun/15` |
| --- | --- |
| 主題 | 妥当なトンネルデータセットの要件 |
| 分類 | 要件分類M-1: 妥当なトンネルオブジェクト |
| 条文 | tun:_BoundarySurfaceは、屋根面（tun:RoofSurface）、外壁面（tun:WallSurface）、底面（tun:GroundSurface）、外部床面（tun:OuterFloorSurface）、内壁面（tun:InteriorWallSurface）、天井面（tun:CeilingSurface）、床面（tun:FloorSurface）、閉鎖面（tun:ClosureSurface）のいずれかとなる。 |

CityGMLでは、トンネルの境界面としてtun:_BoundarySurfaceの下位型となる地物型が用意されている。ただし外部天井面（tun:OuterCeillingSurface）、屋外床面（tun:OuterFloorSurface）は対象外とする（使用しない）。

![Image](images/tocM_08_img001.png)

## M.8.1 境界面の空間属性

屋根面及び壁面、底面の境界面となる地物型の空間属性（面）は、トンネルの立体を構成する境界面に一致する。

**表 M-22 — M-16: 境界面の幾何オブジェクト参照条件**

| ID | `/req/tun/16` |
| --- | --- |
| 主題 | 妥当なトンネルデータセットの要件 |
| 分類 | 要件分類M-1: 妥当なトンネルオブジェクト |
| 条文 | 境界面の空間属性を構成する幾何オブジェクトは、トンネルの立体を構成する境界面として参照されなければならない。 |

## M.8.2 境界面の下位型

## M.8.2.1 tun:RoofSurface

tun:RoofSurfaceは、屋根面を記述するために使用する地物型である。

**表 M-23 — M-17: 外壁面と屋根面の区分条件**

| ID | `/req/tun/17` |
| --- | --- |
| 主題 | 妥当なトンネルデータセットの要件 |
| 分類 | 要件分類M-1: 妥当なトンネルオブジェクト |
| 条文 | 外壁面（tun:WallSurface）と屋根面（tun:RoofSurface）との区分が構造上難しい場合は、外壁面（tun:WallSurface）として取得することを基本とする。 |

## M.8.2.2 tun:WallSurface

tun:WallSurfaceは、壁面を記述するために使用する地物型である。

## M.8.2.3 tun:GroundSurface

tun:GroundSurfaceは、トンネルの底面を記述するために使用する地物型である。

## M.8.2.4 tun:OuterFloorSurface

tun:OuterFloorSurfaceはトンネルの外側を覆う部分であり、通行可能な床面としての機能を有する部分の記述に使用する地物型である。標準製品仕様書では使用しない。

## M.8.2.5 tun:ClosureSurface

tun:ClosureSurfaceは、トンネルの開口部を立体として閉じるために、境界面として設けられた仮想的な面である。トンネルの境界面に開口部が存在するが、開口部内の詳細なデータ作成が不要である場合に、開口部を閉じるために便宜上設けられている。内部空間で用途が異なる等により区分する場合は、tun:HollowSpaceを使用して空間を区分する。

![Image](images/tocM_08_img002.png)

## M.8.2.6 tun:OuterCeilingSurface

tun:OuterCeilingSurfaceは、トンネルの外側を覆う部分であり、天井としての機能を有する面を記述するために使用する地物型である。標準製品仕様書では使用しない。

## M.8.2.7 tun:CeilingSurface

tun:CeilingSurfaceは、トンネルの内部の天井面を記述するために使用する地物型である。

**表 M-24 — M-18: 内壁面と天井面の区分条件**

| ID | `/req/tun/18` |
| --- | --- |
| 主題 | 妥当なトンネルデータセットの要件 |
| 分類 | 要件分類M-1: 妥当なトンネルオブジェクト |
| 条文 | 内壁面（tun:InteriorWallSurface）と天井面（tun:CeilingSurface）との区分が構造上難しい場合は、内壁面（tun:InteriorWallSurface）として取得することを基本とする。 |

## M.8.2.8 tun:InteriorWallSurface

tun:InteriorWallSurfaceは、トンネルの内部空間の区画を区切る壁や仕切りとしての機能を有する面を記述するために使用する地物型である。

## M.8.2.9 tun:FloorSurface

tun:FloorSurfaceは、トンネルの内部空間の下面に位置する舗装等が存在する板状の構造物（床面）を記述するために使用する地物型である。

