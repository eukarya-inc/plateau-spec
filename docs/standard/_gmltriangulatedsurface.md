# 4.28.2.1.9 gml:TriangulatedSurface

## 4.28.2.1.9 gml:TriangulatedSurface

**表 4-944**

| 型の定義 | 三角形網。 | ← |
| --- | --- | --- |
| 上位の型 | gml:Surface | ← |
| ステレオタイプ | << Type >> | ← |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| gml:patches             [ Surface ] | gml:_SurfacePatch [1..*] | 面を構成する要素（SurfacePatch）。標準製品仕様書では使用しない。 gml:Tinの場合は継承に制限がかかるため、gml:patchesではなくgml:trianglePatchesで部品となる要素を持つ。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| gml:trianglePatches             [ TriangulatedSurface ] | gml:Triangle [0..*] | 三角網を構成する三角形。 |

