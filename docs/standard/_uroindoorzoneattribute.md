# 4.2.3.5.35 uro:IndoorZoneAttribute

## 4.2.3.5.35 uro:IndoorZoneAttribute

**表 4-92**

| 型の定義 | 任意の空間に追加するナビゲーション用の属性。 bldg:InteriorWallSurface、bldg:CeilingSurface、bldg:FloorSurfaceに使用する。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IndoorAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:source     [ IndoorAttribute ] | gml::CodeType [0..1] | 原典資料。コードリスト（Common_indoorSource.xml）から選択する。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:floorId     [ IndoorZoneAttribute ] | xs::string [0..1] | 任意の空間が紐づけられている階層の固有ID。 |

