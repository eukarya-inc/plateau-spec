# S.2.2 使用可能な地物型とLOD

## S.2.2 使用可能な地物型とLOD

地形モデルは、LODごとに使用すべき地物型やその空間属性を定めている。

**表 S-2 — S-1: 地形記述のLOD対応条件**

| ID | `/req/dem/1` |
| --- | --- |
| 主題 | 妥当な地形データセットの要件 |
| 分類 | 要件分類S-1: 妥当な地形オブジェクト |
| 条文 | 地形の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

地形の記述には、CityGMLのReliefモジュールに定義されたdem:ReliefFeatureを使用する。

地形はその表現方法に応じて地形要素（dem:_ReliefComponent）を選択する。標準製品仕様書では、地形要素として、dem:TinRelief、dem:MassPointRlief及びdem:BreaklineReliefを地形モデルに含めている。地形の表現では、これら三つの地形要素から一つを選択し、その集まりとして地形を表現する。

地形の各LODにおいて使用可能な地物型とその空間属性を表S-3に示す。

**表 S-3 — 地形モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | 適用 |
| --- | --- | --- | --- | --- | --- | --- |
| dem:ReliefFeature |  | ● | ● | ● | ● | dem:ReliefFeatureを作成する場合は、dem:BreaklineRelief、dem:TINRelief又はdem:ReliefFeatureのいずれかを必須とする |
| dem:BreaklineRelief |  | ■ |  |  |  |  |
| ↑ | dem:ridgeOrValleyLines | ■ |  |  |  | ↑ |
| dem:TINRelief |  | ■ | ■ | ■ | ■ | LOD1、LOD2及びLOD3では、dem:TINReliefを基本とする。 |
| ↑ | dem:tin | ■ | ■ | ■ | ■ | ↑ |
| dem:MassPointRelief |  | ■ | ■ | ■ | ■ | ↑ |
| ↑ | dem:reliefPoints | ■ | ■ | ■ | ■ | ↑ |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

3D都市モデルでは、地形モデルにdem:TINReliefを使用することを推奨している。この理由として、点群で扱う場合には、LAS形式など点群に特化したデータフォーマットが別途存在し、地形モデルには主題属性が付与されていないことから、CityGMLを使用して点群を記述する利点があまりないことが挙げられる。地形を連続した面として取り扱えるようにすることを重視し、dem:TINReliefの採用を推奨する。

