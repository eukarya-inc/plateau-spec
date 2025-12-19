# 4.2.1.5.1 建築物モデル（LOD4）の概要

## 4.2.1.5.1 建築物モデル（LOD4）の概要

建築物モデル（LOD4）は、建築物モデル（LOD3）により表現される建築物の外側の形状に加え、建築物の内側の形状（屋内空間）を表現する。

建築物モデル（LOD4）は、BIMモデルからの変換又は屋内測量によって取得する。BIMモデルからの変換フローは、「[PLATEAU Handbook #03](https://www.mlit.go.jp/plateau/libraries/handbooks/)」を参照のこと。

建築物モデル（LOD4）は、含むべき地物により、LOD4.0、LOD4.1及びLOD4.2に区分する（表4-22）。

標準製品仕様書では原則としてLOD4.0を採用する。ただし、ユースケースの必要に応じてLOD4.1又はLOD4.2を採用できる。

**表 4-22 — LOD4.0, LOD4.1及びLOD4.2の区分**

| 建築物モデル（LOD4）に含むべき地物 | ← | 対応するCityGMLの地物型 | LOD4.0 | LOD4.1 | LOD4.2 |
| --- | --- | --- | --- | --- | --- |
| 建築物 | ← | bldg:Building | ● | ● | ● |
| 建築物部分 | ← | bldg:BuildingPart | ■一棟の建築物を、属性の異なる複数の部分に分ける場合に必須とする。 | ■一棟の建築物を、属性の異なる複数の部分に分ける場合に必須とする。 | ■一棟の建築物を、属性の異なる複数の部分に分ける場合に必須とする。 |
| 屋根面 | ← | bldg:RoofSurface | ● | ● | ● |
| 壁面 | ← | bldg:WallSurface | ● | ● | ● |
| 底面 | ← | bldg:GroundSurface | ● | ● | ● |
| 屋外天井面 | ← | bldg:OuterGroundSurface | 〇 | 〇 | 〇 |
| 屋外床面 | ← | bldg:OuterFloorSurface | 〇 | 〇 | 〇 |
| 屋外付属物 | ← | bldg:BuildingInstallation | ● | ● | ● |
| 部屋 | ← | bldg:Room | ● | ● | ● |
| 天井面 | ← | bldg:CeilingSurface | ● | ● | ● |
| 内壁面 | ← | bldg:InteriorWallSurface | ● | ● | ● |
| 床面 | ← | bldg:FloorSurface | ● | ● | ● |
| 閉鎖面 | ← | bldg:ClosureSurface | ■ BuildingPartを使用する場合、及び、内壁面、天井面、床面が無いが建築確認申請上部屋として区分されている空間を区切る場合に必須とする。 | ■ BuildingPartを使用する場合、及び、内壁面、天井面、床面が無いが建築確認申請上部屋として区分されている空間を区切る場合に必須とする。 | ■ BuildingPartを使用する場合、及び、内壁面、天井面、床面が無いが建築確認申請上部屋として区分されている空間を区切る場合に必須とする。 |
| 窓 | ← | bldg:Window | ● | ● | ● |
| 扉 | ← | bldg:Door | ● | ● | ● |
| 屋内付属物 | 階段 | bldg:IntBuildingInstallation |  | ● | ● |
| ↑ | スロープ | bldg:IntBuildingInstallation |  | ● | ● |
| ↑ | 輸送設備 | bldg:IntBuildingInstallation |  | ● | ● |
| ↑ | 柱 | bldg:IntBuildingInstallation |  | ● | ● |
| ↑ | デッキ・ステージ | bldg:IntBuildingInstallation |  | ● | ● |
| ↑ | 梁 | bldg:IntBuildingInstallation |  |  | 〇 |
| ↑ | パネル | bldg:IntBuildingInstallation |  |  | 〇 |
| ↑ | 手すり | bldg:IntBuildingInstallation |  |  | 〇 |
| 家具 | ← | bldg:BuildingFurniture |  |  | 〇 |
| 階 | ← | grp:CityObjectGroup | ● | ● | ● |
| 任意設定空間（例：防火区画） | ← | grp:CityObjectGroup |  |  | 〇 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← |

建築物モデル（LOD4）に含むべき地物を、図4-6に示す。

![Image](images/建築物モデル_lod4の概要__img001.png)

LOD4.0、LOD4.1及びLOD4.2それぞれの取得イメージを表4-23に示す。

**表 4-23 — 建築物モデル（LOD4）の取得イメージ**

| LOD | 取得イメージと説明 |
| --- | --- |
| LOD4.0 | ![Image](images/建築物モデル_lod4の概要__img002.png)LOD4.0は建築物の外形（上図1）に加え、建築物の内部を表現する。このとき、建築物の内部を部屋（bldg:Room）に区切り、各部屋の形状を立体として表現する（上図2）。また、部屋の立体の境界面を、天井面（bldg:CeilingSurface）、内壁面（bldg:InteriorWallSurface）、床面（bldg:FloorSurface）又は閉鎖面（bldg:ClosureSurface）のいずれかに区分する（上図3）。さらに、各部屋の天井面、内壁面又は床面に存在する扉（bldg:Door）及び窓（bldg:Window）を区分する（上図4）。閉鎖面は、内壁面や天井面、床面はないが、建築確認申請では部屋となっている空間を区切る場合に仮想的な境界面として使用する。 建築物の階を表現する場合は、CityObjectGroupを使用する。上図1のように、建築物が複数の階から構成される場合、上図4に示す同じ階の部屋を、CityObjectGroupを使用してグループ化する。このとき、CityObjectGroupの名称（gml:name）は階を識別する名称となる。 なお、CityGMLでは、壁面や天井面などは全て面として表現する。一方、現実世界の壁には厚みがある。1つの壁が建築物の外形を示す外壁と部屋の外形を示す内壁との機能を備えていた場合（上図5）、建築物の外形となる面（bldg:WallSurface）と部屋の外形となる面（bldg:InteriorWallSurface）の2枚の面として表現され、それらの面の間には隙間（壁の厚み）ができる（何もない）。 また、LOD4.0では建築物の内部に存在する付属物や家具を表現しない。 |
| LOD4.1 | ![Image](images/建築物モデル_lod4の概要__img003.png)LOD4.1ではLOD4.0に、屋内付属物（bldg:IntBuildingInstallation）として、階段、スロープ、輸送設備（エスカレータ、エレベータ及び動く歩道）、柱及びデッキ・ステージが追加される。上図の例では、LOD4.0に加えて、階段、踊り場、エレベータ、柱が付属物として追加された。 |
| LOD4.2 | ![Image](images/建築物モデル_lod4の概要__img004.png)LOD4.2ではLOD4.1に屋内の付属物（bldg:IntBuildingInstallation）として、手すり、パネル及び梁が付属物として追加される。また、机やいすなどの移動可能な家具（bldg:BuildingFurniture）が追加される。 上図の例では、LOD4.2に加えて屋内付属物として階段の手すりとパネル（間仕切り）、また、家具として机及び椅子が追加された。 |

