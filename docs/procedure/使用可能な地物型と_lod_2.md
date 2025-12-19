# F.2.2 使用可能な地物型とLOD

## F.2.2 使用可能な地物型とLOD

交通（徒歩道）モデルは、LODごとに、使用すべき地物型やその空間属性が異なる。

**表 F-2 — F-1: 徒歩道記述のLOD対応条件**

| ID | `/req/trk/1` |
| --- | --- |
| 主題 | 妥当な交通（徒歩道）データセットの要件 |
| 分類 | 要件分類F-1: 妥当な交通（徒歩道）オブジェクト |
| 条文 | 徒歩道の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

徒歩道の記述に使用するtran:Trackは、道路と同様に、tran:TrafficAreaとtran:AuxiliaryTrafficAreaの集まりとして構成される。tran:TrafficAreaは、歩道や車線のように、通行の用に供される領域である。一方、tran:AuxiliaryTrafficAreaは、縁石や植栽のように補助的な役割を果たす領域である。

tran:Trackの空間属性は、LOD0では、線として表現する。LOD1以降では面として表現する。tran:TrafficArea及びtran:AuxiliaryTrafficAreaは、いずれも面として表現する。ただし、これら二つの地物型はLOD2以降でのみ使用可能となる。すなわち、LOD1では徒歩道は歩車道の区別がない面として表現され、LOD2以降では道路の歩道や車道あるいは車線など、徒歩道内を区分できるようになる。また、LOD0で表現する線は、ネットワーク（中心線）又は徒歩道縁線のいずれかを選択できる。

交通（徒歩道）モデルの各LODにおいて使用可能な地物型と空間属性を、表F-3に示す。

**表 F-3 — 交通（徒歩道）モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| tran:Track |  | ● | ● | ● | ● |  |
| ↑ | tran:lod0Network | ■ |  |  |  | LOD0はネットワークを原則とするが、数値地形図との互換性を保つために、徒歩道縁を選択できる。 |
| ↑ | uro:lod0Geometry | ■ |  |  |  | ↑ |
| ↑ | tran:lod1MultiSurface |  | ● |  |  |  |
| ↑ | tran:lod2MultiSurface |  |  | ● |  |  |
| ↑ | tran:lod3MultiSurface |  |  |  | ● |  |
| tran:TrafficArea |  |  |  | ● | ● |  |
| ↑ | tran:lod2MultiSurface |  |  | ● |  |  |
| ↑ | tran:lod3MultiSurface |  |  |  | ● |  |
| tran:AuxiliaryTrafficArea |  |  |  | ● | ● |  |
| ↑ | tran:lod2MultiSurface |  |  | ● |  |  |
| ↑ | tran:lod3MultiSurface |  |  |  | ● |  |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

