# 9.8.5.4 ImplicitGeometryにより地下埋設物の形状を表現する場合の関連役割

## 9.8.5.4 ImplicitGeometryにより地下埋設物の形状を表現する場合の関連役割

地下埋設物の形状を、ImplicitGeometryにより表現する場合、frn:CityFurnitureから継承する関連役割を使用する。

**表 9-14**

| 継承する関連役割 | ← | ← |
| --- | --- | --- |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| frn:lod2ImplicitRepresentation | core:ImplicitGeometry [0..1] | LOD2の幾何オブジェクトの代替として使用する繰り返しオブジェクト。 |
| frn:lod3ImplicitRepresentation | core:ImplicitGeometry [0..1] | LOD3の幾何オブジェクトの代替として使用する繰り返しオブジェクト。 |
| frn:lod4ImplicitRepresentation | core:ImplicitGeometry [0..1] | LOD4の幾何オブジェクトの代替として使用する繰り返しオブジェクト。 |

