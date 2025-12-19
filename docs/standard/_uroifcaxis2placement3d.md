# 4.2.3.5.1 uro:IfcAxis2Placement3D

## 4.2.3.5.1 uro:IfcAxis2Placement3D

**表 4-58**

| 型の定義 | ローカル座標系の変換を定義する座標系情報を設定するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:location     [ IfcAxis2Placement3D ] | Point [0..1] | 3 次元ローカル座標系における原点。 |
| uro:axis     [ IfcAxis2Placement3D ] | gml::doubleList [0..1] | ローカル座標系におけるZ 軸ベクトルを示し、アフィン変換のパラメータとして使用する。デフォルト値は (0, 0, 1)。uro:refDirection を設定した場合は必ず設定する。 |
| uro:refDirection     [ IfcAxis2Placement3D ] | gml::doubleList [0..1] | ローカル座標系におけるX 軸ベクトルを示し、アフィン変換のパラメータとして使用する。デフォルト値は(1, 0, 0)。uro:axis を設定した場合は必ず設定する。 |

