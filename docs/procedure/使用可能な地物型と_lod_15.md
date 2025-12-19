# T.2.2 使用可能な地物型とLOD

## T.2.2 使用可能な地物型とLOD

水部の記述には、CityGMLのWaterBodyモジュールに定義されたwtr:WaterBodyを用いて記述する。wtr:WaterBodyは、詳細度（LOD）ごとに、使用すべき地物型やその空間属性が異なる。

**表 T-2 — T-1: 水部記述のLOD対応条件**

| ID | `/req/wtr/1` |
| --- | --- |
| 主題 | 妥当な水部データセットの要件 |
| 分類 | 要件分類T-1: 妥当な水部オブジェクト |
| 条文 | 水部の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

WaterBodyモジュールには、水部（wtr:WaterBody）と、水部の境界面（wtr:_BoundarySurface）が定義される。水部の境界面は、水面（wtr:WaterSurface）、水底面（wtr:WaterGroundSurface）及び閉鎖面（wtr:WaterClosureSurface）に区分される。

LOD0では、Wtr:Waterbodyのみを使用し、水部の形状を線として表現する。LOD1では、wtr:WaterBodyのみを使用し、水部の形状を面として表現する。LOD2及びLOD3では、wtr:WaterBodyを使用して水部の形状を立体として表現し、その立体の境界面をwtr:WaterSurface、wtr:WaterGroundSurface又はwtr:WaterClosureSurfaceに区分する。

水部の各LODにおいて使用可能な地物型とその空間属性を表T-3に示す。

**表 T-3 — 水部モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| wtr:WaterBody |  | ● | ● | ● | ● |  |
| ↑ | lod0MultiCurve | ■ |  |  |  | 河川中心線（wtr:lod0MultiCurve）を基本とする。縁線、界線又は点として取得する場合はuro:lod0Geometryを使用する。 |
| ↑ | lod0Geometry | ■ |  |  |  | ↑ |
| ↑ | lod1MultiSurface |  | ● |  |  |  |
| ↑ | lod2Solid |  |  | ● |  |  |
| ↑ | lod3Solid |  |  |  | ● |  |
| wtr:WaterSurface |  |  |  | ● | ● | 水部の形状を構成する面のうち、水面に使用する。 |
| ↑ | lod2Surface |  |  | ● |  |  |
| ↑ | lod3Surface |  |  |  | ● |  |
| wtr:WaterGroundSurface |  |  |  | ● | ● | 水底の境界面に使用する。 |
| ↑ | lod2Surface |  |  | ● |  |  |
| ↑ | lod3Surface |  |  |  | ● |  |
| wtr:WaterClosureSurface |  |  |  | ■ | ■ | 仮想的な面を作成したい場合に使用する。 |
| ↑ | lod2Surface |  |  | ■ |  | WaterClosureSurfaceを作る場合は必須とする。 |
| ↑ | lod3Surface |  |  |  | ■ | ↑ |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

