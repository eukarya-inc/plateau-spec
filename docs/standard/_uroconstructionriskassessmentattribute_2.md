# 4.12.3.2.7 uro:ConstructionRiskAssessmentAttribute

## 4.12.3.2.7 uro:ConstructionRiskAssessmentAttribute

**表 4-585**

| 型の定義 | 構造物の損傷及び対応状況に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:surveyYear     [ ConstructionRiskAssessmentAttribute ] | xs::gYear [0..1] | 点検が実施された年度。 |
| uro:riskType     [ ConstructionRiskAssessmentAttribute ] | gml::CodeType [1..1] | 判定区分。コードリスト（ConstructionRiskAssessmentAttribute_riskType.xml）より選択する。 |
| uro:status     [ ConstructionRiskAssessmentAttribute ] | gml::CodeType [0..1] | 対応状況。コードリスト（ ConstructionRiskAssessmentAttribute_status.xml）より選択する。 |
| uro:referenceDate     [ ConstructionRiskAssessmentAttribute ] | xs::date [1..1] | 判定区分や措置状況の情報が記載された損傷マップの更新時点。 |

