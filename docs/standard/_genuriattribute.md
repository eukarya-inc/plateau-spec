# 4.21.3.1.5 gen:uriAttribute

## 4.21.3.1.5 gen:uriAttribute

**表 4-829**

| 型の定義 | URI型の汎用属性。CityGML又はi-URに定義されていないURI型の属性を追加したい場合に使用する。 | ← |
| --- | --- | --- |
| 上位の型 | gen:_genericAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gen:name     [ _genericAttribute ] | xs::string [1..1] | 汎用属性の名称。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gen:value     [ uriAttribute ] | xs::anyURI [1..1] | 汎用属性の値。 |

