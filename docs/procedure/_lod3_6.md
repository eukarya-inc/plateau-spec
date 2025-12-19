# F.2.3.4 LOD3

## F.2.3.4 LOD3

交通（徒歩道）モデル（LOD3）では、徒歩道の形状を面として表現する。このとき、徒歩道（tran:Track）は、横断構成要素であるtran:TrafficAreaとtran:AuxiliaryTrafficAreaに分解される。すなわち、tran:Trackの空間属性は、これを構成するtran:TrafficAreaとtran:AuxiliaryTrafficAreaの空間属性の集まりとなる。

**表 F-10 — F-8: LOD3 tran:Trackの空間属性一致条件**

| ID | `/req/trk/8` |
| --- | --- |
| 主題 | 妥当な交通（徒歩道）データセットの要件 |
| 分類 | 要件分類F-1: 妥当な交通（徒歩道）オブジェクト |
| 条文 | LOD3におけるtran:Trackの空間属性は、これを構成するtran:TrafficArea及びtran:AuxiliaryTrafficAreaの空間属性の集まりと一致しなければならない。 |

このとき、交通（徒歩道）オブジェクトは、交通（徒歩道）モデル（LOD3）の定義に従ったものでなければならない。

**表 F-11 — F-9: 徒歩道のLOD3形状定義**

| ID | `/req/trk/9` |
| --- | --- |
| 主題 | 妥当な交通（徒歩道）データセットの要件 |
| 分類 | 要件分類F-1: 妥当な交通（徒歩道）オブジェクト |
| 条文 | 徒歩道のLOD3の形状は、交通（徒歩道）モデル（LOD3）の定義に従う。 |

