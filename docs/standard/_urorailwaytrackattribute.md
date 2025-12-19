# 4.4.3.2.5 uro:RailwayTrackAttribute

## 4.4.3.2.5 uro:RailwayTrackAttribute

**表 4-200**

| 型の定義 | 軌道の情報。軌道ごとに、LOD2以上で作成する。LOD3では軌道中心線の平面線形要素ごとに作成する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:routeName     [ RailwayTrackAttribute ] | xs::string [0..1] | 鉄道路線の名称。 |
| uro:directionType     [ RailwayTrackAttribute ] | gml::CodeType [0..1] | 上り下りの別。コードリスト（RailwayTrackAttribute_directionType.xml） から選択する。 |
| uro:trackType     [ RailwayTrackAttribute ] | gml::CodeType [0..1] | 軌道の種類。コードリスト（ RailwayTrackAttribute_trackType.xml）から選択する。 |
| uro:startPost     [ RailwayTrackAttribute ] | xs::string [0..1] | 開始キロ程。 |
| uro:endPost     [ RailwayTrackAttribute ] | xs::string [0..1] | 終了キロ程。 |
| uro:alignmentType     [ RailwayTrackAttribute ] | gml::CodeType [0..1] | 軌道中心線の線形要素の種別。コードリスト（RailwayTrackAttribute_alignmentType.xml）から選択する。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:controlPoint             [ RailwayTrackAttribute ] | uro:ControlPoint [0..*] | 軌道中心線上の線形変化点。 |
| uro:lod2Network             [ RailwayTrackAttribute ] | gml:GeometricComplex [0..1] | 軌道中心線。高さは0とする。 |
| uro:lod3Network             [ RailwayTrackAttribute ] | gml:GeometricComplex [0..1] | 軌道中心線。高さは軌道中心線には、下り本線上の高さ（標高）を与える。 |

