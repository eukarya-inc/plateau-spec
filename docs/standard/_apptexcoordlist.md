# 4.22.3.1.3 app:TexCoordList

## 4.22.3.1.3 app:TexCoordList

**表 4-835**

| 型の定義 | テクスチャの座標とテクスチャを貼り付ける幾何オブジェクトへの参照の対。 | ← |
| --- | --- | --- |
| 上位の型 | app:_TextureParameterization | ← |
| ステレオタイプ | << Type >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| app:textureCoordinates     [ TexCoordList ] | gml::doubleList [1..*] | テクスチャの座標。UV座標（画像の横方向をU軸、縦方向をV軸とし、座標値を0から1の小数値を用いて記述する座標）により記述する。 |
| app:ring     [ TexCoordList ] | xs::anyURI [1..*] | 幾何オブジェクトへの参照。相対パスにより記述する。参照する幾何オブジェクトは、gml:LinearRingとする。 |

