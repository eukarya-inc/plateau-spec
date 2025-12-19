# D.3.1.3 LOD2

## D.3.1.3 LOD2

交通（道路）モデル（LOD2）では、道路の形状を面として表現する。このとき、道路（tran:Road）は、横断構成要素であるtran:TrafficAreaとtran:AuxiliaryTrafficAreaに分解される。すなわち、tran:Roadの空間属性は、これを構成するtran:TrafficAreaとtran:AuxiliaryTrafficAreaの空間属性の集まりとなる。

**表 D-11 — D-6: LOD2 tran:Roadの空間属性一致条件**

| ID | `/req/tran/6` |
| --- | --- |
| 主題 | 妥当な交通（道路）データセットの要件 |
| 分類 | 要件分類D-1: 妥当な交通（道路）オブジェクト |
| 条文 | LOD2におけるtran:Roadの空間属性は、これを構成するtran:TrafficArea及びtran:AuxiliaryTrafficAreaの空間属性の集まりと一致しなければならない。 |

