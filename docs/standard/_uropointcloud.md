# 4.23.3.1.2 uro:PointCloud

## 4.23.3.1.2 uro:PointCloud

**表 4-841**

| 型の定義 | 地物の三次元点群データ。三次元点群データとは、地形、地物等を表す三次元座標を持つ多数の点データ及びその内容を表す属性データを、計算処理が可能な形態で表現したものをいう。 ［出典：[作業規定の準則](https://www.gsi.go.jp/gijyutukanri/gijyutukanri41018.html)］ | ← |
| --- | --- | --- |
| 上位の型 | uro:AbstractPointCloud | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:mimeType     [ PointCloud ] | gml::CodeType [0..1] | 三次元点群データの外部ファイルフォーマット。コードリスト（ PointCloud_mimeType.xml）より選択する。三次元点群データの外部ファイルフォーマットは、LAS又はLAZとする。 |
| uro:pointFile     [ PointCloud ] | xs::anyURI [0..1] | 三次元点群データの外部ファイルへの参照。相対パスにより記述する。 |
| uro:pointFileSrsName     [ PointCloud ] | xs::string [0..1] | 三次元点群データの外部ファイルの座標参照系。 EPSG:[EPSGコード]  [EPSGコード]は、EPSGにより指定された半角数字の組合せによる識別子とする。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:pointCloudSpecAttribute             [ PointCloud ] | uro:PointCloudSpecAttribute [0..1] | 三次元点群データの仕様。 |
| uro:points             [ PointCloud ] | gml:MultiPoint [0..1] | 三次元点群データの幾何形状。 インラインにより三次元点群データの幾何形状を記述する場合に使用する。 |

