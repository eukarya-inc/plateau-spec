# 4.15.3.1.6 uro:Duct

## 4.15.3.1.6 uro:Duct

**表 4-702**

| 型の定義 | トラフ、洞道、鞘管、CAB、情報BOX。管路やケーブルを保護するための設備。 | ← |
| --- | --- | --- |
| 上位の型 | uro:UtilityLink | ← |
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
| frn:class     [ CityFurniture ] | gml::CodeType [0..1] | 都市設備の区分。コードリスト（CityFurniture_class.xml）より選択する。道路付属物は、1000とする。 通信設備（例：電話ボックス、電柱）は1010とする。 保安設備（例：門、フェンス）は1020とする。 その他の施設は1030とする。 |
| frn:function     [ CityFurniture ] | gml::CodeType [0..*] | 都市設備の種類。コードリスト（CityFurniture_function.xml）より選択する。なお、区画線と道路標示について、形状・寸法・色が同一、かつ、設置される場所が同一のものは、区画線とする。 |
| frn:usage     [ CityFurniture ] | gml::CodeType [0..*] | 都市設備の主な使い道。コードリストより選択する。本属性を使用する場合は、コードリスト（CityFurniture_usage.xml）を作成する。 |
| uro:occupierType     [ UtilityNetworkElement ] | gml::CodeType [0..1] | 事業者の種類。コードリスト（UtilityNetworkElement_occupierType.xml）より選択する。 |
| uro:occupierName     [ UtilityNetworkElement ] | gml::CodeType [0..1] | 事業者の名称。コードリスト（UtilityNetworkElement_occupierName.xml）より選択する。 |
| uro:year     [ UtilityNetworkElement ] | xs::gYear [0..1] | 埋設された年度。 |
| uro:yearType     [ UtilityNetworkElement ] | gml::CodeType [0..1] | 埋設された年度の確からしさ。コードリスト（UtilityNetworkElement_yearType.xml）より選択する。 |
| uro:administrator     [ UtilityNetworkElement ] | gml::CodeType [0..1] | 埋設物の主管事業者。コードリスト（UtilityNetworkElement_administrator.xml）より選択する。 |
| uro:routeStartNode     [ UtilityLink ] | xs::string [0..1] | 路線又は系統の開始ノード。`[路線名]`又は`[路線]`-`[区間名]`とする。 |
| uro:startNode     [ UtilityLink ] | xs::string [0..1] | 開始ノードとなる設備（uro:Appurtenance）又は、格納設備（uro:UtilityNodeContainerの下位クラス）の識別子。 uro:Appurtenance又はuro:UtilityNodeContainerの下位クラスuro:FacilityIdAttributeの属性idの値に一致する。 |
| uro:routeEndNode     [ UtilityLink ] | xs::string [0..1] | 路線又は系統の開始ノード。`[路線名]`又は`[路線]`-`[区間名]`とする。 |
| uro:endNode     [ UtilityLink ] | xs::string [0..1] | 終了ノードとなる設備（uro:Appurtenance）又は、格納設備（uro:UtilityNodeContainerの下位クラス）の識別子。 uro:Appurtenance又はuro:UtilityNodeContainerの下位クラスuro:FacilityIdAttributeの属性idの値に一致する。 |
| uro:depth     [ UtilityLink ] | gml::LengthType [0..1] | 土被りの深さ。単位はm。 |
| uro:minDepth     [ UtilityLink ] | gml::LengthType [0..1] | 土被りの深さが幅をもつ場合の最小深さ。単位はm。 |
| uro:maxDepth     [ UtilityLink ] | gml::LengthType [0..1] | 土被りの深さが幅をもつ場合の最大深さ。単位はm。 |
| uro:maxWidth     [ UtilityLink ] | gml::LengthType [0..1] | 埋設物が存在する最大幅。単位はｍ。 |
| uro:offset     [ UtilityLink ] | gml::LengthType [0..1] | オフセット量。単位はm。 |
| uro:material     [ UtilityLink ] | gml::CodeType [0..1] | 材質の種類。コードリスト（UtilityNetworkElement_material.xml）より選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:width     [ Duct ] | gml::LengthType [0..1] | 外側の幅。単位は㎜とする。 |
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
| frn:lod1Geometry             [ CityFurniture ] | gml:_Geometry [0..1] | LOD1の幾何形状。形状をそれが占有している範囲（面）に一律の高さを与えて押し出した立体（gml:Solid）として表現する。ただし、他の地物の面と一体として表現されるべきものは面（gml:MultiSurface）として表現する。詳細は各モデルのLOD定義を参照すること。 |
| frn:lod1ImplicitRepresentation             [ CityFurniture ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD1）への参照。標準製品仕様書では埋設物モデルのみ使用することを可とする。 |
| frn:lod2Geometry             [ CityFurniture ] | gml:_Geometry [0..1] | LOD2の幾何形状。その主要な部分の外形を面の集まり又は立体として表現する。詳細は各モデルのLOD定義を参照すること。 |
| frn:lod2ImplicitRepresentation             [ CityFurniture ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD2）への参照。標準製品仕様書では埋設物モデルのみ使用することを可とする。 |
| frn:lod3Geometry             [ CityFurniture ] | gml:_Geometry [0..1] | LOD3の幾何形状。外形を面の集まり又は立体として表現する。詳細は各モデルのLOD定義を参照すること。 |
| frn:lod3ImplicitRepresentation             [ CityFurniture ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD3）への参照。標準製品仕様書では埋設物モデルのみ使用することを可とする。 |
| frn:lod4Geometry             [ CityFurniture ] | gml:_Geometry [0..1] | LOD4の幾何形状。外形に加え、内部を表現する。詳細は各モデルのLOD定義を参照すること。標準製品仕様書では地下埋設物の場合にのみ使用する。 |
| frn:lod4ImplicitRepresentation             [ CityFurniture ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD4）への参照。標準製品仕様書では埋設物モデルのみ使用することを可とする。 |
| uro:cityFurnitureDetailAttribute             [ CityFurniture ] | uro:CityFurnitureDetailAttribute [0..*] | 都市設備の詳細な内容。一つの道路標識柱に複数の道路標識が存在する場合のように、複合的な設備の場合は、設備ごとに記述する。 |
| uro:frnDataQualityAttribute             [ CityFurniture ] | uro:DataQualityAttribute [1..1] | 作成するデータの品質に関する情報。必須とする。 |
| uro:frnDmAttribute             [ CityFurniture ] | uro:DmAttribute [0..*] | 公共測量標準図式による図形表現に必要な情報。 |
| uro:frnFacilityAttribute             [ CityFurniture ] | uro:FacilityAttribute [0..*] | uro:frnFacilityTypeAttribute.classによって指定された分野における施設管理情報。 |
| uro:frnFacilityIdAttribute             [ CityFurniture ] | uro:FacilityIdAttribute [0..1] | uro:frnFacilityTypeAttribute.classによって指定された分野における施設の識別情報。 |
| uro:frnFacilityTypeAttribute             [ CityFurniture ] | uro:FacilityTypeAttribute [0..*] | 特定分野における施設の分類情報。 |
| uro:frnKeyValuePairAttribute             [ CityFurniture ] | uro:KeyValuePairAttribute [0..*] | コード型の属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 |
| uro:offsetDepth             [ UtilityNetworkElement ] | uro:OffsetDepth [0..*] | このリンク上に存在するオフセットデプス情報。 |
| uro:thematicShape             [ UtilityNetworkElement ] | uro:ThematicShape [0..*] | このリンクの主題的な形状情報。高さをもった埋設物の中心線の情報。 |
| uro:lengthAttribute             [ UtilityLink ] | uro:LengthAttribute [0..*] | このリンクの実長、亘長の情報。 |

