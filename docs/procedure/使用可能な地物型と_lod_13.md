# R.2.2 使用可能な地物型とLOD

## R.2.2 使用可能な地物型とLOD

植生モデルは、LODごとに、使用すべき地物型やその空間属性が異なる。

**表 R-2 — R-1: 植生記述のLOD対応条件**

| ID | `/req/veg/1` |
| --- | --- |
| 主題 | 妥当な植生データセットの要件 |
| 分類 | 要件分類R-1: 妥当な植生オブジェクト |
| 条文 | 植生の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

植生の記述には、CityGMLのVegetationモジュールに定義されたveg:SolitalyVegetationObject又はveg:PlantCoverを使用する。veg:SolitalyVegetationObjectは、街路樹のような単独木を記述するための型である。また、veg:PlantCoverは、歩道上の植栽など、個々の樹木を識別せず、まとめて記述する場合に使用する型である。

- 植生モデル（LOD0）では、植生の形状を点又は線で記述する。CityGMLのVegetationモジュールでは、植生のLODとしてLOD1以上が定義されている。植生モデルで使用する植生のLOD0は、i-URにより拡張した植生の空間属性である。
- 植生モデル（LOD1）では、veg:SolitalyVegetationObject及びveg:PlantCoverを立体として表現する。
- 植生モデル（LOD2）では、veg:SolitalyVegetationObject及びveg:PlantCoverを面の集まり又は立体として表現する。
- 植生モデル（LOD3）では、植生モデル（LOD2）と同様に、veg:SolitalyVegetationObject及びveg:PlantCoverを面の集まり又は立体として表現するが、その形状はより詳細化される。

**表 R-3 — 植生モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| veg:SolitaryVegetationObject |  | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  | Point又は MultiPointとする。 |
| ↑ | veg:lod1Geometry |  | ● |  |  | Solidとする。 |
| ↑ | veg:lod2Geometry |  |  | ● |  | MultiSurface又はSolidとする。 |
| ↑ | veg:lod3Geometry |  |  |  | ● | MultiSurface又はSolidとする。 |
| veg:PlantCover |  | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  | Point、 MultiPoint又はMultiCurveとする。 |
| ↑ | veg:lod1MultiSurface |  |  |  |  |  |
| ↑ | veg:lod1MultiSolid |  | ● |  |  |  |
| ↑ | veg:lod2MultiSurface |  |  | ■ |  | いずれかが必須となる。 |
| ↑ | veg:lod2MultiSolid |  |  | ■ |  | ↑ |
| ↑ | veg:lod3MultiSurface |  |  |  | ■ | いずれかが必須となる。 |
| ↑ | veg:lod3MultiSolid |  |  |  | ■ | ↑ |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

