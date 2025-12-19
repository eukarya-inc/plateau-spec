# 4.2.3.5.13 uro:IfcOpeningElement

## 4.2.3.5.13 uro:IfcOpeningElement

**表 4-70**

| 型の定義 | IFCで記述された、床や壁に設けられた開口部の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcElement | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:globalId     [ IfcRoot ] | xs::string [0..1] | そのオブジェクトが単一に識別できる唯一な識別子。22桁の文字列により表現する。 |
| uro:name     [ IfcRoot ] | xs::string [0..1] | オブジェクトの名称。 |
| uro:description     [ IfcRoot ] | xs::string [0..1] | オブジェクトの文字情報による追加説明。 |
| uro:objectType     [ IfcObject ] | xs::string [0..1] | オブジェクトの特定の型を示す。 |
| uro:tag     [ IfcElement ] | xs::string [0..1] | オブジェクトのシリアルナンバー、ポジションナンバーなどの識別番号。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:nominalArea     [ IfcOpeningElement ] | gml::MeasureType [0..1] | 全体の面積。単位はm2。 |
| uro:nominalVolume     [ IfcOpeningElement ] | gml::MeasureType [0..1] | 全体の体積。単位はm3。 |

