# 4.2.3.5.15 uro:IfcProjectedCRS

## 4.2.3.5.15 uro:IfcProjectedCRS

**表 4-72**

| 型の定義 | 投影座標参照系の情報を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcCoordinateReferenceSystem | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:name     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | 座標参照系の名称。 EPSG: `[EPSGコード]` `[EPSGコード]`は、EPSGにより指定された半角数字の組合せによる識別子とする。 |
| uro:description     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | EPSGコードの説明情報。 |
| uro:geodeticDatum     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | 測地原子の識別子。 JGD2011とする。 |
| uro:verticalDatum     [ IfcCoordinateReferenceSystem ] | xs::string [0..1] | 垂直原子。TPを原則とする。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:mapUnit     [ IfcProjectedCRS ] | xs::string [0..1] | 座標軸の単位。mとする。 |
| uro:mapProjection     [ IfcProjectedCRS ] | xs::string [0..1] | 投影座標系の名称。Japan Plane Rectangularとする。 |
| uro:mapZone     [ IfcProjectedCRS ] | xs::string [0..1] | 平面直角座標系の系。半角数字1～19までのいずれかとする。 |

