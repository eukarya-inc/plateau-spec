# 4.21.3.1.4 gen:dateAttribute

## 4.21.3.1.4 gen:dateAttribute

**表 4-828**

| 型の定義 | 日付型の汎用属性。CityGML又はi-URに定義されていない日付型の属性を追加したい場合に使用する。 | ← |
| --- | --- | --- |
| 上位の型 | gen:_genericAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gen:name     [ _genericAttribute ] | xs::string [1..1] | 汎用属性の名称。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gen:value     [ dateAttribute ] | xs::date [1..1] | 汎用属性の値。 |

