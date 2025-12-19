# L.4 brid:BridgeRoom

## L.4 brid:BridgeRoom

brid:BridgeRoomは、橋梁の内部空間の記述に使用する地物型である。橋梁モデル（LOD4）においてのみ使用する。brid:BridgeRoomは、橋梁と一体的な建屋があり、人や物の出入りが可能な場合に、その内部空間を表現するため使用する。

## L.4.1 橋梁内部空間の空間属性

橋梁モデル（LOD4）において、橋梁の内部空間の形状は、立体により表現する。立体は境界面に区分し、それぞれの境界面は、brid:CeilingSurface、brid:FloorSurface、brid:InteriorWallSurface又はbrid:ClosureSurfaceのいずれかにならなければならない。

**表 L-12 — L-7: 境界面ポリゴンの条件**

| ID | `/req/brid/7` |
| --- | --- |
| 主題 | 妥当な橋梁データセットの要件 |
| 分類 | 要件分類L-1: 妥当な橋梁オブジェクト |
| 条文 | brid:BridgeRoomのgml:Solidを構成する境界面のgml:Polygonは、以下のいずれかの地物型のLOD4幾何オブジェクトに含まれなければならない。 |
| A | brid:boundedByによりこのbrid:BridgeRoomが参照する境界面（brid:_BoundarySurface）の下位型 |
| B | 境界面に含まれる開口部（brid:_Opening） |

