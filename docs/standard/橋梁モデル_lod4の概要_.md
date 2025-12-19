# 4.11.1.5.1 橋梁モデル（LOD4）の概要

## 4.11.1.5.1 橋梁モデル（LOD4）の概要

橋梁モデル（LOD4）は、橋梁の詳細な形状及び橋梁内部の空間を表現する。

橋梁モデル（LOD4）に含むべき地物を表4-500に示す。

**表 4-500 — 橋梁モデル（LOD4）に含むべき地物**

| 橋梁モデル（LOD4）に含むべき地物 | ← | 対応するCityGMLの地物型 | LOD4 |
| --- | --- | --- | --- |
| 橋梁 | ← | Bridge | ● |
| 橋梁部分 | ← | BridgePart | ■一つの橋梁を主題属性の異なる複数の部分に分ける場合は必須とする。 横断歩道橋・ペデストリアンデッキは必須とする。 |
| 屋根面 | ← | RoofSurface | ■屋根がある場合は必須とする。 |
| 外壁面 | ← | WallSurface | ● |
| 底面 | ← | GroundSurface | ● |
| 屋外天井面 | ← | OuterCeilingSurface | ● |
| 屋外床面 | ← | OuterFloorSurface | ■屋根が無い場合は必須とする。 |
| 構造上不可欠な要素 | ← | BridgeConstructionElement | ● |
| 外部付属物 | ← | BridgeInstallation | ● |
| 窓 | ← | Window | 〇 |
| 扉 | ← | Door | 〇 |
| 部屋 | ← | BridgeRoom | 〇 |
| 天井面 | ← | CeilingSurface | 〇 |
| 内壁面 | ← | InteriorWallSurface | 〇 |
| 床面 | ← | FloorSurface | 〇 |
| 閉鎖面 | ← | ClosureSurface | 〇 |
| 内部付属物 | ← | IntBridgeInstallation | 〇 |
| 家具 | ← | BridgeFurniture | 〇 |
| ●：必須■：条件付必須〇：任意 | ← | ← | ← |

橋梁の詳細な形状は、IFCにおける橋梁モデル(IFC Bridge)と整合を図る。

ただし、IFCにおける橋梁モデルはまだ国際標準となっていないことから、[IFC Bridge Fast Track Project Report WP2: Conceptual Model](https://www.buildingsmart.org/standards/domains/infrastructure/ifc-bridge/)を参考に、IFC Bridgeを構成するクラスとCityGMLのBridgeモジュールとの対応付けを示す（表4-501）。

なお、橋梁内部の空間は、建築物モデル（LOD4）と整合を図る。このときの橋梁内部の空間とは、橋梁にこれと一体的な建屋が設けられていた場合の、建屋の内部の空間を指す。

**表 4-501 — （参考）CityGMLのクラスとIFCのクラスとの対応**

| CityGMLの地物型 | 対応付けるIFCのクラス | 説明 |
| --- | --- | --- |
| Bridge | IfcBridgePart, IfcElementAssembly | IfcBridgePartのうち、以下に区分されるものを対象とする。DECK, DECK_SEGMENTIfcElementAssemblyのうち、以下に区分されるものを対象とする。DECK |
| BridgeConstructionElement | IfcBridgePart, IfcElementAssembly | IfcBridgePartのうち、以下に区分されるものを対象とする。ABUTMENT,FOUNDATION,PIER,PIER_SEGMENT,PYLON,SUBSTRUCTURE,SUPERSTRUCTURE,SURFACESTRUCTURE,IfcElementAssemblyのうち、以下に区分されるものを対象とする。ARCH,BEAM_GRID,GIRDER,REINFORCEMENT_UNIT,RIGID_FRAME,SLAB_FIELD,TRUSS,ABUTMENT,PIAR,PYLON,CROSS_BRACING, |
| BridgeInstallation | IfcElementAssembly | IfcElementAssemblyのうち、ACCESSORY_ASSEMBLYに区分されるものを対象とする。 |

CityGMLの地物型は、IFCにおいて部材の集まりを示すIfcBridgePart及びIfcElementAssemblyと対応付ける。

IFCでは、梁（IfcBeam）やスラブ（IfcSlab）、支承（IfcBearing）といった橋梁を構成する部材がクラスとして定義されているが、これらのクラスとCityGMLの地物型とを対応付けると、3D都市モデルとしては詳細すぎる表現となる。

そこで、これらの部材クラスの集まりであるIfcBridgePart及びIfcElementAssemblyとCityGMLの地物型とを対応付けた。このとき、IfcBridgePart及びIfcElementAssemblyの属性PredefinedTypeによりCityGMLの地物型であるbrid:Bridge、brid:BridgeConstructionElement又はbrid:BridgeInstallationへの振り分けを行っている。

![Image](images/橋梁モデル_lod4の概要__img001.png)

［出典：[[6]](https://www.buildingsmart.org/standards/domains/infrastructure/ifc-bridge/)］

