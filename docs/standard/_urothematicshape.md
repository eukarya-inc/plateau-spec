# 4.15.3.1.14 uro:ThematicShape

## 4.15.3.1.14 uro:ThematicShape

**表 4-710**

| 型の定義 | 地下埋設物の主題的な形状（高さをもった中心線又は中心点）を表すデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:horizontalType     [ ThematicShape ] | gml::CodeType [0..1] | uro:shapeがもつ水平位置の種類。標準製品仕様書では使用しない。 |
| uro:heightType     [ ThematicShape ] | gml::CodeType [0..1] | uro:shapeがもつ高さの種類。コードリスト（ThematicShape_heightType.xml）より選択する。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:shape             [ ThematicShape ] | gml:_Geometry [1..1] | 形状。リンクに対しては高さをもった中心線（gml:CompositeCurve）、ノードに対しては高さをもった中心点（gml:Point）とする。 |

