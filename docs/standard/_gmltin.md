# 4.28.2.1.10 gml:Tin

## 4.28.2.1.10 gml:Tin

**表 4-945**

| 型の定義 | 不規則三角形網。 | ← |
| --- | --- | --- |
| 上位の型 | gml:TriangulatedSurface | ← |
| ステレオタイプ | << Type >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:stopLines     [ Tin ] | gml::LineStringSegment [0..*] | TINの生成を止める境界線。 |
| gml:breakLines     [ Tin ] | gml::LineStringSegment [0..*] | 地形の変化点をつなぐ線分。 |
| gml:maxLength     [ Tin ] | gml::LengthType [1..1] | TINを構成する三角形の最大辺長。 |
| gml:controlPoint     [ Tin ] | gml::posList [1..1] | TIN生成の制御点リスト。 |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| gml:patches             [ Surface ] | gml:_SurfacePatch [1..*] | 面を構成する要素（SurfacePatch）。標準製品仕様書では使用しない。 gml:Tinの場合は継承に制限がかかるため、gml:patchesではなくgml:trianglePatchesで部品となる要素を持つ。 |
| gml:trianglePatches             [ TriangulatedSurface ] | gml:Triangle [0..*] | 三角網を構成する三角形。 |

