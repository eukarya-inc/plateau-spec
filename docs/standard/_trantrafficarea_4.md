# 4.7.3.1.2 tran:TrafficArea

## 4.7.3.1.2 tran:TrafficArea

**表 4-287**

| 型の定義 | これを部品として保持する各地物型により、tran:TrafficAreaを以下の通り定義する。道路、徒歩道又は広場の場合車両や人が通行可能な領域とする。LOD2 及びLOD3.0 の場合は、車道部として、車両の利用が想定された車線や路肩その他一体的な舗装がされた全ての道路の部分を対象とする。また、歩道部として、歩道及び歩道上に設置された植栽の範囲を対象とする。![Image](images/_trantrafficarea_4_img001.png)LOD2及びLOD3.0におけるtran:TrafficAreaの例LOD3.1 の場合は、LOD3.0 の車道部のうち、車線を細分する。![Image](images/_trantrafficarea_4_img002.png)LOD3.1におけるtran:TrafficAreaの例LOD3.2 及びLOD3.3 の場合は、LOD3.1 の歩道部から歩道上の植栽を除いた範囲を歩道部とする。![Image](images/_trantrafficarea_4_img003.png)LOD3.2及びLOD3.3におけるtran:TrafficAreaの例LOD3.4の場合は、コードリストの区分に従う。![Image](images/_trantrafficarea_4_img004.png)LOD3.4におけるtran:TrafficAreaの例鉄道の場合軌道。軌道とは、施工基面上の道床（スラブを含む）、軌きょう（レールとまくらぎとを、はしご状に組み立てたもの。）及び直接これらに付帯する施設。[ JIS E 1001：2001]LOD2の場合は、軌道中心線に加え、道床の外周に囲まれた範囲を取得する。高さは0とする。![Image](images/_trantrafficarea_4_img005.png)LOD2における鉄道のtran:TrafficAreaの例LOD3.0の場合は軌道中心線に加え、道床の外周に囲まれた範囲を取得する。軌道中心線の各頂点には、軌道中心線上の勾配変化点の標高に基づき、高さを与える。また、道床の高さは、軌道中心線上の高さとする。![Image](images/_trantrafficarea_4_img006.png)LOD3.0における鉄道のtran:TrafficAreaの例LOD3.1 の場合は、LOD3.0 の軌道中心線、道床に加え、レールを取得する。高さはそれぞれの水平位置における標高とする。15 ㎝以上の高さの差を取得する。![Image](images/_trantrafficarea_4_img007.png)LOD3.1における鉄道のtran:TrafficAreaの例LOD3.2 の場合は、LOD3.1 の軌道中心線、道床及びレールの範囲を取得する。高さはそれぞれの水平位置における標高とする。15 ㎝未満の高さの差を取得する。![Image](images/_trantrafficarea_4_img008.png)LOD3.2におけるtran:TrafficAreaの例1 つの鉄道オブジェクトに含まれる交通領域は、属性の変化が無い限り、延長方向では区分しない。 LOD3 では、軌道中心線の平面線形が変化する位置（円曲線及び緩和曲線の開始地点及び終了地点）で区切る。航路の場合法令により指定された進行方向に区切った航路の部分。 高さは0とする。![Image](images/_trantrafficarea_4_img009.png)LOD2における航路のtran:TrafficAreaの例 | ← |
| --- | --- | --- |
| 上位の型 | tran:_TransportationObject | ← |
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
| tran:class     [ TrafficArea ] | gml::CodeType [0..1] | 交通の分類。標準製品仕様書では使用しない。 |
| tran:function     [ TrafficArea ] | gml::CodeType [0..*] | 交通領域の機能。コードリスト（TrafficArea_function.xml）より選択する。航路の場合は使用しない。 |
| tran:usage     [ TrafficArea ] | gml::CodeType [0..*] | 交通領域の利用方法。標準製品仕様書では使用しない。 |
| tran:surfaceMaterial     [ TrafficArea ] | gml::CodeType [0..1] | 表層舗装の有無及び材質。複数の表層舗装が混在している場合は、最も面積を占める舗装とする。道路、徒歩道及び広場の場合は、コードリスト（TrafficArea_surfaceMaterial.xml）より選択する。鉄道及び航路の場合は、標準製品仕様書では使用しない。 |
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
| tran:lod2MultiSurface             [ TrafficArea ] | gml:MultiSurface [0..1] | これを使用する各地物型により、交通領域のtran:lod2MultiSurfaceを以下の通り定義する。 詳細な取得基準は、各地物型のLOD定義に従うこと。道路、徒歩道及び広場の場合区画線や縁石等により示される境界線に囲まれた領域のうち、通行可能な道路の部分（歩道部、車道部、車道交差部）。高さは0とする。隣接するtran:TrafficArea又はtran:AuxiliaryTrafficAreaとの境界線の座標を一致させる。  tran:TrafficAreaのtran:lod2MultiSurfaceは、同一のtran:Roadのオブジェクトに含まれる他のtran:TrafficAreaやtran:AuxiliaryTrafficAreaのtran:lod2MultiSurfaceと重なることはない。（ただし、立体的な構造をもつ道路を除く） 車道交差部での区切りは、LOD1と同様とする。分離帯がある場合には、車道交差部の範囲を分離帯までとする。 境界線として区画線を使用する場合は、区画線の中心を境界線とする。鉄道の場合道床の正射影の外周を取得する。各頂点に軌道中心線上の高さを与える。航路の場合進行方向に区切られた航路の正射影の外周を取得する。高さは0とする。 |
| tran:lod3MultiSurface             [ TrafficArea ] | gml:MultiSurface [0..1] | これを使用する各地物型により、交通領域のtran:lod3MultiSurfaceを以下の通り定義する。詳細な取得基準は、各地物型のLOD定義に従うこと。道路、徒歩道及び広場の場合区画線や縁石等により示される境界線に囲まれた領域のうち、通行可能な道路の部分。 LOD3.0の場合、横断方向に連続する交通領域の高さは一律とし、車道の標高とする。  LOD3.1～LOD3.4では、各水平位置における標高とする。 隣接するtran:TrafficArea又はtran:AuxiliaryTrafficAreaとの境界線の座標を一致させる。  tran:TrafficAreaのtran:lod3MultiSurfaceは、同一のtran:Roadのオブジェクトに含まれる他のtran:TrafficAreaやtran:AuxiliaryTrafficAreaのtran:lod3MultiSurfaceと重なることはない。  LOD3.0の場合、車道交差部での区切りはLOD2と同様とする。LOD3.1～LOD3.4では、停止線がある場合にはこれの延長とし、停止線がない場合には、LOD2と同様とするが、ユースケースに応じて決定できる。 境界線として区画線を使用する場合は、区画線の中心を境界線とする。鉄道の場合LOD3.0の場合：軌道中心線、レールに囲まれた範囲及び道床を面として取得する。各頂点に、軌道中心線上の高さを与える。LOD3.1の場合：軌道中心線、レールに囲まれた範囲、道床及びレールを面として取得する。各頂点に、それぞれの水平位置における標高を与える。15㎝以上の高さの差を取得する。LOD3.2の場合：軌道中心線、レールに囲まれた範囲、道床及びレールを面として取得する。各頂点に、それぞれの水平位置における標高を与える。15㎝未満の高さの差を取得する。航路の場合使用しない。 |
| uro:railwayTrackAttribute             [ TrafficArea ] | uro:RailwayTrackAttribute [0..*] | 軌道中心線の線形情報。鉄道でのみ使用する。 |
| uro:trafficAreaStructureAttribute             [ TrafficArea ] | uro:TrafficAreaStructureAttribute [0..1] | 交通領域の構造。道路の交通領域の場合、かつ、車線を区分しない場合にのみ作成する。交通領域内の代表車線数を記述する。 |

