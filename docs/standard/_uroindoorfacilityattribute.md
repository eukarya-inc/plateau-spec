# 4.2.3.5.30 uro:IndoorFacilityAttribute

## 4.2.3.5.30 uro:IndoorFacilityAttribute

**表 4-87**

| 型の定義 | 施設に追加するナビゲーション用の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IndoorAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:source     [ IndoorAttribute ] | gml::CodeType [0..1] | 原典資料。コードリスト（Common_indoorSource.xml）から選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:weekdayHours     [ IndoorFacilityAttribute ] | xs::string [0..1] | 施設の営業時間（平日）。平日でも曜日により営業時間が異なる場合は、各曜日の営業時間を記載。 |
| uro:weekendHours     [ IndoorFacilityAttribute ] | xs::string [0..1] | 施設の営業時間（土日祝祭日）。土日祝祭日により営業時間が異なる場合は、それぞれの営業時間を記載。 |
| uro:phone     [ IndoorFacilityAttribute ] | xs::string [0..1] | 施設の電話番号。 |
| uro:website     [ IndoorFacilityAttribute ] | xs::string [0..1] | 施設のウェブサイトアドレス（URL）。 |

