# N.2.2 使用可能な地物型とLOD

## N.2.2 使用可能な地物型とLOD

その他の構造物モデルは、LODごとに使用すべき地物型やその空間属性を定めている。

**表 N-2 — N-1: その他の構造物記述のLOD対応条件**

| ID | `/req/cons/1` |
| --- | --- |
| 主題 | 妥当なその他の構造物データセットの要件 |
| 分類 | 要件分類N-1: 妥当なその他の構造物オブジェクト |
| 条文 | その他の構造物の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

その他の構造物は、橋梁（brid:Bridge）やトンネル（tun:Tunnel）と同様に、本体（uro:OtherConstruction）とその付属物（uro:ConstructionInstallation）から構成される。また、本体の外形は、境界面（uro:_BoundarySurface）に区分される。境界面はその向きや機能により、屋根面（uro:RoofSurface）、外壁面（uro:WallSurface）、底面（uro:GroundSurface）、閉鎖面（uro:ClosureSurface）、屋外床面（uro:OuterFloorSurface）及び屋外天井面（uro:OuterRoofSurface）に区分する。

- その他の構造物モデル（LOD0）では、uro:OtherConstructionのみを使用し、その他の構造物の形状を点、線又は面として表現する。
- その他の構造物モデル（LOD1）では、uro:OtherConstructionのみを使用し、その他の構造物の形状を立体として表現する。
- その他の構造物モデル（LOD2）では、uro:OtherConstructionを使用し、その他の構造物の形状を立体として表現するとともに、その境界面を区分する。
- その他の構造物モデル（LOD3）では、uro:OtherConstructionを使用し、その他の構造物の形状を立体として表現するとともに、その境界面を区分し、また、その他の構造物の付属物をuro:ConstructionInstallationにより表現できる。

その他の構造物の各LODにおいて使用可能な地物型と空間属性を、表N-3に示す。

**表 N-3 — その他の構造物モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| uro:OtherConstruction |  | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  | 数値地形図（DM）の取得方法に従う。 |
| ↑ | uro:lod1Geometry |  | ● |  |  | 立体となる。 |
| ↑ | uro:lod2Geometry |  |  | ● |  | 立体となる。 |
| ↑ | uro:lod3Geometry |  |  |  | ● | 立体となる。 |
| uro:RoofSurface |  |  |  | ● | ● | 構造物の外形を構成する面のうち、上向きの面に使用する。 |
| ↑ | uro:lod2MultiSurface |  |  | ● |  |  |
| ↑ | uro:lod3MultiSurface |  |  |  | ● | ↑ |
| uro:WallSurface |  |  |  | ● | ● | 構造物の外形を構成する面のうち、側方の面に使用する。 |
| ↑ | uro:lod2MultiSurface |  |  | ● |  |  |
| ↑ | uro:lod3MultiSurface |  |  |  | ● | ↑ |
| uro:GroundSurface |  |  |  | ● | ● | 構造物の外形を構成する面のうち、下向きの面に使用する。 |
| ↑ | uro:lod2MultiSurface |  |  | ● |  |  |
| ↑ | uro:lod3MultiSurface |  |  |  | ● | ↑ |
| uro:ClosureSurface |  |  |  | ■ | ■ | 行政界で区切る場合に必須とする。 |
| ↑ | uro:lod2MultiSurface |  |  | ■ |  | ClosureSurfaceを作成する場合は必須とする。 |
| ↑ | uro:lod3MultiSurface |  |  |  | ■ | ↑ |
| uro:OuterCeilingSurface |  |  |  |  | 〇 |  |
| ↑ | uro:lod2MultiSurface |  |  |  |  | OuterCeilingSurfaceを作成する場合は必須とする。 |
| ↑ | uro:lod3MultiSurface |  |  |  | ■ | ↑ |
| uro:OuterFloorSurface |  |  |  |  | 〇 |  |
| ↑ | uro:lod2MultiSurface |  |  |  |  | OuterFloorSurfaceを作成する場合は必須とする。 |
| ↑ | uro:lod3MultiSurface |  |  |  | ■ | ↑ |
| uro:ConstructionlInstallation |  |  |  |  | ■ | LOD3.1において必須とする。 |
| ↑ | uro:lod2Geometry |  |  |  |  |  |
| ↑ | uro:lod3Geometry |  |  |  | ■ | OtherConstructionInstallationを作成する場合は必須とする。 MultiSurfaceとする。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

