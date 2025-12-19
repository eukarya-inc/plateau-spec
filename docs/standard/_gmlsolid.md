# 4.28.2.1.7 gml:Solid

## 4.28.2.1.7 gml:Solid

**表 4-942**

| 型の定義 | 立体。以下を満たさなければならない。gml:Solidの境界を構成する曲面が、自己交差していない。gml:Solidは閉じている（水密である）。gml:Solidの内部が連続している。gml:Solidの境界を構成する曲面が、適切な方向を向いている。gml:Solidの境界を構成する曲面が、重なっていない。![Image](images/_gmlsolid_img001.png)妥当なgml:Solidの例 | ← |
| --- | --- | --- |
| 上位の型 | gml:_Solid | ← |
| ステレオタイプ | << Type >> | ← |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| gml:exterior             [ Solid ] | gml:_Surface [1..1] | 立体の外殻。gml:CompositeSurfaceを使用する。 |
| gml:interior             [ Solid ] | gml:_Surface [0..*] | 立体の内殻。gml:CompositeSurfaceを使用する。 |

