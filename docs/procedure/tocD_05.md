# D.5 tran:AuxiliarlyTrafficArea

## D.5 tran:AuxiliarlyTrafficArea

tran:AuxiliaryTrafficAreaは、植樹帯や分離帯のように交通領域の機能を補助するために設けられた領域（交通補助領域）を示す。通常は、車両や人の通行には用いられない領域である。CityGMLでは、地物型として植樹帯や分離帯は用意されておらず、属性tran:functionを用いて区分する。3D都市モデルでは、LODに応じて交通補助領域を区分するために、LODごとに属性tran:functionの定義域を定めている。なお、この定義域は、[道路基盤地図情報（整備促進版）製品仕様書（案）](https://www.nilim.go.jp/lab/bcg/siryou/tnn/tnn0848pdf/ks084811.pdf)において道路面を構成する地物の区分との整合を考慮している。

**表 D-20 — D-10: 交通補助領域の区分方法**

| ID | `/req/tran/10` |
| --- | --- |
| 主題 | 妥当な交通（道路）データセットの要件 |
| 分類 | 要件分類D-1: 妥当な交通（道路）オブジェクト |
| 条文 | LODに応じた交通補助領域の区分には、属性tran:functionを用いる。 |

LOD2、LOD3.0及びLOD3.1では、交通補助領域は、分離帯や交通島のように道路上に設けられた構造物を区分する（表D-21）。

**表 D-21 — LOD2、LOD3.0及びLOD3.1での交通補助領域の区分**

| ファイル名 | AuxiliaryTrafficArea_function.xml | ← |
| --- | --- | --- |
| ファイルURL | [https://www.geospatial.jp/iur/codelists/3.2/AuxiliaryTrafficArea_function.xml](https://www.geospatial.jp/iur/codelists/3.2/AuxiliaryTrafficArea_function.xml) | ← |
| 大分類 | ← | 定義 |
| コード | 説明 | ↑ |
| 3000 | 島 | 車両の走行を制御し、安全な交通を確保するために設置される分離帯及び交通島。路面電車停車所が設けられた島を含む。 |

LOD3.2及びLOD3.3では、分離帯や交通島のように道路上に設けられた構造物に加え、歩道上の植栽を区分する（表D-22）。

**表 D-22 — LOD3.2及びLOD3.3での交通補助領域の区分**

| ファイル名 | AuxiliaryTrafficArea_function.xml | ← |
| --- | --- | --- |
| ファイルURL | [https://www.geospatial.jp/iur/codelists/3.2/AuxiliaryTrafficArea_function.xml](https://www.geospatial.jp/iur/codelists/3.2/AuxiliaryTrafficArea_function.xml) | ← |
| 大分類 | ← | 定義 |
| コード | 説明 | ↑ |
| 3000 | 島 | 車両の走行を制御し、安全な交通を確保するために設置される分離帯及び交通島。路面電車停車所が設けられた島を含む。 |
| 5000 | 植栽 | 植樹帯及び植樹ます。 |

LOD3.4では、より詳細に道路内を区分できる（表D-23）。この区分は、[道路基盤地図情報（整備促進版）製品仕様書（案）](https://www.nilim.go.jp/lab/bcg/siryou/tnn/tnn0848pdf/ks084811.pdf)において道路面を構成する地物の区分と整合する。
 LOD3.4の区分をどこまで詳細に行うかは、ユースケースに応じて決定してよい。

**表 D-23 — LOD3.4での交通補助領域の区分**

| ファイル名 | AuxiliaryTrafficArea_function.xml | ← | ← | ← |
| --- | --- | --- | --- | --- |
| ファイルURL | [https://www.geospatial.jp/iur/codelists/3.2/AuxiliaryTrafficArea_function.xml](https://www.geospatial.jp/iur/codelists/3.2/AuxiliaryTrafficArea_function.xml) | ← | ← | ← |
| 大分類コード | 大分類定義 | 小分類コード | 小分類定義 | 定義 |
| 1000 | 車道部 |  | ← | 主として自動車が利用する道路の部分のうち、自動車の通行の用に供されない（物理的に通行が可能であっても、道路設計上、車両が通行することが想定されていない）を部分。非常駐車帯や中央帯の区分が不要な場合には、通行が想定されていない範囲を全て車道部として取得する。 |
| ↑ | ↑ | 1060 | 非常駐車帯 | 左側路肩に設けられる、故障車等が本線車線から退避し一時的に駐車するための道路の部分。 |
| ↑ | ↑ | 1080 | 中央帯 | 車線を往復の方向別に区分するための道路の部分。 |
| ↑ | ↑ | 1090 | 側帯 | 運転者の視線を誘導し、側方余裕をもたせるため、路肩及び中央帯にも受けられる道路の部分。 |
| ↑ | ↑ | 1100 | 路肩 | 道路の主要構造を保護し、車道の機能を確保するため、車道部や歩道部に連続して設置される道路の部分。 |
| ↑ | ↑ | 1110 | 停車帯 | 車両が停車するために設けられる道路の部分。 |
| ↑ | ↑ | 1120 | 乗合自動車停車所 | バス乗客の乗降のため、本線車線から分離しても受けられる道路の部分。 |
| 3000 | 島 |  | ← | 交通島、分離帯の区分が不要な場合は、島として取得する。 |
| ↑ | ↑ | 3010 | 交通島 | 車両の走行を制御し歩行者を保護するために設置される島状の道路の部分。 |
| ↑ | ↑ | 3020 | 分離帯 | 同方向又は対方向の交通流を分離するために設置される島状の道路の部分。 |
| 4000 | 路面電車停車所 |  | ← | 路面電車の乗降、待合のための停留場として利用される島状の部分。 |
| 5000 | 植栽 |  | ← | 植樹帯、植樹ますの区分をしない場合には全て植栽として取得する。 |
| ↑ | ↑ | 5010 | 植樹帯 | 植栽のために工作物により区切られる道路の帯状の部分。 |
| ↑ | ↑ | 5020 | 植樹ます | 歩道上に設置される植栽のためのます。 |
| 6000 | 自転車駐車場 |  | ← | 自転車駐車場のうち、駐車区画の部分。 |
| 7000 | 自動車駐車場 |  | ← | 自動車駐車場のうち、駐車区画の部分。 |

## D.5.1 交通補助領域の空間属性

tran:AuxiliaryTrafficAreaも、道路面を構成する要素であるため、tran:TrafficAreaと同様の要件を満たす必要がある。

**表 D-24 — D-11: tran:AuxiliaryTrafficAreaの重なり禁止条件**

| ID | `/req/tran/11` |
| --- | --- |
| 主題 | 妥当な交通（道路）データセットの要件 |
| 分類 | 要件分類D-1: 妥当な交通（道路）オブジェクト |
| 条文 | 同一のtran:Roadに含まれるtran:AuxiliaryTrafficAreaの空間属性は、他のtran:TrafficAreaやtran:AuxiliaryTrafficAreaの空間属性と重なってはならない。また、同一のtran:Roadに含まれるtran:AuxiliaryTrafficAreaの空間属性は、隣接するtran:TrafficAreaやtran:AuxiliaryTrafficAreaの空間属性と、境界線を共有しなければならない。 |

## D.5.2 交通補助領域の主題属性

交通補助領域の主題属性であるtran:function及びtran:surfaceMaterialについては、交通領域と同様である。

