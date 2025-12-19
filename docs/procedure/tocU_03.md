# U.3 urf:Boundary

## U.3 urf:Boundary

urf:Boundaryは、区域の境界に属性をもたせたい場合に使用する型である。区域の境界に属性をもたせる必要が無い場合は、urf:Boundaryは使用しなくてよい。

## U.3.1 urf:Boundaryの空間属性

urf:Boundaryは、urf:Zoneの境界となる。そのため、urf:Boundaryの空間属性（gml:MultiCurve）は、urf:Zoneの空間属性（gml:MultiSurface）を構成する曲線と一致しなければならない。

**表 U-8 — U-5: urf:Boundaryの重なり条件**

| ID | `/req/area/5` |
| --- | --- |
| 主題 | 妥当な区域データセットの要件 |
| 分類 | 要件分類U-1: 妥当な区域オブジェクト |
| 条文 | urf:Boundaryの空間属性（gml:MultiCurve）は、これをurf:boundaryによって参照するurf:Zoneの空間属性（gml:MultiSurface）に含まれるgml:Polygonの外周又は内周と重ならなければならない。 |

## U.3.2 urf:Boundaryの主題属性

区域の境界に行政界や地番界、道路区域などが指定される場合がある。また、これらの線からの相対的な距離（オフセット）により指定される場合もある。このような境界線の種類やオフセット量をurf:Boundaryの主題属性として記述する。

