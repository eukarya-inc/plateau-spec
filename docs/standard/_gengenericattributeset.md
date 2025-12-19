# 4.21.3.1.7 gen:genericAttributeSet

## 4.21.3.1.7 gen:genericAttributeSet

**表 4-831**

| 型の定義 | 汎用属性のグループ。CityGML又はi-URに定義されていない属性をグループ化して追加したい場合に使用する。 | ← |
| --- | --- | --- |
| 上位の型 | gen:_genericAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gen:name     [ _genericAttribute ] | xs::string [1..1] | 汎用属性の名称。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gen:codeSpace     [ genericAttributeSet ] | gml::CodeType [0..1] | 汎用属性のグループを維持管理する機関のURI。文字列とする。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| gen:_genericAttribute             [ genericAttributeSet ] | gen:_genericAttribute [1..*] | グループに含みたい汎用属性。 |

