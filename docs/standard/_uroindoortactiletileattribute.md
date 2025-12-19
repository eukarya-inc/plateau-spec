# 4.2.3.5.34 uro:IndoorTactileTileAttribute

## 4.2.3.5.34 uro:IndoorTactileTileAttribute

**表 4-91**

| 型の定義 | 視覚障碍者用誘導ブロックに追加するナビゲーション用の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IndoorAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:source     [ IndoorAttribute ] | gml::CodeType [0..1] | 原典資料。コードリスト（Common_indoorSource.xml）から選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:startNode     [ IndoorTactileTileAttribute ] | xs::string [0..1] | 視覚障碍者誘導用ブロック等の開始位置の固有ID。接続するブロック（点）がある場合に入力する。 |
| uro:endNode     [ IndoorTactileTileAttribute ] | xs::string [0..1] | 視覚障碍者誘導用ブロック等の終了位置の固有ID。接続するブロック（点）がある場合に入力する。 |
| uro:category     [ IndoorTactileTileAttribute ] | gml::CodeType [0..1] | 視覚障碍者誘導用ブロック等の種類。コードリスト（IndoorTactileTileAttribute_category.xml）より選択する。 |
| uro:roof     [ IndoorTactileTileAttribute ] | gml::CodeType [0..1] | 屋根の有無。コードリスト（IndoorTactileTileAttribute_roof.xml）より選択する。 |
| uro:floorId     [ IndoorTactileTileAttribute ] | xs::string [0..1] | 誘導ブロックが紐づけられている階層の固有ID。 |

