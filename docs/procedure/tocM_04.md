# M.4 tun:HollowSpace

## M.4 tun:HollowSpace

tun:HollowSpaceは、トンネルの内部空間の記述に使用する地物型である。トンネルモデル（LOD4）においてのみ使用する。

## M.4.1 トンネル内部空間の空間属性

トンネルモデル（LOD4）において、トンネルの内部空間の形状は、立体により表現する。立体は、境界面に区分し、それぞれの境界面は、tun:CeilingSurface、tun:FloorSurface、tun:InteriorWallSurface又はtun:ClosureSurfaceのいずれかにならなければならない。

**表 M-17 — M-11: 境界面のポリゴンの条件**

| ID | `/req/tun/11` |
| --- | --- |
| 主題 | 妥当なトンネルデータセットの要件 |
| 分類 | 要件分類M-1: 妥当なトンネルオブジェクト |
| 条文 | tun:HollowSpaceのgml:Solidを構成する境界面のgml:Polygonは、以下のいずれかの地物型のLOD4幾何オブジェクトに含まれなければならない |
| A | tun:boundedByによりこのtun:HollowSpaceが参照する境界面（tun:_BoundarySurface）及びその開口部（tun:_Opening） |

