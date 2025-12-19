# 4.2.3.1.16 bldg:IntBuildingInstallation

## 4.2.3.1.16 bldg:IntBuildingInstallation

**表 4-43**

| 型の定義 | 建築物の内側に設置された、恒久的に存在する固定的な設備（屋内付属物）。屋内付属物は、建築物の付帯的な設備であり、主要な部分であってはならない。また、屋内付属物は、建築物（bldg:Building）又は部屋（bldg:Room）と接していなければならない。![Image](images/_bldgintbuildinginstallation_img001.png)bldg:IntBuildingInstallationの例（階段、手すり）LOD4では、この屋内付属物を含む建築物に適用されたLOD4の細分に従い、以下を取得する。 LOD4.0：屋内付属物を取得しない（bldg:IntBuildingInstallationは取得しない）。  LOD4.1：階段、スロープ、輸送設備（エレベータ、エスカレータ及び動く歩道）、柱、デッキ・ステージ LOD4.2：LOD4.1の取得対象に加え、梁・手すり・パネル等の全ての建築物の屋外付属物及び全ての建築物の屋外付属物 | ← |
| --- | --- | --- |
| 上位の型 | core:_CityObject | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |
| core:creationDate     [ _CityObject ] | xs::date [0..1] | データが作成された日。運用上必須とする。 |
| core:terminationDate     [ _CityObject ] | xs::date [0..1] | データが削除された日。 |
| core:relativeToTerrain     [ _CityObject ] | core::RelativeToTerrainType [0..1] | 地表面との相対的な位置関係。標準製品仕様書では使用しない。 |
| core:relativeToWater     [ _CityObject ] | core::RelativeToWaterType [0..1] | 水面との相対的な位置関係。標準製品仕様書では使用しない。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| bldg:class     [ IntBuildingInstallation ] | gml::CodeType [0..1] | 屋内付属物の形態による区分。コードリスト（IntBuildingInstallation_class.xml）より選択する。 |
| bldg:function     [ IntBuildingInstallation ] | gml::CodeType [0..*] | 屋内付属物の主たる働き。コードリスト（IntBuildingInstallation_function.xml）より選択する。 |
| bldg:usage     [ IntBuildingInstallation ] | gml::CodeType [0..*] | 屋内付属物の主な使い道。標準製品仕様書では使用しない。 |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| core:externalReference             [ _CityObject ] | core:ExternalReference [0..*] | 外部への参照。標準製品仕様書では使用しない。 |
| gen:dateAttribute             [ _CityObject ] | gen:dateAttribute [0..*] | 日付型属性。属性を追加したい場合に使用する。 |
| gen:doubleAttribute             [ _CityObject ] | gen:doubleAttribute [0..*] | 実数型属性。属性を追加したい場合に使用する。 |
| gen:genericAttributeSet             [ _CityObject ] | gen:genericAttributeSet [0..*] | 汎用属性のセット。属性を追加したい場合に使用する。 |
| gen:intAttribute             [ _CityObject ] | gen:intAttribute [0..*] | 整数型属性。属性を追加したい場合に使用する。 |
| gen:measureAttribute             [ _CityObject ] | gen:measureAttribute [0..*] | 単位付き数値型属性。属性を追加したい場合に使用する。 |
| gen:stringAttribute             [ _CityObject ] | gen:stringAttribute [0..*] | 文字列型属性。属性を追加したい場合に使用する。 |
| gen:uriAttribute             [ _CityObject ] | gen:uriAttribute [0..*] | URI型属性。属性を追加したい場合に使用する。 |
| uro:pointCloud             [ _CityObject ] | uro:AbstractPointCloud [0..1] | ポイントクラウドへの参照。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| bldg:boundedBy             [ IntBuildingInstallation ] | bldg:_BoundarySurface [0..*] | 屋内付属物を構成する内壁、天井等の境界面への参照。屋内付属物の境界面が部屋（bldg:Room）の境界面となる場合にのみ作成する。例えば、下図（平面図）のように部屋内に屋内付属物があった場合、この屋内付属物を除く空間（gml:Solid）をRoomとしたい場合は、屋内付属物の境界面を内壁面（bldg:InteriorWallSurface）とする。![Image](images/_bldgintbuildinginstallation_img002.png)bldg:IntBuildingInstallationの境界面の取得例ただし、部屋の空間から屋内付属物を除く必要が無い場合は、屋内付属物の形状を構成する面を、境界面（bldg:_BoundarySurface）にする必要はない。また、ユースケースによりエレベータの出入口を、エレベータの扉を使って表現する必要がある場合は、bldg:boundedBy関連役割により、エレベータの扉が存在する境界面を内壁面（bldg:InteriorWallSurface）として区分し、この内壁面に扉（bldg:Door）を作成することでエレベータの扉を表現可能となる。 |
| bldg:lod4Geometry             [ IntBuildingInstallation ] | gml:_Geometry [0..1] | 屋内付属物のLOD4の外形。屋内付属物の外形（外側から見える形）を構成する面を取得する。面の各頂点に屋内付属物の高さを与える。gml:MultiSurfaceにより記述することを基本とする。容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。 |
| uro:ifcIntBuildingInstallationAttribute             [ IntBuildingInstallation ] | uro:IfcAttribute [0..*] | IDM・MVDで定義されるIFCに含まれる情報。bldg:IntBuildingInstallationに使用可能なデータ型は以下とする。uro:IfcBuildingElementこのとき、uro:IfcBuildingElement の属性uro:elementTypeの値は適用されたLOD4の詳細に応じて以下となる。LOD4.1： IfcRamp、IfcRampFlight、IfcStair、IfcStairFlight、IfcTransportElement 、IfcColumn、IfcBuildingElementProxyのいずれかとなる。LOD4.2：IfcBeam、IfcColumn、IfcPlate、IfcRailing、IfcRamp、IfcRampFlight、IfcStair、IfcStairFlight、IfcBuildingElementProxy、IfcTransportElementのいずれかとなる。 |
| uro:indoorIntInstallationAttribute             [ IntBuildingInstallation ] | uro:IndoorAttribute [0..*] | 屋内ナビゲーションに必要な情報。bldg:IntBuildingInstallationに付与可能なデータ型は以下とする。 uro:IndoorFurnishingAttribute  uro:IndoorTactileTileAttribute  uro:IndoorZoneAttribute  uro:IndoorUserDefinedAttribute |

