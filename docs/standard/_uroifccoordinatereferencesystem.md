# 4.2.3.5.6 uro:IfcCoordinateReferenceSystem

## 4.2.3.5.6 uro:IfcCoordinateReferenceSystem

**表 4-63**

| 型の定義 | 座標参照系の情報を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:name     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | 座標参照系の名称。 EPSG: `[EPSGコード]` `[EPSGコード]`は、EPSGにより指定された半角数字の組合せによる識別子とする。 |
| uro:description     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | EPSGコードの説明情報。 |
| uro:geodeticDatum     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | 測地原子の識別子。 JGD2011とする。 |
| uro:verticalDatum     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | 垂直原子。TPを原則とする。 |

