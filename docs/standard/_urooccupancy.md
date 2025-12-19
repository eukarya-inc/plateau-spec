# 4.13.3.1.20 uro:Occupancy

## 4.13.3.1.20 uro:Occupancy

**表 4-637**

| 型の定義 | 構造物に居住又は格納される人、動物、その他の移動可能な物体についての定量的な情報。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:interval     [ Occupancy ] | gml::CodeType [0..1] | 占有する期間の種類。コードリストから選択する。この属性を使用する場合は、コードリスト（Occupancy_interval.xml）を作成すること。 |
| uro:numberOfOccupants     [ Occupancy ] | xs::integer [1..1] | 占有物の数。 |
| uro:occupantType     [ Occupancy ] | gml::CodeType [0..1] | 占有物の種類。コードリストから選択する。この属性を使用する場合は、コードリスト（Occupancy_occupantType.xml）を作成すること。 |

