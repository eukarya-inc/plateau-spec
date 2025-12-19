# H.2.3.3 LOD2

## H.2.3.3 LOD2

交通（航路）モデル（LOD2）では、航路（uro:Waterway）の形状を面として記述する。LOD2では、uro:Waterwayはtran:TrafficAreaに区分できる。このとき、uro:Waterwayの空間属性は、これを構成するtran:TrafficAreaの空間属性の集まりとなる。

**表 H-6 — H-4: LOD2 uro:Waterwayの空間属性一致条件**

| ID | `/req/wwy/4` |
| --- | --- |
| 主題 | 妥当な交通（航路）データセットの要件 |
| 分類 | 要件分類H-1: 妥当な交通（航路）オブジェクト |
| 条文 | LOD2におけるuro:Waterwayの空間属性は、これを構成するtran:TrafficAreaの空間属性の集まりと一致しなければならない。 |

tran:TrafficAreaは、法令により指定された船舶の航行方向により区分する。航行方向の指定が無い場合、交通（航路）モデル（LOD2）においてtran:TrafficAreaの形状と、uro:Waterwayの形状は等しくなり、これは交通（航路）モデル（LOD1）と一致する。

