# R.3.1.1.2 植生モデル（LOD0）の作成手順

## R.3.1.1.2 植生モデル（LOD0）の作成手順

- DMデータより、植生に該当するデータを抽出する。
植生モデル（LOD0）は、公共測量標準図式に従った表現である。よって、DMデータから、植生に該当する分類コードが付与されたデータを抽出する。 
表R-10に植生に関する公共測量標準図式の分類コードを示す。ただし、都市により独自の分類コードを使用していることがあるため、都市ごとの公共測量作業規程を確認する必要がある。

**表 R-10 — 植生に該当する公共測量標準図式の分類コード**

| 分類コード | 名称 | 対応するクラス |
| --- | --- | --- |
| 2238 | 並木 | veg::SolitaryVegetationObject |
| 4221 | 独立樹（広葉樹） | veg::SolitaryVegetationObject |
| 4222 | 独立樹（針葉樹） | veg::SolitaryVegetationObject |
| 6301 | 植生界 | veg::PlantCover |
| 6302 | 耕地界 | veg::PlantCover |
| 6311 | 田 | veg::PlantCover |
| 6313 | 畑 | veg::PlantCover |
| 6314 | さとうきび畑 | veg::PlantCover |
| 6315 | パイナップル畑 | veg::PlantCover |
| 6317 | 桑畑 | veg::PlantCover |
| 6318 | 茶畑 | veg::PlantCover |
| 6319 | 果樹園 | veg::PlantCover |
| 6321 | その他の樹木畑 | veg::PlantCover |
| 6323 | 芝地 | veg::PlantCover |
| 6331 | 広葉樹林 | veg::PlantCover |
| 6332 | 針葉樹林 | veg::PlantCover |
| 6333 | 竹林 | veg::PlantCover |
| 6334 | 荒地 | veg::PlantCover |
| 6335 | はい松地 | veg::PlantCover |
| 6336 | しの地（笹地） | veg::PlantCover |
| 6337 | やし科樹林 | veg::PlantCover |

