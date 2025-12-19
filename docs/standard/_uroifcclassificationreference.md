# 4.2.3.5.5 uro:IfcClassificationReference

## 4.2.3.5.5 uro:IfcClassificationReference

**表 4-62**

| 型の定義 | IFCで記述された分類に関する属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:location     [ IfcClassificationReference ] | xs::anyURI [0..1] | 分類の外部ソース情報。 |
| uro:itemReference     [ IfcClassificationReference ] | gml::CodeType [0..1] | 分類コード。コードリスト（IfcClassificationReference_itemReference.xml）から選択する。この属性を使用する場合は、コードリストを作成する。 |
| uro:name     [ IfcClassificationReference ] | xs::string [0..1] | コードに対応するラベル。 |
| uro:referencedSource     [ IfcClassificationReference ] | IfcClassification [0..1] | 分類の諸元。 |

