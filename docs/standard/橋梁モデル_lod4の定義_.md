# 4.11.1.5.2 橋梁モデル（LOD4）の定義

## 4.11.1.5.2 橋梁モデル（LOD4）の定義

橋梁モデル（LOD4）の定義として、含むべき地物型、各地物型の空間属性の型、取得基準、取得方法及び補足を示す。

**表 4-502**

| LOD |  | 地物型 | 空間属性の型 | 取得基準 | 取得方法 | 補足 |
| --- | --- | --- | --- | --- | --- | --- |
| LOD4 | ● | Bridge | MultiSurface | IfcBridgePartのうち、DECK又はDECK_SEGMENTに区分されるもの。 IfcElementAssemblyのうち、DECKに区分されるもの。 | 構成要素となる全ての部材の形状を統合し、面の集まりとして表現する。 |  |
| LOD4 | ■ | BridgePart | MultiSurfce | 一つの橋梁を、主題属性の異なる複数の部分に分ける場合に必須とする。階段やスロープのある横断歩道橋、ペデストリアンデッキ及び跨線橋の場合必須とする。 | 屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）及び閉鎖面（ClosureSurface）を境界面とする立体を作成する。 | 取得方法が異なる場合は必須とする。本体を構造上分けて作成したい場合は任意とする。横断歩道橋、ペデストリアンデッキ及び跨線橋の場合は階段、スロープ、踊り場を分ける。道路橋、桟道橋及び鉄道橋の場合、階段、スロープ、踊り場は付属物（BridgeInstallation）となる。 |
| LOD4 | ■ | RoofSurface | MultiSurface | 屋根が存在する場合 | 屋根の外周を取得し、棟（屋根の頂部であり、屋根の分水嶺となる箇所）及び谷（屋根と屋根のつなぎの谷状の部分）で区切る。区切った面の各頂点に屋根の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4 | ● | GroundSurface | MultiSurface |  | 橋梁の側面と、地表との交線により囲まれた面を取得する。面の各頂点に、地表面の高さを与える。 |  |
| LOD4 | ● | WallSurface | MultiSurface |  | 橋梁の外形のうち、屋根面（RoofSurface）、底面（GroundSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）及び閉鎖面（ClosureSurface）以外の面を取得する。 |  |
| LOD4 | ■ | ClosureSurface | MultiSurface | 窓や扉のない開口部 | 開口部の外周を面として取得する。 |  |
| LOD4 | ● | OuterCeilingSurface | MultiSurface |  | 橋梁の床版・主桁部分の外壁のうち、下向きとなる面の外周を取得する。面の各頂点に、外壁の高さを与える。 |  |
| LOD4 | ■ | OuterFloorSurface | MultiSurface | 屋根が存在しない場合 | 上部工の外周を取得する。床版の外周を取得する。外周の各頂点に床版の高さを与える。 |  |
| LOD4 | ● | BridgeConstructionElement | MultiSurface | IfcBridgePart IfcElementAssembly (ARCH, BEAM_GRID, GIRDER, REINFORCEMENT_UNIT, RIGID_FRAME, SLAB_FIELD, TRUSS, ABUTMENT, PIAR, PYLON, CROSS_BRACING) | 構造物の外形（外側から見える形）を構成する面を取得する。面の各頂点に構造物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4 | ● | BridgeInstallation | MultiSurface | IfcElementAssembly (ACCESSORY＿ASSEMBLY) | 屋外付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋外付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4 | 〇 | Door | MultiSurface | ユースケースの必要に応じて作成する。 | 扉（Door）の外周を取得する。 |  |
| LOD4 | 〇 | Window | MultiSurface | ユースケースの必要に応じて作成する。 | 窓（Window）の外周を取得する。 |  |
| LOD4 | 〇 | BridgeRoom | Solid又はMultiSurface | 全てを対象とする。 | 天井面（CeilingSurface）、内壁面（InteriorWallSurface）、閉鎖面（ClosureSurface）及び床面（FloorSurface）を境界面とする立体を作成する。 |  |
| LOD4 | 〇 | CeilingSurface | MultiSurface | 全てを対象とする。 | 天井の外周を取得する。 |  |
| LOD4 | 〇 | InteriorWallSurface | MultiSurface | 全てを対象とする。 | 部屋（Room）を区切る内壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4 | 〇 | FloorSurface | MultiSurface | 全てを対象とする。 | 床の外周を取得する。 |  |
| LOD4 | 〇 | IntBridgeInstallation | MultiSurface | 階段、スロープ、エスカレータ、輸送設備（エレベータ、エスカレータ、動く歩道）、柱、デッキ、ステージ、手すり、パネル、梁 | 屋内付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋内付属物の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。建築物モデル（LOD4）と同様とする。 |
| LOD4 | 〇 | CeilingSurface | MultiSurface | 全てを対象とする。 | 天井の外周を取得する。 |  |
| LOD4 | 〇 | InteriorWallSurface | MultiSurface | 全てを対象とする。 | 部屋（Room）を区切る内壁の角を結ぶ外周を取得する。角となる場所で区切る。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| LOD4 | 〇 | FloorSurface | MultiSurface | 全てを対象とする。 | 床の外周を取得する。 |  |
| LOD4 | 〇 | BridgeFurniture |  |  | 可動設備の外形（外側から見える形）を構成する面を取得する。面の各頂点に可動設備の高さを与える。 | 曲面の場合は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう平面に分割する。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← |

