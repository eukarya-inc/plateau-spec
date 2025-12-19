# 4.13.3.1.19 uro:ConstructionEvent

## 4.13.3.1.19 uro:ConstructionEvent

**表 4-636**

| 型の定義 | 構造物の設計から施工、維持管理にいたるイベント。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:event     [ ConstructionEvent ] | gml::CodeType [1..1] | イベントの種類。コードリストから選択する。この属性を使用する場合は、コードリスト（ConstructionEvent_event.xml）を作成すること。 |
| uro:dateOfEvent     [ ConstructionEvent ] | xs::date [1..1] | イベントが生じた日付。 |
| uro:description     [ ConstructionEvent ] | xs::string [0..1] | イベントの説明。 |

