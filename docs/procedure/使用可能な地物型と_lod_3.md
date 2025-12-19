# G.2.2 使用可能な地物型とLOD

## G.2.2 使用可能な地物型とLOD

交通（広場）モデルは、LODごとに、使用すべき地物型やその空間属性が異なる。

**表 G-2 — G-1: 広場記述のLOD対応条件**

| ID | `/req/squr/1` |
| --- | --- |
| 主題 | 妥当な交通（広場）データセットの要件 |
| 分類 | 要件分類G-1: 妥当な交通（広場）オブジェクト |
| 条文 | 広場の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

広場は、CityGMLのTransportationモジュールに定義されたtran:Squareを用いて記述する。tran:Squareは広場面に該当する。 広場面は、歩道や車道のように、機能の異なる複数の要素から構成される。この広場を構成する要素は、tran:TrafficArea又はtran:AuxiliaryTrafficAreaを用いて記述する。tran:TrafficAreaは、車線や歩道のように、通行の用に供される領域である。一方、tran:AuxiliaryTrafficAreaは、車両待機所や車両乗降所のように補助的な役割を果たす領域である。tran:Squareは、tran:TrafficArea及びtran:AuxiliaryTrafficAreaの集まりとして構成される。

tran:Squareの空間属性は、LOD0では、線として表現する。LOD1以降では面として表現する。tran:TrafficArea及びtran:AuxiliaryTrafficAreaは、いずれも面として表現する。ただし、これら二つの地物型はLOD2以降でのみ使用可能となる。すなわち、LOD1では広場は歩車道の区別がない面として表現され、LOD2以降では広場の歩道や車道あるいは車線など、広場内の構造を区別できるようになる。広場の各LODにおいて使用可能な地物型と空間属性を、表G-3に示す。

**表 G-3 — 交通（広場）モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| tran:Square |  | ■ | ● | ● | ● | 広場が道路に接している場合はLOD0を必須とする。 |
| ↑ | lod0Network | ■ |  |  |  | LOD0はネットワークを原則とするが、数値地形図との互換性を保つために、道路縁及び車道の界線を選択できる。 |
| ↑ | uro:lod0Geometry | ■ |  |  |  |  |
| ↑ | lod1MultiSurface |  | ● |  |  |  |
| ↑ | lod2MultiSurface |  |  | ● |  |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |
| tran:TrafficArea |  |  |  | ● | ● |  |
| ↑ | lod2MultiSurface |  |  | ● |  |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |
| tran:AuxiliaryTrafficArea　 |  |  |  | ● | ● |  |
| ↑ | lod2MultiSurface　 |  |  | ● |  |  |
| ↑ | lod3MultiSurface |  |  |  | ● |  |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

