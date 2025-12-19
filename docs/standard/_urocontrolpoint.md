# 4.4.3.2.6 uro:ControlPoint

## 4.4.3.2.6 uro:ControlPoint

**表 4-201**

| 型の定義 | 軌道中心線上の変化点。平面線形の変化点、縦断線形の変化点及び勾配変化点を含む。線形の変化点は、開始点と終了点の対となり、勾配変化点の場合は変化した点1点を指す。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:startPost     [ ControlPoint ] | xs::string [0..1] | 平面線形の変化点及び縦断線形の変化点の場合は、円曲線、緩和曲線、縦曲線の開始点の起点からのキロ程。勾配変化点の場合は、勾配変化点の起点からのキロ程。 |
| uro:endPost     [ ControlPoint ] | xs::string [0..1] | 平面線形の変化点及び縦断線形の変化点の場合は、円曲線、緩和曲線、縦曲線の終了点の起点からのキロ程。勾配変化点の場合は不要とする。 |
| uro:function     [ ControlPoint ] | gml::CodeType [1..1] | 変化点の種類。コードリスト（ControlPoint_function.xml）から選択する。 |
| uro:parameter     [ ControlPoint ] | ControlPointType [1..1] | 変化点のパラメータ。変化点の種類に応じて、uro:ControlPointTypeの選択肢から一つを選択する。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:endPoint             [ ControlPoint ] | gml:Point [0..1] | 円曲線、緩和曲線又は縦曲線の終了点の座標値。軌道中心線上に存在しなければならない。 |
| uro:startPoint             [ ControlPoint ] | gml:Point [0..1] | 円曲線、緩和曲線、縦曲線の開始点又は勾配変化点の座標値。軌道中心線上に存在しなければならない。 |

