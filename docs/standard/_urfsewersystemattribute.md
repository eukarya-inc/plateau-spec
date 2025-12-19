# 4.10.3.7.13 urf:SewerSystemAttribute

## 4.10.3.7.13 urf:SewerSystemAttribute

**表 4-392**

| 型の定義 | 都市計画法第11条第1項第3号に定める下水道について定めるべき事項。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:startLocation     [ SewerSystemAttribute ] | xs::string [0..1] | 都市計画法第11条第2項で定める位置`［起点の町丁目又は字］`。 |
| urf:endLocation     [ SewerSystemAttribute ] | xs::string [0..1] | 都市計画法第11条第2項で定める位置`［終点の町丁目又は字］`。 |
| urf:systemType     [ SewerSystemAttribute ] | gml::CodeType [0..1] | 下水道法施行規則第19条第1項第3号で定められる種別。コードリスト（SewerSystemAttribute_systemType.xml）より選択する。 |
| urf:drainageArea     [ SewerSystemAttribute ] | gml:StringOrRefType [0..1] | 都市計画法施行令第6条第1項第6号に定められた排水区域。 |

