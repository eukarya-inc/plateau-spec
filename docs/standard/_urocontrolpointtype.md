# 4.4.3.2.7 uro:ControlPointType

## 4.4.3.2.7 uro:ControlPointType

**表 4-202**

| 型の定義 | 軌道中心線上の変化点の種類を指定するための共用型。いずれか一つの属性を選択する。 円曲線の変化点の場合は、uro:circularCurve、緩和曲線の変化点の場合はuro:slope、傾斜変化点の変化点の場合は、uro:slope、縦曲線の変化点の場合は、uro:verticalCurveを選択する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << Union >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:circularCurve     [ ControlPointType ] | CircularCurveType [1..1] | 円曲線のパラメータ。 |
| uro:transitionCurve     [ ControlPointType ] | TransitionCurveType [1..1] | 緩和曲線のパラメータ。 |
| uro:slopeType     [ ControlPointType ] | SlopeType [1..1] | 傾斜変化点のパラメータ。 |
| uro:verticalCurve     [ ControlPointType ] | VerticalCurveType [1..1] | 縦曲線のパラメータ。 |

