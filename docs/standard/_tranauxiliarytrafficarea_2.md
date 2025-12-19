# 4.4.3.1.3 tran:AuxiliaryTrafficArea

## 4.4.3.1.3 tran:AuxiliaryTrafficArea

**表 4-195**

| 型の定義 | これを部品として保持する各地物型により、tran:TrafficAreaを以下の通り定義する。道路、徒歩道又は広場の場合道路、徒歩道又は広場を構成する領域のうち、交通領域の機能を補助するために設けられた領域。LOD2、LOD3.0及びLOD3.1の場合は、道路内の島状の施設（交通島及び分離帯、路面電車停車所）を対象とする。![Image](images/_tranauxiliarytrafficarea_2_img001.png)LOD2、LOD3.0及びLOD3.1での道路のtran:AuxiliaryTrafficAreaの取得例LOD3.2 及びLOD3.3 の場合は、上記に加え、歩道部に設置された植栽を対象とする。![Image](images/_tranauxiliarytrafficarea_2_img002.png)LOD3.2及びLOD3.3での道路のtran:AuxiliaryTrafficAreaの取得例LOD3.4 に場合は、tran:function により指定されるコードリストの区分に従う。![Image](images/_tranauxiliarytrafficarea_2_img003.png)LOD3.4での道路のtran:AuxiliaryTrafficAreaの取得例鉄道の場合鉄道用地のうち、道床を除く範囲。LOD2の場合は取得しない。（tran:TrafficAreaのみを取得する。）LOD3.0の場合は取得しない。（tran:TrafficAreaのみを取得する。）LOD3.1の場合は、鉄道敷地界及び道床の外周に囲まれた範囲を取得する。高さはそれぞれの水平位置における標高とする。15㎝以上の高さの差を取得する。![Image](images/_tranauxiliarytrafficarea_2_img004.png)LOD3.1におけるtran:AuxiliaryTrafficAreaの例LOD3.2の場合は、鉄道敷地界及び道床の外周に囲まれた範囲を取得する。高さはそれぞれの水平位置における標高とする。15㎝未満の高さの差を取得する。![Image](images/_tranauxiliarytrafficarea_2_img005.png)LOD3.2におけるtran:AuxiliaryTrafficAreaの例1 つの道路オブジェクトに含まれる交通補助領域は、属性の変化が無い限り、延長方向では区分しない（例：延長方向に連続する分離帯を細分しない）。航路の場合使用しない。 | ← |
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
| tran:class     [ AuxiliaryTrafficArea ] | gml::CodeType [0..1] | 交通の分類。標準製品仕様書では使用しない。 |
| tran:function     [ AuxiliaryTrafficArea ] | gml::CodeType [0..*] | 区画線や路面標示、道路標識等により示された交通補助領域の機能。コードリスト（AuxiliaryTrafficArea_function.xml）より選択する。鉄道及び航路の場合は使用しない。 |
| tran:usage     [ AuxiliaryTrafficArea ] | gml::CodeType [0..*] | 交通補助領域の利用方法。標準製品仕様書では使用しない。 |
| tran:surfaceMaterial     [ AuxiliaryTrafficArea ] | gml::CodeType [0..1] | 表層舗装の有無及び材質。複数の表層舗装が混在している場合は、最も面積を占める舗装とする。コードリスト（AuxiliaryTrafficArea_surfaceMaterial.xml）より選択する。航路の場合は使用しない。 |
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
| tran:lod2MultiSurface             [ AuxiliaryTrafficArea ] | gml:MultiSurface [0..1] | これを使用する各地物型により、交通補助領域のtran:lod2MultiSurfaceを以下の通り定義する。 詳細な取得基準は、各地物型のLOD定義に従うこと。道路、徒歩道及び広場の場合縁石等により示される境界線に囲まれた領域のうち、通行の用に供しない道路の部分（分離帯、交通島、路面電車停車所）。高さは0とする。隣接するtran:TrafficArea又はtran:AuxiliaryTrafficAreaとの境界線の座標を一致させる。 tran: AuxiliaryTrafficAreaのtran:lod2MultiSurfaceは、同一のtran:Roadのオブジェクトに含まれる他のtran:TrafficAreaやtran:AuxiliaryTrafficAreaのtran:lod2MultiSurfaceと重なることはない。（ただし、立体的な構造をもつ道路を除く。） 車道交差部での区切りは、LOD1と同様とする。分離帯がある場合には、車道交差部の範囲を分離帯までとする。 境界線として区画線を使用する場合は、区画線の中心を境界線とする。鉄道の場合使用しない。 |
| tran:lod3MultiSurface             [ AuxiliaryTrafficArea ] | gml:MultiSurface [0..1] | これを使用する各地物型により、交通補助領域のtran:lod3MultiSurfaceを以下の通り定義する。 詳細な取得基準は、各地物型のLOD定義に従うこと。道路、徒歩道及び広場の場合縁石等により示される境界線に囲まれた領域のうち、通行の用に供しない道路の部分。LOD3.0の場合、横断方向に連続する交通領域の高さは一律とし、車道の標高とする。LOD3.1～LOD3.4では、各水平位置における標高とする。隣接するtran:TrafficArea又はtran:AuxiliaryTrafficAreaとの境界線の座標を一致させる。tran: AuxiliaryTrafficAreaのtran:lod3MultiSurfaceは、同一のtran: AuxiliaryTrafficAreaのtran:lod2MultiSurfaceは、同一のtran:Roadのオブジェクトに含まれる他のtran:TrafficAreaやtran:AuxiliaryTrafficAreaのtran:lod2MultiSurfaceと重なることはない。（ただし、立体的な構造をもつ道路を除く。）車道交差部での区切りは、LOD1と同様とする。分離帯がある場合には、車道交差部の範囲を分離帯までとする。境界線として区画線を使用する場合は、区画線の中心を境界線とする。鉄道の場合鉄道敷地界及び道床の外周に囲まれた範囲。LOD3.0の場合は取得しない。LOD3.1の場合は、15㎝以上の高さの差を取得する。LOD3.2の場合は、15㎝未満の高さの差を取得する。 |

