# 4.12.3.2.2 uro:TunnelStructureAttribute

## 4.12.3.2.2 uro:TunnelStructureAttribute

**表 4-580**

| 型の定義 | トンネルの構造に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:tunnelType     [ TunnelStructureAttribute ] | gml::CodeType [0..1] | トンネルの区分。コードリスト（TunnelStructureAttribute_tunnelType.xml）より選択する。 |
| uro:tunnelSubtype     [ TunnelStructureAttribute ] | gml::CodeType [0..1] | トンネルの詳細な区分。コードリスト（TunnelStructureAttribute_tunnelSubType.xml）より選択する。 |
| uro:mouthType     [ TunnelStructureAttribute ] | gml::CodeType [0..*] | 坑口の種類。コードリスト（TunnelStructureAttribute_mouthType.xml）より選択する。 |
| uro:length     [ TunnelStructureAttribute ] | gml::LengthType [0..1] | トンネルの長さ。単位はmとする。 |
| uro:width     [ TunnelStructureAttribute ] | gml::LengthType [0..1] | トンネルの幅員。単位はmとする。 |
| uro:area     [ TunnelStructureAttribute ] | gml::MeasureType [0..1] | トンネルの面積。単位はm2とする。 |
| uro:innerHeight     [ TunnelStructureAttribute ] | gml::LengthType [0..1] | トンネルの内空高。単位はmとする。 |
| uro:effectiveHeight     [ TunnelStructureAttribute ] | gml::LengthType [0..1] | トンネルの有効高。単位はmとする。 |
| uro:slopeType     [ TunnelStructureAttribute ] | gml::CodeType [0..1] | トンネルが地下横断歩道である場合の、昇降形式。コードリスト（ConstructionStructureAttribute_slopeType.xml）より選択する。 |

