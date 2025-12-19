# 4.22.3.1.1 app:Appearance

## 4.22.3.1.1 app:Appearance

**表 4-833**

| 型の定義 | 地物の外観。 地物の幾何オブジェクトとなる面に色又はテクスチャを指定する。![Image](images/_appappearance_img001.png)Apperanceを使用してbldg:Buildingの外観を設定した例　　　　（左：色の指定　右：テクスチャの指定） | ← |
| --- | --- | --- |
| 上位の型 | gml:_Feature | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| app:theme     [ Appearance ] | xs::string [0..1] | 主題。固定値とし、「rgbTexture」とする。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| app:surfaceDataMember             [ Appearance ] | app:_SurfaceData [0..*] | 使用するテクスチャ及びパラメータ又は色及びパラメータへの参照。 |

