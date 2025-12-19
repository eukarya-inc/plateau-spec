# 4.2.3.5.33 uro:IndoorSpaceAttribute

## 4.2.3.5.33 uro:IndoorSpaceAttribute

**表 4-90**

| 型の定義 | 物理的な空間に追加するナビゲーション用の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IndoorAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:source     [ IndoorAttribute ] | gml::CodeType [0..1] | 原典資料。コードリスト（Common_indoorSource.xml）から選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:floorId     [ IndoorSpaceAttribute ] | xs::string [0..1] | 物理的な空間が紐づけられている階層の固有ID。 |
| uro:isRestricted     [ IndoorSpaceAttribute ] | xs::boolean [0..1] | 業務用エリアなど一般の人の進入制限の有無。1：進入制限あり0：進入制限なし |
| uro:suite     [ IndoorSpaceAttribute ] | xs::string [0..1] | 地図表示用の注記ラベル。 |
| uro:isPublic     [ IndoorSpaceAttribute ] | xs::boolean [0..1] | 地図情報としての公開可否。1：公開可0：公開不可 |
| uro:tollType     [ IndoorSpaceAttribute ] | gml::CodeType [0..1] | 有料施設の区分。コードリスト（IndoorSpaceAttribute_tollType.xml）より選択する。 |

