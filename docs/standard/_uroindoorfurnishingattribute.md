# 4.2.3.5.31 uro:IndoorFurnishingAttribute

## 4.2.3.5.31 uro:IndoorFurnishingAttribute

**表 4-88**

| 型の定義 | 設備に追加するナビゲーション用の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IndoorAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:source     [ IndoorAttribute ] | gml::CodeType [0..1] | 原典資料。コードリスト（Common_indoorSource.xml）から選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:floorId     [ IndoorFurnishingAttribute ] | xs::string [0..1] | 固定設置物が紐づけられている階層の固有ID。 |

