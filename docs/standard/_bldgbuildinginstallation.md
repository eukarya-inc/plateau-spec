# 4.2.3.1.15 bldg:BuildingInstallation

## 4.2.3.1.15 bldg:BuildingInstallation

**表 4-42**

| 型の定義 | 建築物の外側（屋外）に設置され、建築物の外観を特徴づける設備。建築物の付帯的な設備であり、主要な部分であってはならない。また、建築物（bldg:Building）と接していなければならない。 建築物の屋外付属物には以下を含む。ただし、全て屋外に設置され、建築物と接するもののみを対象とする。 バルコニー、ポーチ、アーケード、テラス、サンテラス、回廊、エントランスホール、ダクト、装飾的な柱、デッキ、屋根飾り、出窓、ドーマー、（建築物の一部としての）煙突、看板、換気口、（建築物の一部としての）塔、階段、カーポート、物置、アンテナ、外階段や歩道に設けられた屋根、手すり、スロープ、パネル（内装・外装の仕上げ等で利用される板材）、エレベータ、エスカレータ、動く歩道など。![Image](images/_bldgbuildinginstallation_img001.png)bldg:BuildingInstallationの例（左：屋根面に設置された建築物の屋外付属物 右：外壁面に設置された建築物の屋外付属物）ユースケースの要求に応じて、取得対象とする建築物の屋外付属物を限定してもよく、また、建築物の屋外付属物として取得せず建築物の一部として取得してもよい。 | ← |
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
| bldg:class     [ BuildingInstallation ] | gml::CodeType [0..1] | 建築物の屋外付属物の形態による区分。コードリスト（BuildingInstallation_class.xml）より選択する。建築物の外側に取り付けられた付属物の場合は、1000となる。 |
| bldg:function     [ BuildingInstallation ] | gml::CodeType [0..*] | 建築物の屋外付属物の主たる働き。コードリスト（BuildingInstallation_function.xml）より選択する。 |
| bldg:usage     [ BuildingInstallation ] | gml::CodeType [0..*] | 建築物の屋外付属物の主な使い道。標準製品仕様書では使用しない。 |
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
| bldg:boundedBy             [ BuildingInstallation ] | bldg:_BoundarySurface [0..*] | 建築物の屋外付属物を構成する外壁、屋根等の境界面への参照。建築物の屋外付属物の境界面が建築物（bldg:Building）の境界面となる場合にのみ作成する。 例えば、下図（平面図）のように建築物に建築物の屋外付属物があった場合、この建築物の屋外付属物を含む空間（gml:Solid）をBuildingとしたい場合は、建築物の屋外付属物の境界面を外壁面（bldg:WallSurface）とする。 ![Image](images/_bldgbuildinginstallation_img002.png)bldg:BuildingInstallationの境界面の例建築物の空間に建築物の屋外付属物を含まない場合は、建築物の屋外付属物を構成する面を、境界面（bldg:_BoundarySurface）に区別する必要はない。 |
| bldg:lod2Geometry             [ BuildingInstallation ] | gml:_Geometry [0..1] | 建築物の屋外付属物のLOD2の形状。 屋外付属物の外形（外側から見える形）を構成する面を取得し、面の各頂点に屋外付属物の高さを与える。各面は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう取得する。gml:MultiSurfaceを使用することを基本とする。容積の算出等、ユースケースで必要な場合はgml:Solidを使用する。![Image](images/_bldgbuildinginstallation_img003.png)bldg:BuildingInstallationの取得例（屋外階段） |
| bldg:lod3Geometry             [ BuildingInstallation ] | gml:_Geometry [0..1] | 建築物の屋外付属物のLOD3の形状。屋外付属物の外形（外側から見える形）を構成する面を取得し、面の各頂点に屋外付属物の高さを与える。各面は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう取得する。gml:MultiSurfaceを使用することを基本とする。容積の算出等、ユースケースで必要な場合はgml:Solidを使用する。 |
| bldg:lod4Geometry             [ BuildingInstallation ] | gml:_Geometry [0..1] | 建築物の屋外付属物のLOD4の外形。屋外付属物の外形（外側から見える形）を構成する面を取得し、面の各頂点に屋外付属物の高さを与える。各面は、データセットが採用する地図情報レベルの水平及び高さの誤差の標準偏差に収まるよう取得する。gml:MultiSurfaceにより記述することを基本とする。容積の算出等、ユースケースで必要な場合はgml:Solidを使用する。 |
| uro:ifcBuildingInstallationAttribute             [ BuildingInstallation ] | uro:IfcAttribute [0..*] | IDM・MVDで定義されるIFCに含まれる情報。bldg:BuildingInstallationに付与可能なデータ型は以下とする。uro:IfcBuildingElementこのとき、uro:IfcBuildingElement の属性uro:elementTypeの値はIfcBeam、IfcColumn、IfcPlate、IfcRailing、IfcRamp、IfcRampFlight、IfcSlab、IfcStair、IfcStairFlight、IfcBuildingElementProxy、IfcTransportElementのいずれかとなる。 |

