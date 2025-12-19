# 4.26.3.6.5 uro:LongevityMeasures

## 4.26.3.6.5 uro:LongevityMeasures

**表 4-896**

| 型の定義 | 長寿命化対策に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:fiscalYearForCountermeasures     [ LongevityMeasures ] | xs::gYear [1..1] | 対策実施の年度（将来の予定を含む）。 |
| uro:countermeasuresCost     [ LongevityMeasures ] | CountermeasuresCost [0..1] | 対策費用（将来の予定を含む）。 |
| uro:description     [ LongevityMeasures ] | xs::string [0..1] | 対策の内容（将来の予定を含む）。 |

