# 4.23.3.1.1 uro:AbstractPointCloud

## 4.23.3.1.1 uro:AbstractPointCloud

**表 4-840**

| 型の定義 | 地物の三次元点群データを表す抽象型。 | ← |
| --- | --- | --- |
| 上位の型 | gml:_Feature | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |

