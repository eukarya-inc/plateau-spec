# 4.2.3.1.1 bldg:Building

## 4.2.3.1.1 bldg:Building

**表 4-28**

| 型の定義 | 居住その他の目的をもって構築された建築物。普通建物、堅ろう建物、普通無壁舎及び堅ろう無壁舎に区分する。 普通建物とは、3階未満の建物及び3階以上の木造等で建築された建物をいう。 堅ろう建物とは、鉄筋コンクリート等で建築された建物で、地上3階以上又は3階相当以上の高さのものやスタンドを備えた競技場をいう。 普通無壁舎とは、側壁のない建物、温室及び工場内の建物類似の構築物で、3階未満のものをいう。 堅ろう無壁舎とは、鉄筋コンクリート等で建築された側壁のない建物及び建物類似の構築物で、地上3階以上又は3階相当以上の高さのものをいう。 ［出典：[作業規程の準則　付録7　公共測量標準図式](https://www.gsi.go.jp/gijyutukanri/gijyutukanri41018.html)］![Image](images/_bldgbuilding_img001.png)bldg:Buildingの例LOD0からLOD3 までは、建築物の屋外の形状を表現する。 LOD4では、建築物の屋外の形状に加え、屋内の形状を表現する。 | ← |
| --- | --- | --- |
| 上位の型 | bldg:_AbstractBuilding | ← |
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
| bldg:class     [ _AbstractBuilding ] | gml::CodeType [0..1] | 形態による区分。コードリスト（Building_class.xml）より選択する。地下街の場合は使用しない。 |
| bldg:function     [ _AbstractBuilding ] | gml::CodeType [0..*] | 主たる働き。標準製品仕様書では使用しない。 |
| bldg:usage     [ _AbstractBuilding ] | gml::CodeType [0..*] | 主な使い道。コードリスト（ Building_usage.xml）より選択する。用途の区分は、都市計画基礎調査実施要領（国土交通省都市局）による区分とする。複数の建築物で一体の施設を構成しているものについては、一体としての用途とする。店舗等併用住宅、同共同住宅、作業所併用住宅は、1/3 以上が住宅のものとする。複合用途の建築物（商業系複合施設及び併用住宅を除く）については、主たる用途により分類する。複数の用途を記述する場合は、主たる用途を最初に記載する。 |
| bldg:yearOfConstruction     [ _AbstractBuilding ] | xs::gYear [0..1] | 建築された年。 |
| bldg:yearOfDemolition     [ _AbstractBuilding ] | xs::gYear [0..1] | 解体された年。 |
| bldg:roofType     [ _AbstractBuilding ] | gml::CodeType [0..1] | 屋根形状の種類。 コ－ドリスト( Building_roofType.xml）より選択する。地下街の場合は使用しない。 |
| bldg:measuredHeight     [ _AbstractBuilding ] | gml::LengthType [0..1] | 計測により取得した建築物の地上の最低点から最高点までの高さ。単位はm (uom=”m”）とする。地下街の場合は使用しない。 |
| bldg:storeysAboveGround     [ _AbstractBuilding ] | xs::nonNegativeInteger [0..1] | 地上階の階数。地下街の場合は使用しない。 |
| bldg:storeysBelowGround     [ _AbstractBuilding ] | xs::nonNegativeInteger [0..1] | 地下階の階数。 |
| bldg:storeyHeightsAboveGround     [ _AbstractBuilding ] | gml::MeasureOrNullListType [0..1] | 地上の各階の高さを、地表面に最も近い階から列挙する。地下街の場合は使用しない。 |
| bldg:storeyHeightsBelowGround     [ _AbstractBuilding ] | gml::MeasureOrNullListType [0..1] | 地下の各階の高さを、地表面に最も近い階から列挙する。 |
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
| bldg:address             [ _AbstractBuilding ] | core:Address [0..*] | 建築物に付与された住所。 CityGMLでは複数個の記述が可能（多重度[0..*]）であるが、標準製品仕様書では、最大1個とする。 |
| bldg:boundedBy             [ _AbstractBuilding ] | bldg:_BoundarySurface [0..*] | 建築物を構成する外壁、屋根等の境界面。 |
| bldg:consistsOfBuildingPart             [ _AbstractBuilding ] | bldg:BuildingPart [0..*] | 階数や屋根の種別が異なる複合的な一つの建築物を、複数の建築物の集まりとして記述する場合の、部品となる建築物。 bldg:BuildingのLOD2、LOD3又はLOD4においてのみ使用できる。  bldg:BuildingPartには使用しない。 |
| bldg:interiorBuildingInstallation             [ _AbstractBuilding ] | bldg:IntBuildingInstallation [0..*] | 建築物の内部に付属する、階段、手すり、柱等の固定設備。建築物の内部の外観を特徴づける設備であり、恒久的に設置されている、固定されたもののみを対象とする。 なお、bldg:interiorBuildingInstallationを用いて記述する内部の固定設備は、個々の部屋（bldg:Room）に属さない設備を対象とする。 個々の部屋に付属する設備は、bldg:Roomのbldg:roomInstallationとして記述する。  bldg:interiorBuildingInstallationにより建築物内部の付属物を取得する場合、この建築物には、必ずLOD4の形状（bldg:lod4Solid又はbldg:lod4MultiSurface）が無ければならない。 |
| bldg:interiorRoom             [ _AbstractBuilding ] | bldg:Room [0..*] | 建築物の内部に存在する部屋。 bldg:interiorRoomにより建築物内部の部屋を取得する場合、この建築物には、必ずLOD4の形状（bldg:lod4Solid又はbldg:lod4MultiSurface）が無ければならない。 |
| bldg:lod0FootPrint             [ _AbstractBuilding ] | gml:MultiSurface [0..1] | 地表面と外壁面との交線に囲まれた面。 bldg:lod0FootPrint又はbldg:lod0RoofEdgeのいずれか一方が出現する。  bldg:lod0RoofEdgeを使用することを原則とする。 |
| bldg:lod0RoofEdge             [ _AbstractBuilding ] | gml:MultiSurface [0..1] | 建築物の上方からの正射影の外周。 bldg:lod0FootPrint又はbldg:lod0RoofEdgeのいずれか一方が出現する。 建築物のLOD0 の形状は、建築物モデル（LOD0）の定義に従う。 |
| bldg:lod1Solid             [ _AbstractBuilding ] | gml:_Solid [0..1] | 建築物の外周の上方からの正射影を取得し、地上から一律の高さを与えて上向きに押し出した立体。![Image](images/_bldgbuilding_img002.png)LOD1 立体イメージ一律の高さは中央値を原則とする。 |
| bldg:lod2MultiSurface             [ _AbstractBuilding ] | gml:MultiSurface [0..1] | 建築物の主要構造を保護又はこれに付随する設備の外形を示す面。 LOD2では、建築物をSolid により記述するため、MultiSurface は使用しない。 |
| bldg:lod2Solid             [ _AbstractBuilding ] | gml:_Solid [0..1] | 建築物の主要構造の外形を示す立体であり、屋根面（RoofSurface）、外壁面（WallSurface）及び底面（GroundSurface）を境界面とする。![Image](images/_bldgbuilding_img003.png)LOD2立体イメージ建築物をbldg:BuildingPart の集まりとして記述する場合、bldg:Buildingのlod2Solidは使用しない。 |
| bldg:lod3MultiSurface             [ _AbstractBuilding ] | gml:MultiSurface [0..1] | 建築物の主要構造を保護又はこれに付随する設備の外形を示す面。 LOD3では、建築物をSolid により記述するため、MultiSurface は使用しない。 |
| bldg:lod3Solid             [ _AbstractBuilding ] | gml:_Solid [0..1] | 建築物の詳細な形状を示す立体であり、屋根面（RoofSurface）、外壁面（WallSurface）、底面（GroundSurface）及び開口部の面（境界面の内空として作成されている場合）を境界面とする。![Image](images/_bldgbuilding_img004.png)LOD3立体イメージ建築物をbldg:BuildingPart の集まりとして記述する場合、bldg:Buildingのlod3Solidは使用しない。 |
| bldg:lod4MultiSurface             [ _AbstractBuilding ] | gml:MultiSurface [0..1] | 建築物の詳細な形状を示す面の集まりであり、屋根面（RoofSurface）、外壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、及び底面（GroundSurface）から構成する。 bldg:lod4Solid又はbldg:lod4MultiSurfaceのいずれかが出現する。  BIMモデルからの変換により取得する場合はMultiSurfaceとする。 |
| bldg:lod4Solid             [ _AbstractBuilding ] | gml:_Solid [0..1] | 建築物の詳細な形状を示す立体であり、屋根面（RoofSurface）、外壁面（WallSurface）、屋外床面（OuterFloorSurface）、屋外天井面（OuterCeilingSurface）、及び底面（GroundSurface）を境界面とする。 bldg:lod4Solid又はbldg:lod4MultiSurfaceのいずれかが出現する。 測量により取得する場合は、Solidとする。 |
| bldg:outerBuildingInstallation             [ _AbstractBuilding ] | bldg:BuildingInstallation [0..*] | 建築物の外側に付属する小屋根、外階段、バルコニー等の設備。建築物の外側の外観を特徴づける設備であり、恒久的に設置されているもののみを対象とする。 |
| uro:bldgDataQualityAttribute             [ _AbstractBuilding ] | uro:DataQualityAttribute [0..1] | 作成したデータの品質に関する情報。原則必須とする。bldg:Buildingがbldg:BuildingPartをもたない場合、bldg:Buildingの品質属性は必須とする。bldg:Buildingがbldg:BuildingPartをもち、bldg:BuildingPartが品質属性をもつ場合は、bldg:Buildingの品質属性は省略する。bldg:Buildingがbldg:BuildingPartをもち、bldg:BuildingPartが品質属性をもつ場合は、bldg:Buildingの品質属性は省略する。 |
| uro:bldgDisasterRiskAttribute             [ _AbstractBuilding ] | uro:DisasterRiskAttribute [0..*] | 当該建築物に対する災害リスクに関する情報。 i-URでは複数個の記述が可能（多重度[0..*]）であるが、標準製品仕様書では、uro:LandSlideRiskAttributeの出現回数は最大3回とする。  bldg:BuildingPartには使用しない。 |
| uro:bldgDmAttribute             [ _AbstractBuilding ] | uro:DmAttribute [0..*] | 公共測量標準図式による図形表現に必要な情報。建築物の一部（bldg:BuildingPart）には使用しない。 |
| uro:bldgFacilityAttribute             [ _AbstractBuilding ] | uro:FacilityAttribute [0..*] | uro:bldgFacilityTypeAttributeによって指定された分野における施設管理情報。 bldg:BuildingPartには使用しない。 |
| uro:bldgFacilityIdAttribute             [ _AbstractBuilding ] | uro:FacilityIdAttribute [0..1] | uro:bldgFacilityTypeAttribute.classによって指定された分野における施設の識別情報。 bldg:BuildingPartには使用しない。 |
| uro:bldgFacilityTypeAttribute             [ _AbstractBuilding ] | uro:FacilityTypeAttribute [0..*] | 特定分野における施設の分類情報。 bldg:BuildingPartには使用しない。 |
| uro:bldgKeyValuePairAttribute             [ _AbstractBuilding ] | uro:KeyValuePairAttribute [0..*] | コード型の属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 bldg:BuildingPartには使用しない。 |
| uro:bldgRealEstateIDAttribute             [ _AbstractBuilding ] | uro:RealEstateIDAttribute [0..1] | 建築物に紐づく不動産IDの情報。 |
| uro:bldgUsecaseAttribute             [ _AbstractBuilding ] | uro:BuildingUsecaseAttribute [0..1] | 建築物を使用するユースケースのための属性。標準製品仕様書では使用しない。 |
| uro:buildingDetailAttribute             [ _AbstractBuilding ] | uro:BuildingDetailAttribute [0..*] | 建築物に関する基礎的な情報。 bldg:BuildingPartにuro:buildingDetailAttributeが記述されている場合は、bldg:Buildingには出現しない。 |
| uro:buildingIDAttribute             [ _AbstractBuilding ] | uro:BuildingIDAttribute [1..1] | 建築物の識別情報。必ず1つ作成する。 |
| uro:ifcBuildingAttribute             [ _AbstractBuilding ] | uro:IfcAttribute [0..*] | IDM・MVDで定義されるIFCのクラス及びプロパティセットに含まれる情報。bldg:BuildingPartには使用しない。 bldg:Buildingに付与可能なデータ型は、以下とする。  uro:IfcProject  uro:IfcBuilding  uro:IfcSite  uro:IfcCoordinateReferenceSystem  uro:IfcProjectedCRS  uro:IfcMapConversion  uro:IfcPsetBuildingCommon  uro:IfcPsetSiteCommon |
| uro:indoorBuildingAttribute             [ _AbstractBuilding ] | uro:IndoorAttribute [0..*] | 屋内ナビゲーションに必要な情報。bldg:BuildingPartには使用しない。 bldg:Buildingに付与可能なデータ型は、以下とする。  uro:IndoorFacilityAttribute  uro:IndoorZoneAttribute  uro:IndoorUserDefinedAttribute |
| uro:largeCustomerFacilityAttribute             [ _AbstractBuilding ] | uro:LargeCustomerFacilityAttribute [0..*] | 当該建築物が大規模集客施設である場合の立地状況への参照。大規模集客施設の場合にのみ付与する。 bldg:BuildingPartには使用しない。 |

