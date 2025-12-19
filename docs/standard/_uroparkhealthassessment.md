# 4.26.3.6.4 uro:ParkHealthAssessment

## 4.26.3.6.4 uro:ParkHealthAssessment

**表 4-895**

| 型の定義 | 健全度調査結果に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:assessmentFiscalYear     [ ParkHealthAssessment ] | xs::gYear [1..1] | 健全度調査の実施年度。 |
| uro:deteriorationStatus     [ ParkHealthAssessment ] | xs::string [0..1] | 健全度調査結果で確認した劣化状況。 |
| uro:condition     [ ParkHealthAssessment ] | gml::CodeType [0..1] | 健全度調査結果で確認した健全度。コードリスト（ Common_parkHealthAssessmentCondition.xml）より選択する。 |
| uro:urgency     [ ParkHealthAssessment ] | gml::CodeType [0..1] | 健全度調査結果で確認した対策の緊急度。コードリスト（ Common_parkHealthAssessmentUrgency.xml）より選択する。 |

