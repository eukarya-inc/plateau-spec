# 4.2.3.5.18 uro:IfcPsetOpeningElementCommon

## 4.2.3.5.18 uro:IfcPsetOpeningElementCommon

**表 4-75**

| 型の定義 | Ifcで記述された開口部に共通の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:reference     [ IfcPsetOpeningElementCommon ] | xs::string [0..1] | 参照用のID。 |
| uro:purpose     [ IfcPsetOpeningElementCommon ] | xs::string [0..1] | この開口部の目的。（例：換気、アクセス） |
| uro:fireExit     [ IfcPsetOpeningElementCommon ] | xs::boolean [0..1] | この開口部が火災時の非常用出口として機能するよう設計されているか。1：設計されている0：されていない |
| uro:protectedOpening     [ IfcPsetOpeningElementCommon ] | xs::boolean [0..1] | この開口部が、防火上の観点で保護されているとみなせるかどうか。みなされる場合、該当する法令のもの確保された開口部としてカウントする。1：みなされる0：みなされない |
| uro:parallelJambs     [ IfcPsetOpeningElementCommon ] | xs::boolean [0..1] | 湾曲した開口部のわき柱が平行になるように意図されているか否か。1：意図されている0：されていない |

