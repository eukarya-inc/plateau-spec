# 4.2.3.5.14 uro:IfcProject

## 4.2.3.5.14 uro:IfcProject

**表 4-71**

| 型の定義 | IFCで記述されたプロジェクトに適用される属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcObject | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:globalId     [ IfcRoot ] | xs::string [0..1] | そのオブジェクトが単一に識別できる唯一な識別子。22桁の文字列により表現する。 |
| uro:name     [ IfcRoot ] | xs::string [0..1] | オブジェクトの名称。 |
| uro:description     [ IfcRoot ] | xs::string [0..1] | オブジェクトの文字情報による追加説明。 |
| uro:objectType     [ IfcObject ] | xs::string [0..1] | オブジェクトの特定の型を示す。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:longName     [ IfcProject ] | xs::string [0..1] | 人が認識可能な名称。 |
| uro:phase     [ IfcProject ] | xs::string [0..1] | プロジェクトの状態。計画、完成など。 |
| uro:representationContexts     [ IfcProject ] | IfcGeometricRepresentationContext [0..1] | プロジェクト内のIfcProduct オブジェクトの形状表現に適用されるコンテキスト。 |
| uro:unitsInContext     [ IfcProject ] | IfcUnit [0..*] | 使用される単位系情報。 |

