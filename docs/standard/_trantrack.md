# 4.5.3.1.1 tran:Track

## 4.5.3.1.1 tran:Track

**表 4-237**

| 型の定義 | 徒歩道。徒歩道とは、人や車両等の通行の用に供される道路のうち、道路法第3 条に示された道路の種類及び建築基準法第42 条による「道路」を除く道路を指す。徒歩道には、作業規程の準則付録７公共測量標準図式における徒歩道及び庭園路（ただし、庭園路のうち、自動車ターミナル内の道路は、広場として取得するため、徒歩道には含まない）を含む。ここで、作業規程の準則付録７公共測量標準図式における徒歩道とは、「道路縁及び軽車道に接続するもの、登山、観光等に利用されるもの、神社等主要な地点へ到達するもの、耕地の区画等の景観を表現するために必要なもの」であり、庭園路とは、「公園、住宅地等で自動車の通行を規制している道路及び工場等特定の敷地内の道路」である。 徒歩道の延長方向は、以下の場所で区切る。交差部（四差路、多差路及び三差路）道路構造の変化点（トンネル、橋梁）位置正確度（地図情報レベル）や取得方法同一のLOD において、連続する徒歩道の境界は一致しなければならない。 tran:Track は、LOD0 ではネットワーク（中心線）又は徒歩道縁により取得する。  LOD1 以上では、面として取得する。  LOD2 以上では、tran:Track の面を、tran:TrafficArea とtran:AuxiliaryTrafficArea に細分する。 さらに、LOD3 では、各地物の面に高さを付与する。 以下に、取得例を示す。LOD0 における徒歩道の取得例![Image](images/_trantrack_img001.png)徒歩道モデル（LOD0）の取得例ネットワークで取得する場合は、徒歩道の中心線とする。徒歩道が道路と接する場合、道路中心線まで伸ばす。LOD1 における徒歩道の取得例![Image](images/_trantrack_img002.png)徒歩道モデル（LOD1）の取得例徒歩道のLOD1（面）は、徒歩道の境界に囲まれた範囲とする。徒歩道が道路と接する場合、その境界線は道路（tran:Road）の境界線と一致しなければならない。LOD2 における徒歩道の取得例![Image](images/_trantrack_img003.png)徒歩道モデル（LOD2）の取得例徒歩道のLOD2 は、LOD1（面）をtran:TrafficArea（車道、車道交差部、歩道）及びtran:AuxiliaryTrafficArea（島）に区分する。このとき、隣接するtran:TrafficArea 及びtran:AuxiliaryTrafficArea の面の境界線は座標が一致していなければならない。また、徒歩道が道路と接続する場合、接続する境界線は一致しなければならない。なお、歩道及び車道の区分が島又は路面標示により示されていない場合は、歩道として取得する。LOD3 における広場の取得例徒歩道のLOD3 は、LOD2 と同様に徒歩道の面をtran:TrafficArea 及びtran:AuxiliaryTrafficArea に区分する。このとき、それぞれの面は高さをもつ。また、LOD2 よりもさらに細かい種類にtran:TrafficArea及びtran:AuxiliaryTrafficArea を分けることができる。「高さの表現」及び「広場内の表現」の組み合わせにより、LOD3.0、LOD3.1、LOD3.2、LOD3.3 及び LOD3.4 に分かれるが、標準製品仕様は、原則としてLOD3.0 とする。LOD3.0 では、徒歩道の横断方向に一律の高さ（車道の高さ）を付し、高さの差は表現しない。![Image](images/_trantrack_img004.png)徒歩道モデル（LOD3.0）の取得例（横断面）徒歩道に車道が無い場合は歩道の高さとする。段の表現を行わないため、徒歩道に階段が存在する場合、階段の段は表現されず、最下段と最上段を結ぶ一定の斜度をもった面として表現される。![Image](images/_trantrack_img005.png)徒歩道モデル（LOD3.0）の取得例（階段がある場合）LOD2と同様、道路と接続する場合は、境界線が一致していなければならない。属性tran:functionは、コードリスト（Track_function.xml）より選択する。属性tran:usageは、標準製品仕様書では使用しない。 | ← |
| --- | --- | --- |
| 上位の型 | tran:TransportationComplex | ← |
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
| tran:class     [ TransportationComplex ] | gml::CodeType [0..1] | 交通の分類。コードリスト（TransportationComplex_class.xml）より選択する。航路の場合は使用しない。 |
| tran:function     [ TransportationComplex ] | gml::CodeType [0..*] | 機能の区分。使用する下位の地物型に応じたコードリストを選択する。 |
| tran:usage     [ TransportationComplex ] | gml::CodeType [0..*] | 利用方法の区分。使用する下位の地物型に応じたコードリストを選択する。 |
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
| tran:auxiliaryTrafficArea             [ TransportationComplex ] | tran:AuxiliaryTrafficArea [0..*] | 道路、徒歩道又は広場の場合は、これらを構成する要素のうち、交通領域の機能を補助する領域。鉄道の場合は、鉄道用地を構成する要素のうち、交通領域（軌道）の機能を補助する領域。航路の場合は、この関連役割は使用しない。 |
| tran:lod0Network             [ TransportationComplex ] | gml:GeometricComplex [0..*] | 各地物型により、tran:lod0Networkを以下の通り定義する。いずれも高さは0とする。 詳細な取得基準は、各地物型のLOD定義に従うこと。道路の場合は、道路の連続性を表現する線とする。鉄道の場合は、軌道中心線とする。軌道中心線は、路線ごとに作成し、路線に上下線等複数の軌道が存在する場合には、それぞれの軌道中心線を取得し、これらの組をgml:GeometricComplexとする。徒歩道の場合は、徒歩道の中心線及びこの端点と徒歩道に接続する道路のLOD0上の点を結ぶ線とする。広場の場合は、広場の中心線とする。広場の中心線は、広場の区域に含まれる道路縁又は庭園路等により示される2本の縁線の中心をつないだ線分とする。広場の中心線は、広場に接続する道路のLOD0上の点まで伸ばす。航路の区域の中心線とする。 |
| tran:lod1MultiSurface             [ TransportationComplex ] | gml:MultiSurface [0..1] | 各地物型により、tran:lod1MultiSurfaceを以下の通り定義する。詳細な取得基準は、各地物型のLOD定義に従うこと。道路の場合は、道路縁により囲まれた道路の範囲とする。車道交差部では、隅切りを結ぶ線により区切ることを基本とする。道路両側の隅切り位置が道路延長方向に大きく異なる場合は、より交差点より遠い隅切り位置より横断方向に区切る。隅切りが無い場合は、交差する道路の道路縁の接点を結ぶ線により区切る。鉄道の場合は、一対のレールとレールとの間の範囲とし、高さは0とする。路線ごとに作成し、路線に上下線等複数の軌道が存在する場合には、それぞれのレールに囲まれた領域を取得し、これらの組をgml:MultiSurfaceとする。徒歩道の場合は、徒歩道縁に囲まれた範囲とする。車道交差部では、隅切りを結ぶ線により区切ることを基本とする。徒歩道両側の隅切り位置が道路延長方向に大きく異なる場合は、より交差点より遠い隅切り位置より横断方向に区切る。隅切りが無い場合は、交差する道の道路縁の接点をつないだ境界で車道部を区切る。高さは0とする。広場の場合は都市計画において指定された区域に一致する。高さは0とする。航路の場合は、法令により定められた航路の区域とする。航路が交差する場合は、交差する部分の四隅を結ぶ位置で区切る。 |
| tran:lod2MultiSurface             [ TransportationComplex ] | gml:MultiSurface [0..1] | 各地物型により、tran:lod2MultiSurfaceを以下の通り定義する。高さは0とする。詳細な取得基準は、各地物型のLOD定義に従うこと。道路の場合は、道路縁により囲まれた道路の範囲とする。鉄道の場合は、道床の範囲とする。徒歩道の場合は、徒歩道縁に囲まれた範囲とする。広場の場合は広場の範囲とする。tran:lod1MultiSurfaceの範囲に一致する。航路の場合は、法令により定められた航路の区域。tran:lod1MultiSurfaceの範囲に一致する。tran:lod2MultiSurfaceの範囲は、各地物型が参照するtran:TrafficArea及びtran:AuxiliaryTrafficAreaのtran:lod2MultiSurfaceに含まれる、すべてのgml:Polygonにより構成しなければならない。 |
| tran:lod3MultiSurface             [ TransportationComplex ] | gml:MultiSurface [0..1] | 各地物型により、tran:lod3MultiSurfaceを以下の通り定義する。詳細な取得基準は、各地物型のLOD定義に従うこと。道路の場合は、道路縁により囲まれた道路の範囲とする。鉄道の場合は、鉄道用地の範囲とする。徒歩道の場合は、徒歩道縁に囲まれた範囲とする。広場の場合は広場の範囲とする。航路の場合は、標準製品仕様書ではtran:lod3MultiSurfaceは対象外とする。tran:lod3MultiSurfaceの範囲は、各地物型が参照するtran:TrafficArea及びtran:AuxiliaryTrafficAreaのtran:lod3MultiSurfaceに含まれる、すべてのgml:Polygonにより構成しなければならない。 |
| tran:trafficArea             [ TransportationComplex ] | tran:TrafficArea [0..*] | 道路、徒歩道又は広場の場合は、これらを構成する要素のうち、車両や人が通行可能な領域への参照。鉄道の場合は、鉄道用地を構成する要素のうち、車両の通行に使用する領域（軌道）への参照。航路の場合は、構成する要素のうち、船舶が航行可能な領域への参照。 |
| uro:tranDataQualityAttribute             [ TransportationComplex ] | uro:DataQualityAttribute [1..1] | 作成したデータの品質に関する情報。必須とする。 |
| uro:tranDmAttribute             [ TransportationComplex ] | uro:DmAttribute [0..*] | 公共測量標準図式による図形及び注記表現に必要な情報。航路の場合は使用しない。 |
| uro:tranFacilityAttribute             [ TransportationComplex ] | uro:FacilityAttribute [0..*] | uro:tranFacilityTypeAttribute.classによって指定された分野における施設管理情報。 |
| uro:tranFacilityIdAttribute             [ TransportationComplex ] | uro:FacilityIdAttribute [0..1] | uro:tranFacilityTypeAttribute.classによって指定された分野における施設の識別情報。 |
| uro:tranFacilityTypeAttribute             [ TransportationComplex ] | uro:FacilityTypeAttribute [0..*] | 特定分野における施設の分類情報。 |
| uro:tranKeyValuePairAttribute             [ TransportationComplex ] | uro:KeyValuePairAttribute [0..*] | 属性を拡張するための仕組み。コード値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:trackAttribute             [ Track ] | uro:TrackAttribute [0..1] | 徒歩道に関する追加情報。 |

