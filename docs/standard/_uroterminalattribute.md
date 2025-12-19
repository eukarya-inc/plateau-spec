# 4.6.3.2.6 uro:TerminalAttribute

## 4.6.3.2.6 uro:TerminalAttribute

**表 4-273**

| 型の定義 | 自動車ターミナルに関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:SquareUrbanPlanAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:prefecture     [ SquareUrbanPlanAttribute ] | gml::CodeType [0..1] | 広場が位置する都道府県。JIS X0401に定義される2桁の半角数字。コードリスト（Common_localPublicAuthorities.xml）より選択する。 |
| uro:city     [ SquareUrbanPlanAttribute ] | gml::CodeType [0..1] | 広場が位置する市区町村。JIS X0401に定義される2桁の半角数字とJIS X0402に定義される3桁の半角数字とを組み合わせた5桁の半角数字。政令市の場合は、区の市区町村コードとする。コードリスト（Common_localPublicAuthorities.xml）より選択する。 |
| uro:urbanPlanningAreaName     [ SquareUrbanPlanAttribute ] | xs::string [0..1] | 都市計画区域の名称。 |
| uro:enforcer     [ SquareUrbanPlanAttribute ] | xs::string [0..*] | 施行者名。 |
| uro:dateOfDecision     [ SquareUrbanPlanAttribute ] | xs::date [0..1] | 都市計画の決定日。 |
| uro:dateOfRevision     [ SquareUrbanPlanAttribute ] | xs::date [0..*] | 都市計画の変更年月日。 |
| uro:areaPlanned     [ SquareUrbanPlanAttribute ] | gml::MeasureType [0..1] | 計画面積。交通広場の場合は、単位はm2とする。 |
| uro:areaInService     [ SquareUrbanPlanAttribute ] | gml::MeasureType [0..1] | 供用面積。交通広場の場合は、単位はm2とする。 |
| uro:remarks     [ SquareUrbanPlanAttribute ] | xs::string [0..1] | 都市施設の摘要。 |
| uro:status     [ SquareUrbanPlanAttribute ] | gml::CodeType [0..1] | 事業の進捗状況。コードリスト（Common_status.xml）より選択する。 |
| uro:areaImproved     [ SquareUrbanPlanAttribute ] | gml::MeasureType [0..1] | 改良済（用地が計画のとおり確保されており、供用している）の面積又は延長。交通広場の場合は面積で記述する。単位はm2とする。 |
| uro:areaCompleted     [ SquareUrbanPlanAttribute ] | gml::MeasureType [0..1] | 概成済（改良済み以外の区間のうち、都市計画施設と同程度の機能をしている）の面積又は延長。交通広場の場合は面積で記述する。単位はm2とする。 |
| uro:projectStartDate     [ SquareUrbanPlanAttribute ] | xs::date [0..1] | 事業開始年月日。事業に着手していないもの、計画決定時に完成しているものは記入しない。 |
| uro:projectEndDate     [ SquareUrbanPlanAttribute ] | xs::date [0..1] | 事業完了年月日。事業が完了していないもの、事業に着手していないもの、計画決定時に完成しているものは記入しない。 |
| uro:isCompleted     [ SquareUrbanPlanAttribute ] | xs::boolean [0..1] | 計画決定時に完成している場合に1とする。 |
| uro:isAuthorized     [ SquareUrbanPlanAttribute ] | xs::boolean [0..1] | 認可を受けている場合に1とする。 |
| uro:purpose     [ SquareUrbanPlanAttribute ] | xs::string [0..1] | 都市計画の変更を行った場合に、その目的を記述する。 |
| uro:note     [ SquareUrbanPlanAttribute ] | xs::string [0..1] | その他特筆事項。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:terminalType     [ TerminalAttribute ] | gml::CodeType [0..1] | 自動車ターミナルの種類。コードリスト（TerminalAttribute_terminalType.xml）から選択する。 |
| uro:structure     [ TerminalAttribute ] | xs::string [0..1] | 自動車ターミナルの構造。 |
| uro:numberOfBerthsPlanned     [ TerminalAttribute ] | xs::integer [0..1] | 計画バース数。 |
| uro:numberOfBerthsInService     [ TerminalAttribute ] | xs::integer [0..1] | 供用バース数。 |
| uro:userType     [ TerminalAttribute ] | gml::CodeType [0..1] | 一般、専用の別。コードリスト（TerminalAttribute_userType.xml）から選択する。 |

