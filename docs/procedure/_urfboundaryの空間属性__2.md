# U.3.1 urf:Boundaryの空間属性

## U.3.1 urf:Boundaryの空間属性

urf:Boundaryは、urf:Zoneの境界となる。そのため、urf:Boundaryの空間属性（gml:MultiCurve）は、urf:Zoneの空間属性（gml:MultiSurface）を構成する曲線と一致しなければならない。

**表 U-8 — U-5: urf:Boundaryの重なり条件**

| ID | `/req/area/5` |
| --- | --- |
| 主題 | 妥当な区域データセットの要件 |
| 分類 | 要件分類U-1: 妥当な区域オブジェクト |
| 条文 | urf:Boundaryの空間属性（gml:MultiCurve）は、これをurf:boundaryによって参照するurf:Zoneの空間属性（gml:MultiSurface）に含まれるgml:Polygonの外周又は内周と重ならなければならない。 |

