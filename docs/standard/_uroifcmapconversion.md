# 4.2.3.5.12 uro:IfcMapConversion

## 4.2.3.5.12 uro:IfcMapConversion

**表 4-69**

| 型の定義 | 座標参照系の変換情報を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:sourceCRS     [ IfcMapConversion ] | IfcCoordinateReferenceSystemSelect [0..1] | 変換元となる座標参照系の情報。 |
| uro:targetCRS     [ IfcMapConversion ] | IfcCoordinateReferenceSystem [0..1] | 変換先となる座標参照系の情報。 |
| uro:eastings     [ IfcMapConversion ] | gml::LengthType [0..1] | 変換先の座標参照系の座標系の東座標に沿った位置を指定する。右手デカルト座標系の場合、これは x 軸に沿った位置を定める。 |
| uro:northings     [ IfcMapConversion ] | gml::LengthType [0..1] | 変換先の座標参照系の座標系の北座標に沿った位置を指定する。右手デカルト座標系の場合、y 軸に沿った位置を定める。 |
| uro:orthogonalHeight     [ IfcMapConversion ] | gml::LengthType [0..1] | 変換先の座標参照系の垂直座標における位置（高さ）を指定する。右手デカルト座標系の場合、z軸に沿った位置を定める。 |
| uro:xAxisAbscissa     [ IfcMapConversion ] | xs::double [0..1] | 施工基準座標参照系のローカル x 軸の位置を示すベクトルの終点の東座標軸に沿った値を指定する。右手デカルト座標系の場合、x 軸に沿った位置を定める。XAxisOrdinate とともに、マップ座標系の水平面内のローカル x 軸の方向を提供する。 |
| uro:xAxisOrdinate     [ IfcMapConversion ] | xs::double [0..1] | 施工基準座標参照系のローカル x 軸の位置を示すベクトルの終点の北座標軸に沿った値を指定する。右手デカルト座標系の場合、y 軸に沿った位置を定める。XAxisAbscissaとともに、マップ座標系の水平面内のローカル x 軸の方向を提供する。 |
| uro:scale     [ IfcMapConversion ] | xs::double [0..1] | CRS の単位が施工基準座標系の単位と同一でない場合に使用されるスケール。省略した場合は1.0となる。 |

