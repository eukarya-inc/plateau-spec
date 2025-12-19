# 4.12.3.1.1 tun:Tunnel

## 4.12.3.1.1 tun:Tunnel

**表 4-562**

| 型の定義 | トンネルとは、上方を含め周辺が地山や他の構造物で覆われている、交通、輸送等に供する構造物である。[[道路基盤地図情報（整備促進版） 製品仕様書 （案）](https://www.nilim.go.jp/lab/bcg/siryou/tnn/tnn0848pdf/ks084811.pdf)]標準製品仕様書では、その設置の形態から、ボックスカルバート、シェッド、シェルター及び地下横断歩道もトンネルに含める。 ボックスカルバート：他の道路等の下方を横断するため、その盛土の中を横断する形で設置される箱状の構造物。[ [道路基盤地図情報（整備促進版） 製品仕様書 （案）](https://www.nilim.go.jp/lab/bcg/siryou/tnn/tnn0848pdf/ks084811.pdf)]シェッド：落石及び雪崩等を道路外に直接落下させる、又は道路を越えて転進させるため、鋼材やコンクリート材等で道路上を覆う構造物。[ [道路基盤地図情報（整備促進版） 製品仕様書 （案）](https://www.nilim.go.jp/lab/bcg/siryou/tnn/tnn0848pdf/ks084811.pdf)]シェルター：アーチ型又は門型で、道路を完全に覆ったトンネル状の構造物。[ [道路基盤地図情報（整備促進版） 製品仕様書 （案）](https://www.nilim.go.jp/lab/bcg/siryou/tnn/tnn0848pdf/ks084811.pdf)]地下横断歩道：人、自転車等が道路又は鉄道を横断するために構築された地下道。 [ [作業規程の準則　付録7　公共測量標準図式](https://www.gsi.go.jp/gijyutukanri/gijyutukanri41018.html)]高速道路等に存在する延長の長いトンネルは、管理区間及び覆工スパンの境界で区切ることができる。 | ← |
| --- | --- | --- |
| 上位の型 | tun:_AbstractTunnel | ← |
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
| tun:class     [ _AbstractTunnel ] | gml::CodeType [0..1] | トンネルの形態による区分。コードリスト（Tunnel_class.xml ）より選択する。運用上必須とする。 |
| tun:function     [ _AbstractTunnel ] | gml::CodeType [0..*] | トンネルの主たる機能による区分。コードリスト（Tunnel_function.xml）より選択する。運用上必須とする。 |
| tun:usage     [ _AbstractTunnel ] | gml::CodeType [0..*] | トンネルの用途。tun:function で指定された機能と異なる場合に実際の用途を示すためにこの属性を用いる。コードリスト（Tunnel_usage.xml）より選択する。この属性を使用する場合は、コードリストを作成する。 |
| tun:yearOfConstruction     [ _AbstractTunnel ] | xs::gYear [0..1] | トンネルが建築された年度。完成年度とする。運用上必須とする。 |
| tun:yearOfDemolition     [ _AbstractTunnel ] | xs::gYear [0..1] | トンネルが解体された年度。 |
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
| tun:boundedBy             [ _AbstractTunnel ] | tun:_BoundarySurface [0..*] | トンネルを構成する屋根面（tun:RoofSurface）等の境界面。 |
| tun:consistsOfTunnelPart             [ _AbstractTunnel ] | tun:TunnelPart [0..*] | トンネルの部分（tun:TunnelPart）。トンネルの一部（tun:TunnelPart）の場合は使用しない。 |
| tun:interiorHollowSpace             [ _AbstractTunnel ] | tun:HollowSpace [0..*] | トンネルの内部空間（tun:HollowSpace）。 |
| tun:interiorTunnelInstallation             [ _AbstractTunnel ] | tun:IntTunnelInstallation [0..*] | トンネルの内部空間において、その外観を特徴付ける部分のうち、管理用通路・階段のような、トンネルの構造上不可欠ではない付属物（ tun:IntTunnelInstallation ）。内空（tun:HollowSpace ）に含まれない付属物を対象とする。 |
| tun:lod1MultiSurface             [ _AbstractTunnel ] | gml:MultiSurface [0..1] | 外周を構成する面の集まり。標準製品仕様書では使用しない。 |
| tun:lod1Solid             [ _AbstractTunnel ] | gml:Solid [0..1] | 外周に一律の高さを与えた立体。一律の高さは、トンネルの最も低い高さから最も高い高さまでとする。トンネルをtun:TunnelPartの集まりとして記述する場合、この空間属性は空となる。 |
| tun:lod1TerrainIntersection             [ _AbstractTunnel ] | gml:MultiCurve [0..1] | LOD1におけるトンネルと地形との交線。標準製品仕様書では使用しない。 |
| tun:lod2MultiCurve             [ _AbstractTunnel ] | gml:MultiCurve [0..1] | トンネルの立体表現に加え、線状の表現を行う場合に使用する。標準製品仕様書では使用しない。 |
| tun:lod2MultiSurface             [ _AbstractTunnel ] | gml:MultiSurface [0..1] | トンネルの主要構造の外形を構成する面の集まり。標準製品仕様書では使用しない。 |
| tun:lod2Solid             [ _AbstractTunnel ] | gml:Solid [0..1] | トンネルの主要構造の外形を示す立体。この時の立体は外壁等の、境界面により構成される。トンネルを tun:TunnelPart の集まりとして記述する場合、この空間属性は空となる。 |
| tun:lod2TerrainIntersection             [ _AbstractTunnel ] | gml:MultiCurve [0..1] | LOD2におけるトンネルと地形との交線。標準製品仕様書では使用しない。 |
| tun:lod3MultiCurve             [ _AbstractTunnel ] | gml:MultiCurve [0..1] | トンネルの立体表現に加え、線状の表現を行う場合に使用する。標準製品仕様書では使用しない。 |
| tun:lod3MultiSurface             [ _AbstractTunnel ] | gml:MultiSurface [0..1] | トンネルの主要構造の外形を構成する面の集まり。面の集まりが立体の境界としての要件を満たすことが出来ない場合に、LOD3の幾何オブジェクトとして使用する。 |
| tun:lod3Solid             [ _AbstractTunnel ] | gml:Solid [0..1] | トンネルの主要構造の外形を示す立体。この時の立体は、外壁等を区分する境界面及び開口部の面（境界面の内空として作成されている場合）により構成される。 |
| tun:lod3TerrainIntersection             [ _AbstractTunnel ] | gml:MultiCurve [0..1] | LOD3におけるトンネルと地形との交線。標準製品仕様書では使用しない。 |
| tun:lod4MultiCurve             [ _AbstractTunnel ] | gml:MultiCurve [0..1] | トンネルの立体表現に加え、線状の表現を行う場合に使用する。標準製品仕様書では使用しない。 |
| tun:lod4MultiSurface             [ _AbstractTunnel ] | gml:MultiSurface [0..1] | トンネルの外形を構成する面の集まり。このときの面は、外壁等を区分する境界面及び開口部の面（境界面の内空として作成されている場合）により構成される。lod4Solid又はlod4MultiSurfaceのいずれかとする。 |
| tun:lod4Solid             [ _AbstractTunnel ] | gml:Solid [0..1] | トンネルの詳細構造の外形を示す立体。このときの立体は、外壁等を区分する境界面及び開口部の面（境界面の内空として作成されている場合）により構成される。lod4Solid又はlod4MultiSurfaceのいずれかとする。 |
| tun:lod4TerrainIntersection             [ _AbstractTunnel ] | gml:MultiCurve [0..1] | LOD4におけるトンネルと地形との交線。標準製品仕様書では使用しない。 |
| tun:outerTunnelInstallation             [ _AbstractTunnel ] | tun:TunnelInstallation [0..*] | トンネルの外観を特徴付ける部分のうち、坑口や管理用通路、非常用階段のような、トンネルの構造上不可欠ではない付属物（ tun:TunnelInstallation ）。 |
| uro:tunBaseAttribute             [ _AbstractTunnel ] | uro:ConstructionBaseAttribute [0..1] | トンネルの管理に関する基本的な情報。 |
| uro:tunDataQualityAttribute             [ _AbstractTunnel ] | uro:DataQualityAttribute [0..1] | 作成するデータの品質に関する情報。原則必須とする。tun:TunnelPartが品質属性をもつ場合は、省略できる。関連役割uro:publicSurveyDataQualityAttributeは、属性uro:srcScaleLod0、uro:publicSurveySrcDescLod0以外は使用しない。 |
| uro:tunDisasterRiskAttribute             [ _AbstractTunnel ] | uro:DisasterRiskAttribute [0..*] | トンネルの災害リスクに関する情報。uro:DisasterRiskAttributeの下位型を使用して記述する。 tun:TunnelPartには使用しない。 |
| uro:tunDmAttribute             [ _AbstractTunnel ] | uro:DmAttribute [0..*] | LOD0の幾何形状。トンネルのLODの幾何形状は、gml:MultiPoint、gml:MultiCurve又はgml:MultiSurfaceにより記述する。 tun:TunnelPartには使用しない。 |
| uro:tunFacilityAttribute             [ _AbstractTunnel ] | uro:FacilityAttribute [0..*] | uro:tunFacilityTypeAttribute.classによって指定された分野における施設管理情報。トンネルの一部（tun:TunnelPart）の場合は使用しない。 |
| uro:tunFacilityIdAttribute             [ _AbstractTunnel ] | uro:FacilityIdAttribute [0..*] | uro:tunFacilityTypeAttribute.classによって指定された分野における施設の識別情報。トンネルの一部（tun:TunnelPart）の場合は使用しない。 |
| uro:tunFacilityTypeAttribute             [ _AbstractTunnel ] | uro:FacilityTypeAttribute [0..1] | 特定分野における施設の分類情報。トンネルの一部（tun:TunnelPart）の場合は使用しない。 |
| uro:tunFunctionalAttribute             [ _AbstractTunnel ] | uro:TunnelFunctionalAttribute [0..1] | トンネルの機能に関する情報。 |
| uro:tunKeyValuePairAttribute             [ _AbstractTunnel ] | uro:KeyValuePairAttribute [0..*] | 属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 tun:TunnelPartには使用しない。 |
| uro:tunRiskAssessmentAttribute             [ _AbstractTunnel ] | uro:ConstructionRiskAssessmentAttribute [0..1] | トンネルの損傷に関する情報。 |
| uro:tunStructureAttribute             [ _AbstractTunnel ] | uro:TunnelStructureAttribute [0..1] | トンネルの構造に関する情報。 |

