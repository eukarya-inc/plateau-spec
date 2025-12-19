# 4.26.3.6.3 uro:RepairsBeforeParkHealthAssessment

## 4.26.3.6.3 uro:RepairsBeforeParkHealthAssessment

**表 4-894**

| 型の定義 | 健全度調査以前に実施した補修に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:repair     [ RepairsBeforeParkHealthAssessment ] | gml::CodeType [1..1] | 健全度調査以前に実施した補修の有無。コードリスト（Common_parkRepair.xml）から選択する。 |
| uro:repairFiscalYear     [ RepairsBeforeParkHealthAssessment ] | xs::gYear [0..1] | 補修の実施年度。 |

