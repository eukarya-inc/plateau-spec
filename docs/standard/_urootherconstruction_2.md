# 4.13.3.1.1 uro:OtherConstruction

## 4.13.3.1.1 uro:OtherConstruction

**表 4-618**

| 型の定義 | その他の構造物とは、構造物のうち、建築物、橋梁及びトンネルを除く構造物をいう。 「構造物」とは、「目的とする機能を持ち、作用に対して抵抗することを意図として人為的に構築されるもの」（土木・建築にかかる設計の基本、国土交通省）であり、ダム、堤防、床止め、堰、水門・閘門・陸閘、樋門・樋管、伏せ越し、水制、護岸、防波堤その他の土木構造物を指す。ダム：洪水の調整、発電、上水道、農工業等のための各種用水の貯水を目的として設けられた工作物をいい、砂防ダムを含む。[[作業規程の準則　付録7　公共測量標準図式](https://www.gsi.go.jp/gijyutukanri/gijyutukanri41018.html)]堤防：河川管理施設等構造令第17条に示される堤防及び霞堤。河川の流水の氾濫を防ぐ目的をもって、土砂・石礫等によって造られた河川構造物。床止め：河川管理施設等構造令第33条に示される床止め。河床の洗掘を防いで河川河道の勾配等を安定させ、河川の縦断又は横断形状を維持するために、河川を横断して設ける施設（河川構造物設計要領）堰：河川管理施設等構造令第36条に示される堰。河川の流水を制御するために、河川を横断して設けられるダム以外の施設であって、堤防の機能を有しないもの。水門・閘門・陸閘：河川管理施設等構造令第46条に示される水門。水門：河川又は水路を横断して設けられる制水施設であって堤防の機能を有するもの。閘門：運河・放水路などで水量を調節するための水門。陸閘：閉鎖することで海水の侵入を防止するとともに、開放することで堤防等の海側にある港湾、漁港、海浜等を利用するために人や車両等が堤防等を横断できるようにする施設。樋門・樋管：河川管理施設等構造令第46条に示される樋門。海・河川への排水口に設けられている施設。伏せ越し：河川管理施設等構造令第68条に示される伏せ越し。用水施設又は排水施設である開渠が、河川法の適用を受ける河川と交差する場合において、逆サイフォン構造で河底を横過する工作物で、施工方法が開削工法によるもの。水制：河川管理施設等構造令第26条に示される護岸。海岸や河川の水の勢いを弱め海岸や川岸が削られるのを防いだり、流れの方向を整えたりするために水中に設ける工作物。護岸：河川管理施設等構造令第25条に示される護岸。高水敷や他の構造物とともに流水による侵食作用から堤防（掘込河道にあっては堤内地）を保護するために設けるもの。防波堤：波浪を制御する堤防、埠頭、海岸浸食を防ぐ突堤等。[[作業規程の準則　付録7　公共測量標準図式](https://www.gsi.go.jp/gijyutukanri/gijyutukanri41018.html)]建築物は、bldg:Building、橋梁はbrid:Bridge、トンネルはtun:Tunnelとする。排水機場及び揚水機場はbldg:Buildingとする。 堤防のように延長が長く、構造上の切れ目なく続く場合は、管理区間及び市区町村界で区切ることができる。 | ← |
| --- | --- | --- |
| 上位の型 | uro:AbstractConstruction | ← |
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
| uro:conditionOfConstruction     [ AbstractConstruction ] | ConditionOfConstructionValue [0..1] | 構造物の稼働状況。 |
| uro:dateOfConstruction     [ AbstractConstruction ] | xs::date [0..1] | 構造物の完成年月日。 |
| uro:dateOfDemolition     [ AbstractConstruction ] | xs::date [0..1] | 構造物の解体年月日。標準製品仕様書では使用しない。 |
| uro:constructionEvent     [ AbstractConstruction ] | ConstructionEvent [0..*] | 構造物のライフサイクルにおける特定のイベントについての記述。標準製品仕様書では使用しない。 |
| uro:elevation     [ AbstractConstruction ] | Elevation [0..*] | 構造物の指定された地点における標高。 |
| uro:height     [ AbstractConstruction ] | Height [0..*] | 構造物の高さ。指定された2地点間の高さの差により記述する。 |
| uro:occupancy     [ AbstractConstruction ] | Occupancy [0..*] | 構造物に居住又は格納される人、動物、その他の移動可能な物体についての定量的な情報。標準製品仕様書では使用しない。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:class     [ OtherConstruction ] | gml::CodeType [0..1] | 構造物の分類。コードリスト（OtherConstruction_class.xml）より選択する。 |
| uro:function     [ OtherConstruction ] | gml::CodeType [0..*] | 構造物の機能。コードリスト（OtherConstruction_function.xml）より選択する。 |
| uro:usage     [ OtherConstruction ] | gml::CodeType [0..*] | 構造物の利用方法。標準製品仕様書では使用しない。 |
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
| uro:boundedBy             [ AbstractConstruction ] | uro:_BoundarySurface [0..*] | 構造物の立体を構成する境界面への参照。 |
| uro:consBaseAttribute             [ AbstractConstruction ] | uro:ConstructionBaseAttribute [0..1] | 構造物の管理に関する情報。 |
| uro:consDataQualityAttribute             [ AbstractConstruction ] | uro:DataQualityAttribute [0..1] | 作成するデータの品質に関する情報。必須とする。 |
| uro:consDisasterRiskAttribute             [ AbstractConstruction ] | uro:DisasterRiskAttribute [0..*] | その他の構造物の災害リスクに関する情報。uro:DisasterRiskAttributeの下位型を使用して記述する。 |
| uro:consDmAttribute             [ AbstractConstruction ] | uro:DmAttribute [0..*] | LOD0の幾何形状。 構造物のLODの幾何形状は、gml:MultiPoint、gml:MultiCurve又はgml:MultiSurfaceにより記述する。 |
| uro:consFacilityAttribute             [ AbstractConstruction ] | uro:FacilityAttribute [0..*] | uro:consFacilityTypeAttribute.classによって指定された分野における施設管理情報。 |
| uro:consFacilityIdAttribute             [ AbstractConstruction ] | uro:FacilityIdAttribute [0..1] | 構造物の位置を識別するための情報。 河川管理施設の場合は、uro:RiverFacilityIdAttribute及びその下位型を用いて記述する。 |
| uro:consFacilityTypeAttribute             [ AbstractConstruction ] | uro:FacilityTypeAttribute [0..*] | 構造物が管理・利用される分野（河川管理、港湾、漁港、農業等）における区分。 |
| uro:consKeyValuePairAttribute             [ AbstractConstruction ] | uro:KeyValuePairAttribute [0..*] | 属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 |
| uro:consStructureAttribute             [ AbstractConstruction ] | uro:ConstructionStructureAttribute [0..1] | 構造物の構造に関する情報。河川管理施設の場合は、uro:ConstructionStructureAttribute及びその下位型を用いて記述する。 |
| uro:constructionInstallation             [ AbstractConstruction ] | uro:ConstructionInstallation [0..*] | 構造物の付属物への参照。LOD3.1の場合にのみ使用する。 |
| uro:lod0Geometry             [ AbstractConstruction ] | gml:_Geometry [0..1] | 構造物の形状を2次元平面に投影した形状。点、線又は面により表現する。 |
| uro:lod1Geometry             [ AbstractConstruction ] | gml:_Geometry [0..1] | 構造物の形状を、立体として表現する。立体として表現する構造物のうち、護岸、樋門・樋管、水門・閘門、伏せ越し、堤防及びダムは構造物の上方からの正射影の外周を一律の高さで押し出した立体とする。また、水制及び床止めは一団となって設置された構造物の形状を包含する直方体とする。 高さは、構造物の最高高さとする。 |
| uro:lod2Geometry             [ AbstractConstruction ] | gml:_Geometry [0..1] | 構造物の形状を、主要な部分を簡略化した立体を組み合わせた立体（境界面は平面に分割）として表現する。簡略化した立体とは、球体、円錐、角錐、角柱、円柱などの単純な立体図形とする。構造上不可欠ではない付属物（手すり、柵、構造物と一体ではない階段）は表現しない。 |
| uro:lod3Geometry             [ AbstractConstruction ] | gml:_Geometry [0..1] | その他の構造物モデル（LOD3）では、その他の構造物の形状を、主要な部分の外形を構成する特徴点から構成する面を境界面とする立体として表現する。 LOD3は、構造上不可欠ではない付属物（手すり、柵、構造物と一体ではない階段）の表現有無によりLOD3.0及びLOD3.1に区分する。 |
| uro:lod4Geometry             [ AbstractConstruction ] | gml:_Geometry [0..1] | その他の構造物モデル（LOD4）の幾何形状。標準製品仕様書では使用しない。 |

