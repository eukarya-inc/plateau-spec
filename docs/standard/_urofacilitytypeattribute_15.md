# 4.26.3.1.2 uro:FacilityTypeAttribute

## 4.26.3.1.2 uro:FacilityTypeAttribute

**表 4-872**

| 型の定義 | 管理施設の用途に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:class     [ FacilityTypeAttribute ] | gml::CodeType [0..1] | 管理施設の区分。コードリスト（FacilityTypeAttribute_class.xml）から選択する。 |
| uro:function     [ FacilityTypeAttribute ] | gml::CodeType [0..*] | 管理施設の用途。コードリスト（FacilityTypeAttribute_function.xml）から選択する。 |

