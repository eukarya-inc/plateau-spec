# 4.26.3.6.6 uro:CountermeasuresCost

## 4.26.3.6.6 uro:CountermeasuresCost

**表 4-897**

| 型の定義 | 長寿命化対策の費用に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:cost     [ CountermeasuresCost ] | xs::integer [0..1] | 対策の費用（将来の予定を含む）。 |
| uro:costUnit     [ CountermeasuresCost ] | xs::string [1..1] | 費用の単位。通常は“千円”。 |

