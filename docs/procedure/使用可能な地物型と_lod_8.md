# M.2.2 使用可能な地物型とLOD

## M.2.2 使用可能な地物型とLOD

トンネルモデルは、LODごとに、使用すべき地物型やその空間属性が異なる。

**表 M-2 — M-1: トンネル記述のLOD対応条件**

| ID | `/req/tun/1` |
| --- | --- |
| 主題 | 妥当なトンネルデータセットの要件 |
| 分類 | 要件分類M-1: 妥当なトンネルオブジェクト |
| 条文 | トンネルの記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

トンネルは、CityGMLのTunnelモジュールに定義されたtun:Tunnelを用いて記述する。

- トンネルモデル（LOD0）では、tun:Tunnelのみを使用し、トンネルの坑口の形状を点、線又は面として表現する。CityGMLにはトンネルのLOD0の空間属性が定義されていないため、i-URにおいて拡張した属性uro:lod0Geometryを使用する。この空間属性は数値地形図との互換性を保つ情報を記述するためのデータ型uro:DmAttributeの一部である。
- トンネルモデル（LOD1）では、tun:Tunnelを使用し、トンネルの形状を立体として表現する。立体は、トンネルの上方からの外周の正射影に一律の高さを与えた箱モデルである。このとき、外周には坑口を含む。
- トンネルモデル（LOD2）では、トンネルの外形を立体として表現し、立体の各境界面を、屋根面（tun:RoofSurface）、底面（tun:GroundSurface）及び外壁面（tun:WallSurface）に区分する。このとき、トンネルの外形には坑口を含む。
- トンネルモデル（LOD3）では、トンネルの形状を、立体により表現し、立体を構成する各境界面を地物として取得するとともに、トンネルの開口部やトンネルの外側に付いている付属物（tun:TunnelInstallation）を地物として取得する。開口部に扉や窓が設けられている場合は、それぞれtun:Doorやtun:Windowにより表現するが、これらが無い場合には閉鎖面（tun:ClosureSurface）として記述する。
- トンネルモデル（LOD4）では、トンネルモデル（LOD3）による外形に加えて、トンネルの内部空間（tun:HollowSpace）を表現する。トンネルの内部空間は、建築物の屋内空間と同様、内壁面（tun:InteriorWallSurface）や床面（tun:FloorSurface）などの境界面に分ける。また、トンネル内部の付属物（tun:IntTunnelInstallation）や可動設備（tun:TunnelFurniture）を表現できる。

また、トンネルモデル（LOD1）からトンネルモデル（LOD4）では、一つのトンネルを複数の区間（tun:TunnelPart）に分けて記述することができる。主題属性が異なる場合やトンネルモデル（LOD1）をより実際に近い形状で表現する場合に使用できる。

トンネルの各LODにおいて使用可能な地物型とその空間属性の一覧を表M-3に示す。

**表 M-3 — トンネルモデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | LOD4 | 適用 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| tun:Tunnel |  | ● | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  |  | 数値地形図の取得方法に従う。 |
| ↑ | tun:lod1Solid |  | ● |  |  |  |  |
| ↑ | tun:lod2Solid |  |  | ● |  |  |  |
| ↑ | tun:lod3Solid |  |  |  | ● |  |  |
| ↑ | tun:lod4Solid |  |  |  |  | ■ |  |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ■ | Solidを原則とする。BIMから作成する場合はMultiSurfaceとする。 |
| tun:TunnelPart |  |  | ■ | ■ | ■ | ■ | LOD1において、より実際の形状に近い表現を行う場合に必須とする。LOD2以上において、多連形のトンネルを一つのTunnelとして表現する場合に必須とする。LOD2以上において、一つのトンネルを、属性の異なる複数の部分に分ける場合に必須とする。 |
| ↑ | tun:lod1Solid |  | ■ |  |  |  |  |
| ↑ | tun:lod2Solid |  |  | ■ |  |  | tun:TunnelPartを使用する場合は必須とする。 |
| ↑ | tun:lod3Solid |  |  |  | ■ |  | ↑ |
| ↑ | tun:lod4Solid |  |  |  |  | ■ | ↑ |
| tun:TunnelInstallation |  |  |  |  | ● | ● |  |
| ↑ | tun:lod2Geometry |  |  |  |  |  |  |
| ↑ | tun:lod3Geometry |  |  |  | ● |  | MultiSurfaceとする。 |
| ↑ | tun:lod4Geometry |  |  |  |  | ● | MultiSurfaceとする。 |
| tun:RoofSurface |  |  |  | ■ | ■ | ■ | トンネルの外形を構成する面のうち、上向きの面に使用する。 |
| ↑ | tun:lod2MultiSurface |  |  | ■ |  |  | tun:RoofSurfaceを作る場合は必須とする。 |
| ↑ | tun:lod3MultiSurface |  |  |  | ■ |  | ↑ |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ■ | ↑ |
| tun:WallSurface |  |  |  | ● | ● | ● | トンネルの外形を構成する面のうち、側方の面に使用する。 |
| ↑ | tun:lod2MultiSurface |  |  | ● |  |  |  |
| ↑ | tun:lod3MultiSurface |  |  |  | ● |  | ↑ |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ● | ↑ |
| tun:GroundSurface |  |  |  | ■ | ■ | ■ | トンネルの外形を構成する面のうち、下向きの面に使用する。 |
| ↑ | tun:lod2MultiSurface |  |  | ■ |  |  | tun:GroundSurfaceを作る場合は必須とする。 |
| ↑ | tun:lod3MultiSurface |  |  |  | ■ |  | ↑ |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ■ | ↑ |
| tun:OuterFloorSurface |  |  |  |  |  |  | 標準製品仕様書では使用しない。 |
| ↑ | tun:lod2MultiSurface |  |  |  |  |  |  |
| ↑ | tun:lod3MultiSurface |  |  |  |  |  | ↑ |
| ↑ | tun:lod4MultiSurface |  |  |  |  |  | ↑ |
| tun:OuterCeilingSurface |  |  |  |  |  |  | 標準製品仕様書では使用しない。 |
| ↑ | tun:lod2MultiSurface |  |  |  |  |  |  |
| ↑ | tun:lod3MultiSurface |  |  |  |  |  | ↑ |
| ↑ | tun:lod4MultiSurface |  |  |  |  |  | ↑ |
| tun:ClosureSurface |  |  |  |  | ● | ● | 出入口をtun:ClosureSurfaceとして表現する。 |
| ↑ | tun:lod2MultiSurface |  |  |  |  |  |  |
| ↑ | tun:lod3MultiSurface |  |  |  | ● |  | ↑ |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ● | ↑ |
| tun:InteriorWallSurface |  |  |  |  |  | ● |  |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ● |  |
| tun:CeilingSurface |  |  |  |  |  | ■ |  |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ● |  |
| tun:FloorSurface |  |  |  |  |  | ■ |  |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ● |  |
| tun:Door |  |  |  |  | ● | ● |  |
| ↑ | tun:lod3MultiSurface |  |  |  | ● |  |  |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ● |  |
| tun:Window |  |  |  |  | ● | ● |  |
| ↑ | tun:lod3MultiSurface |  |  |  | ● |  |  |
| ↑ | tun:lod4MultiSurface |  |  |  |  | ● |  |
| tun:HollowSpace |  |  |  |  |  | ● |  |
| ↑ | tun:lod4Solid |  |  |  |  | ● |  |
| tun:IntTunnelnstallation |  |  |  |  |  | ■ | LOD4.1及びLOD4.2では必須とする。 |
| ↑ | tun:lod4Geometry |  |  |  |  | ● | MultiSurfaceとする。 |
| tun:TunnelFurniture |  |  |  |  |  | ■ | LOD4.2では必須とする。 |
| ↑ | tun:lod4Geometry |  |  |  |  | ● | MultiSurfaceとする。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← | ← |

