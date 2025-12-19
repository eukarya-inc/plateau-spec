# 4.2.3.5.36 uro:IndoorUserDefinedAttribute

## 4.2.3.5.36 uro:IndoorUserDefinedAttribute

**表 4-93**

| 型の定義 | 任意に追加するナビゲーション用の属性。 bldg:InteriorWallSurface、bldg:CeilingSurface、bldg:FloorSurfaceに使用する。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IndoorAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:source     [ IndoorAttribute ] | gml::CodeType [0..1] | 原典資料。コードリスト（Common_indoorSource.xml）から選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:name     [ IndoorUserDefinedAttribute ] | xs::string [0..1] | フィールド名。 |
| uro:nominalValue     [ IndoorUserDefinedAttribute ] | UserDefinedValue [0..1] | フィールド名に対応する属性値。 |
| uro:description     [ IndoorUserDefinedAttribute ] | xs::string [0..1] | 説明情報。 |
| uro:unit     [ IndoorUserDefinedAttribute ] | xs::string [0..1] | 単位。 |

