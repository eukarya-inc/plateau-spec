# D.3.1.4 LOD3

## D.3.1.4 LOD3

交通（道路）モデル（LOD3）では、道路の形状を面として表現する。このとき、道路（tran:Road）は、横断構成要素であるtran:TrafficAreaとtran:AuxiliaryTrafficAreaに分解される。すなわち、tran:Roadの空間属性は、これを構成するtran:TrafficAreaとtran:AuxiliaryTrafficAreaの空間属性の集まりとなる。

**表 D-12 — D-7: LOD3 tran:Roadの空間属性一致条件**

| ID | `/req/tran/7` |
| --- | --- |
| 主題 | 妥当な交通（道路）データセットの要件 |
| 分類 | 要件分類D-1: 妥当な交通（道路）オブジェクト |
| 条文 | LOD3におけるtran:Roadの空間属性は、これを構成するtran:TrafficArea及びtran:AuxiliaryTrafficAreaの空間属性の集まりと一致しなければならない。 |

