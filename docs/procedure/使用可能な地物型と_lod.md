# E.2.2 使用可能な地物型とLOD

## E.2.2 使用可能な地物型とLOD

交通（鉄道）モデルは、LODごとに、使用すべき地物型やその空間属性が異なる。

**表 E-2 — E-1: 鉄道記述のLOD対応条件**

| ID | `/req/rwy/1` |
| --- | --- |
| 主題 | 妥当な交通（鉄道）データセットの要件 |
| 分類 | 要件分類E-1: 妥当な交通（鉄道）オブジェクト |
| 条文 | 鉄道の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

tran:Railwayは、列車又は車両を走らせるための通路となる範囲を示し、tran:TrafficAreaとtran:AuxiliaryTrafficAreaの集まりとして構成される。tran:TrafficAreaは、レールのような通行の用に供される地物を記述するための地物型である。一方、tran:AuxiliaryTrafficAreaは、軌道の構造を支える補助的な役割を果たす地物を記述するための地物型である。

tran:Railwayの空間属性は、LOD0では線とし、LOD1以降では面とする。また、tran:TrafficAreaは、面（軌道の範囲）及び線（軌道中心線）となる。ただし、この地物型はLOD2以降でのみ使用可能となる。また、tran:AuxiliaryTrafficAreaは面となり、LOD3で使用可能とする。
交通（鉄道）モデルの各LODにおいて使用可能な地物型と空間属性を、 表E-3に示す。

**表 E-3 — 交通（鉄道）モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| tran:Railway |  | ● | ● | ● | ● |  |
| ↑ | tran:lod0Network | ■ |  |  |  | LOD0はネットワークを原則とするが、数値地形図との互換性を保つために、レールの中心線を選択できる。 |
| ↑ | uro:lod0Geometry | ■ |  |  |  | ↑ |
| ↑ | tran:lod1MultiSurface |  | ● |  |  |  |
| ↑ | tran:lod2MultiSurface |  |  | ● |  |  |
| ↑ | tran:lod3MultiSurface |  |  |  | ● |  |
| tran:TrafficArea |  |  |  | ● | ● |  |
| ↑ | tran:lod2MultiSurface |  |  | ● |  |  |
| ↑ | uro:lod2Network |  |  | ● |  | CompositeCurveとする。 |
| ↑ | tran:lod3MultiSurface |  |  |  | ● |  |
| ↑ | uro:lod3Network |  |  |  | ● | CompositeCurveとする。 |
| tran:AuxiliaryTrafficArea |  |  |  |  | ● |  |
| ↑ | tran:lod2MultiSurface |  |  |  |  | 対象外。 |
| ↑ | tran:lod3MultiSurface |  |  |  | ● |  |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

