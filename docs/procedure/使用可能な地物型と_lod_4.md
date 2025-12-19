# H.2.1 使用可能な地物型とLOD

## H.2.1 使用可能な地物型とLOD

航路とは、船舶の通路として法令で定める海域である。標準製品仕様書では、原則として、以下に示す港則法や海上交通安全法によって規定される航路を対象とする。

- 港則法施行規則第8条
- 海上交通安全法第2条

航路の記述には、i-URのUrban Objectモジュールに定義された、uro:Waterwayを使用する。
 uro:Waterwayは、CityGML3.0において、新たに交通（Transportaiton）モジュールに定義された地物型であるtran:Waterwayを、CityGML2.0においても使えるように、i-URにおいて拡張した地物型である。uro:Waterwayは、道路（tran:Road）や鉄道（tran:Railway）と同様に、交通複合体（tran:TransportationComplex）を継承し、詳細度（LOD）ごとに、使用すべき空間属性が異なる。

標準製品仕様書では、交通（航路）モデルのLODは、LOD0からLOD2までを対象とする。

**表 H-2 — H-1: 航路記述のLOD対応条件**

| ID | `/req/wwy/1` |
| --- | --- |
| 主題 | 妥当な交通（航路）データセットの要件 |
| 分類 | 要件分類H-1: 妥当な交通（航路）オブジェクト |
| 条文 | 航路の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

航路（uro:Waterway）は、tran:TransportationComplexを継承していることから、tran:TrafficAreaとtran:AuxiliaryTrafficAreaの集まりによって構成される。tran:TrafficAreaは船舶が通行可能な領域であり、tran:AuxiliaryTrafficAreaはtran:TrafficAreaの機能を補助する地物である。しかし、航路は航行可能な区域のみが指定されていることから、航路の区域にはtran:AuxiliaryTrafficAreaに該当する領域は存在しない。そこで、標準製品仕様書ではtran:AuxiliaryTrafficAreaを対象外とする。

航路の各LODにおいて使用可能な地物型と空間属性を、表H-3に示す。

**表 H-3 — 交通（航路）モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| uro:Waterway |  | ● | ● | ● |  | LOD0、LOD1及びLOD2を対象とする。 |
| ↑ | tran:lod0Network | ● |  |  |  |  |
| ↑ | tran:lod1MultiSurface |  | ● |  |  |  |
| ↑ | tran:lod2MultiSurface |  |  | ● |  |  |
| ↑ | tran:lod3MultiSurface |  |  |  |  |  |
| tran:TrafficArea |  |  |  | ● |  |  |
| ↑ | tran:lod2MultiSurface |  |  | ● |  |  |
| ↑ | tran:lod3MultiSurface |  |  |  |  |  |
| tran:AuxiliaryTrafficArea |  |  |  |  |  | 対象としない。 |
| ↑ | tran:lod2MultiSurface |  |  |  |  |  |
| ↑ | tran:lod3MultiSurface |  |  |  |  |  |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

