# 4.10.3.7.9 urf:ParkAttribute

## 4.10.3.7.9 urf:ParkAttribute

**表 4-388**

| 型の定義 | 都市計画法第11条第1項第2号に定める公園について定めるべき事項。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:parkTypeNumber     [ ParkAttribute ] | gml::CodeType [0..1] | 都市局長通達(昭和44年建設省都計発第102号)Ⅳ(4)②にて定められる一連番号。 |
| urf:parkSizeNumber     [ ParkAttribute ] | gml::CodeType [0..1] | 都市局長通達(昭和44年建設省都計発第102号)Ⅳ(4)②にて定められる規模。コードリスト（ ParkAttribute_parkSizeNumber.xml）より選択する。 |
| urf:parkSerialNumber     [ ParkAttribute ] | xs::string [0..1] | 都市局長通達(昭和44年建設省都計発第102号)Ⅳ(4)②にて定められる区分。コードリスト（ParkAttribute_parkTypeNumber.xml）より選択する。 |

