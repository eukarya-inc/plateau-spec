# Q.2.3.4 LOD4

## Q.2.3.4 LOD4

地下街モデル（LOD4）では、地下街モデル（LOD3）により表現される地下街の外側の形状に加え、地下街の内側の形状（屋内空間）を表現する。このとき、地下街オブジェクトは、地下街モデル（LOD4）の定義に従ったものでなければならない。

**表 Q-8 — Q-6: 地下街のLOD4形状定義**

| ID | `/req/ubld/6` |
| --- | --- |
| 主題 | 妥当な地下街データセットの要件 |
| 分類 | 要件分類Q-1: 妥当な地下街オブジェクト |
| 条文 | 地下街のLOD4の形状は、地下街モデル（LOD4）の定義に従う。 |

地下街モデル（LOD4）は、含むべき地物により、LOD4.0、LOD4.1及びLOD4.2に区分する（表Q-9）。これは、建築物モデル（LOD4）の区分と同一である。標準製品仕様書では原則としてLOD4.0を採用する。ただし、ユースケースの必要に応じてLOD4.1又はLOD4.2を採用できる。

**表 Q-9 — LOD4.0, LOD4.1及びLOD4.2の区分**

| 地下街モデル（LOD4）に含むべき地物 | 対応する地物型 | LOD4.0 | LOD4.1 | LOD4.2 |
| --- | --- | --- | --- | --- |
| 地下街 | uro:UndergroundBuilding | ● | ● | ● |
| 建築物部分 | bldg:BuildingPart | ■ | ■ | ■ |
| 屋根面 | bldg:RoofSurface | ● | ● | ● |
| 外壁面 | bldg:WallSurface | ● | ● | ● |
| 底面 | bldg:GroundSurface | ● | ● | ● |
| 屋外天井面 | bldg:OuterCeilingSurface |  |  |  |
| 屋外床面 | bldg:OuterFloorSurface |  |  |  |
| 屋外付属物 | bldg:BuildingInstallation |  |  |  |
| 部屋 | bldg:Room | ● | ● | ● |
| 天井面 | bldg:CeilingSurface | ● | ● | ● |
| 内壁面 | bldg:InteriorWallSurface | ● | ● | ● |
| 床面 | bldg:FloorSurface | ● | ● | ● |
| 閉鎖面 | bldg:ClosureSurface | ● | ● | ● |
| 窓 | bldg:Window | ● | ● | ● |
| 扉 | bldg:Door | ● | ● | ● |
| 階段 | bldg:IntBuildingInstallation |  | ● | ● |
| スロープ | bldg:IntBuildingInstallation |  | ● | ● |
| 輸送設備 | bldg:IntBuildingInstallation |  | ● | ● |
| 柱 | bldg:IntBuildingInstallation |  | ● | ● |
| デッキ・ステージ | bldg:IntBuildingInstallation |  | ● | ● |
| 梁 | bldg:IntBuildingInstallation |  |  | 〇 |
| パネル | bldg:IntBuildingInstallation |  |  | 〇 |
| 手すり | bldg:IntBuildingInstallation |  |  | 〇 |
| 家具 | bldg:BuildingFurniture |  |  | 〇 |
| 階 | grp:CityObjectGroup | ● | ● | ● |
| 任意設定空間（例：防火区画） | grp:CityObjectGroup |  |  | 〇 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← |

