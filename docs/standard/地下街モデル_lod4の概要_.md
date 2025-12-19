# 4.16.1.5.1 地下街モデル（LOD4）の概要

## 4.16.1.5.1 地下街モデル（LOD4）の概要

地下街モデル（LOD4）は、地下街モデル（LOD3）により表現される地下街の外側の形状に加え、地下街の内側の形状（屋内空間）を表現する。

地下街モデル（LOD4）は、含むべき地物により、LOD4.0、LOD4.1及びLOD4.2に区分する。これは、建築物モデル（LOD4）の区分と同一である。

標準製品仕様書では原則としてLOD4.0を採用する。ただし、ユースケースの必要に応じてLOD4.1又はLOD4.2を採用できる。

**表 4-737 — LOD4.0, LOD4.1及びLOD4.2の区分**

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

LOD4.0、LOD4.1及びLOD4.2それぞれの取得イメージを 表4-738に示す。

**表 4-738 — 地下街モデル（LOD4）の取得例**

| LOD | 取得イメージと説明 |
| --- | --- |
| LOD4.0 | ![Image](images/地下街モデル_lod4の概要__img001.png)LOD4.0は建築物の外形（図1）に加え、建築物の内部を部屋に区分する（図2）。このとき、各部屋の形状は立体として表現し、部屋の立体の境界面を、天井面、内壁面、床面又は閉鎖面のいずれかに区分する（図3）。また、天井面、内壁面又は床面に存在する全ての扉及び窓を表現する（図4）。 CityGMLでは、壁面や天井面などは全て面として表現する。1つの壁が建築物の外形を示す外壁と部屋の外形を示す内壁との機能を備えていた場合、建築物の外形となる面と部屋の外形となる面の2枚の面として表現され、それらの面の間には隙間ができる（何もない）。LOD4.0では地下街の内部に存在する付属物や家具を表現しないため、上図の例でも、付属物である階段、エレベータ、柱等が表現されていない。 なお、地下街の地上への出入口に設けられた建屋は、都市設備（CityFurniture）として取得する。 |
| LOD4.1 | ![Image](images/地下街モデル_lod4の概要__img002.png)LOD4.1ではLOD4.0に、屋内の付属物（bldg:IntBuildingInstallation）として、階段、スロープ、輸送設備（エスカレータ、エレベータ及び動く歩道）、柱及びデッキ・ステージが追加される。上図の例では、LOD4.0に加えて、階段、エスカレータ、スロープ及び柱が付属物として追加された。 |
| LOD4.2 | ![Image](images/地下街モデル_lod4の概要__img003.png)LOD4.2ではLOD4.1に屋内の付属物（bldg:IntBuildingInstallation）として、手すり、パネル及び梁が付属物として追加される。また、机やいすなどの移動可能な家具（bldg:BuildingFurniture）が追加される。上図の例では、LOD4.2に加えて付属物として階段の手すり及び部屋の間仕切りとしてパネル、また、家具としてテーブルやいす及び棚が追加された。 |

