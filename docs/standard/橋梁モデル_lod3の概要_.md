# 4.11.1.4.1 橋梁モデル（LOD3）の概要

## 4.11.1.4.1 橋梁モデル（LOD3）の概要

橋梁モデル（LOD3）では、橋梁の形状を、主要な部分の外形を構成する特徴点から構成する面を境界面とする立体、又は面の集まりとして表現する。

橋梁モデル（LOD3）に含むべき地物を表4-497に示す。

**表 4-497 — 橋梁モデル（LOD3）に含むべき地物と対応するCityGMLの地物型**

| その他の構造物モデル（LOD3）に含むべき地物 | ← | 対応するCityGMLの地物型 | LOD3 |
| --- | --- | --- | --- |
| 橋梁 | ← | Bridge | ● |
| 橋梁部分 | ← | BridgePart | ■一つの橋梁を主題属性の異なる複数の部分に分ける場合は必須とする。 横断歩道橋・ペデストリアンデッキでは必須とする。 |
| 屋根面 | ← | RoofSurface | ■屋根がある場合は必須とする。 |
| 底面 | ← | GroundSurface | ● |
| 外壁面 | ← | WallSurface | ● |
| 閉鎖面 | ← | ClosureSurface | ■ BridgePartを使用する場合は必須とする。 扉のない出入口がある場合は必須とする。 |
| 屋外床面 | ← | OuterFloorSurface | ■屋根が無い場合は必須とする。 |
| 屋外天井面 | ← | OuterCeilingSurface | ■ WallSurfaceとの区分が必要な場合に必須とする。 |
| 構造上不可欠な要素 | トラス、アーチ、パイロン | BridgeConstructionElement | ● |
| ↑ | 吊材、ケーブル | BridgeConstructionElement | ● |
| ↑ | 橋脚、橋台 | BridgeConstructionElement | ■ |
| 屋外付属物 | 高欄、地覆、親柱、庇、庇の柱、エレベータ、エスカレータ、手すり | BridgeInstallation | ● |
| ↑ | 階段、踊り場、スロープ | BridgeInstallation | ● |
| 扉 | ← | Door | 〇 |
| 窓 | ← | Window | 〇 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← |

**表 4-498 — 橋梁モデル（LOD3）の取得イメージ**

|  | LOD3 | ← |
| --- | --- | --- |
| 取得例 | ![Image](images/橋梁モデル_lod3の概要__img001.png) | ![Image](images/橋梁モデル_lod3の概要__img002.png) |
| 説明 | 道路橋及び鉄道橋の場合は、床版及び主桁以外の構造上不可欠な部材をBridgeConstructionElementとして取得する。上図の例では橋脚が該当する。それ以外の橋梁の外観を構成する部材をBridgeInstallationとして取得する。上図の例では高欄が該当する。 | 跨線橋の場合は、道路橋及び鉄道橋と同様に、床版及び主桁以外の構造上不可欠な部材をBridgeConstructionElementとして取得する。上図の例では橋脚が該当する。それ以外の橋梁の外観を構成する部材をBridgeInstallationとして取得する。上図の例では高欄が該当する。 |
|  | LOD3 | ← |
| 取得例 | ![Image](images/橋梁モデル_lod3の概要__img003.png) | ![Image](images/橋梁モデル_lod3の概要__img004.png) |
| 説明 | ケーブル橋の場合、パイロン、ケーブル及び吊材を構造上不可欠な部材（BuildingConstructionElement）として取得する。この時、吊材は一本ずつ取得せず、吊材が存在する範囲をまとめて一つの面として取得してもよい。 | 横断歩道橋、ペデストリアンデッキ及び跨線橋の場合は、本体（上部工、階段及び踊り場）以外の構造上不可欠な部材をBridgeConstructionElementとして取得する。上図の例では橋脚が該当する。それ以外の橋梁の外観を構成する部材をBridgeInstallationとして取得する。上図の例では高欄が該当する。横断歩道橋、ペデストリアンデッキ及び跨線橋の本体（上部工、階段及び踊り場）に屋根がある場合、庇はBridgeInstallationとして取得する。 |

