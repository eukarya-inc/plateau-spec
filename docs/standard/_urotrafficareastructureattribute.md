# 4.3.3.2.6 uro:TrafficAreaStructureAttribute

## 4.3.3.2.6 uro:TrafficAreaStructureAttribute

**表 4-158**

| 型の定義 | 交通領域の構造。 | ← |
| --- | --- | --- |
| 上位の型 | uro:TrafficAreaAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:numberOfLanes     [ TrafficAreaStructureAttribute ] | xs::integer [0..1] | 交通領域内の合計（一方通行区間の場合を除く）の車線数。道路構造令第 2 条第 7 号の登坂車線、同第 2 条第 6 号にいう付加追越車線、同第 2 条 8 号の屈折車線、同第 2 条第 9 号の変速車線及び同第 2 条第 14 号の停車帯、及びゆずり車線は車線数には含めない。交差点付近において、右左折のための車線が設けられている場合はこの数を含まない。  LOD2及びLOD3.0の車道部のみ（車線を分けない場合）にこの属性を付与する。 |

