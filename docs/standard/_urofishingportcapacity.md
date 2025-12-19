# 4.26.3.4.2 uro:FishingPortCapacity

## 4.26.3.4.2 uro:FishingPortCapacity

**表 4-890**

| 型の定義 | 漁港施設の能力を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | uro:FishingPortAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:facilityId     [ FacilityAttribute ] | xs::string [0..1] | 施設の用途や区分ごとに付与された管理番号。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:capacity     [ FishingPortCapacity ] | xs::string [0..1] | 能力。当該施設が、外郭施設、輸送施設、漁港施設用地、増殖及び養殖用施設、漁獲物施設、漁業用通信施設、環境整備施設、廃船施設、廃油施設、浄化施設、漁港管理施設のいずれかの場合に記述する。 |
| uro:weightCapacity     [ FishingPortCapacity ] | gml::MeasureType [0..1] | 能力_耐重量。 当該施設が係留施設の場合に記述する。 |
| uro:hullForm     [ FishingPortCapacity ] | xs::integer [0..1] | 能力_係船能力_船型。 当該施設が係留施設の場合に記述する。 |
| uro:shipNumber     [ FishingPortCapacity ] | xs::integer [0..1] | 能力_係船能力_隻数。 当該施設が係留施設の場合に記述する。 |
| uro:waterDepth-2m     [ FishingPortCapacity ] | gml::MeasureType [0..1] | 能力_水深別内訳_2ｍ未満の面積。 当該施設が水域施設の場合に記述する。 |
| uro:waterDepth2-3m     [ FishingPortCapacity ] | gml::MeasureType [0..1] | 能力_水深別内訳_2～3ｍ未満の面積。 当該施設が水域施設の場合に記述する。 |
| uro:waterDepth3-6m     [ FishingPortCapacity ] | gml::MeasureType [0..1] | 能力_水深別内訳_3～6ｍ未満の面積。 当該施設が水域施設の場合に記述する。 |
| uro:waterDepth6-m     [ FishingPortCapacity ] | gml::MeasureType [0..1] | 能力_水深別内訳_6ｍ以上の面積。 当該施設が水域施設の場合に記述する。 |
| uro:heightAboveAWL     [ FishingPortCapacity ] | gml::LengthType [0..1] | 能力_種類_灯台_平均水面上の高さ。 当該施設が航行補助施設の場合に記述する。 |
| uro:heightOnFoundations     [ FishingPortCapacity ] | gml::LengthType [0..1] | 能力_種類_灯台_基礎上の高さ。 当該施設が航行補助施設の場合に記述する。 |
| uro:luminousRange     [ FishingPortCapacity ] | gml::LengthType [0..1] | 能力_光音電波の到達距離。 当該施設が航行補助施設の場合に記述する。 |
| uro:luminousColor     [ FishingPortCapacity ] | xs::string [0..1] | 能力_灯色。 当該施設が航行補助施設の場合に記述する。 |
| uro:candlePower     [ FishingPortCapacity ] | xs::integer [0..1] | 能力_燭光数。 当該施設が航行補助施設の場合に記述する。 |
| uro:lightType     [ FishingPortCapacity ] | xs::string [0..1] | 能力_灯質の種類。 当該施設が航行補助施設の場合に記述する。 |
| uro:period     [ FishingPortCapacity ] | xs::string [0..1] | 能力_灯質の周期。 当該施設が航行補助施設の場合に記述する。 |
| uro:maximumGroundingWeight     [ FishingPortCapacity ] | xs::integer [0..1] | 能力_入きょ又は上架できる最大船舶の総重量。 当該施設が漁船漁具保全施設の場合に記述する。 |
| uro:handleablePower     [ FishingPortCapacity ] | xs::integer [0..1] | 能力_取り扱いできる機関の馬力数。 当該施設が漁船漁具保全施設の場合に記述する。 |
| uro:maximumWaterSupply     [ FishingPortCapacity ] | xs::integer [0..1] | 能力_最大給水能力。 当該施設が補給施設の場合に記述する。 |
| uro:maximumRefueling     [ FishingPortCapacity ] | xs::string [0..1] | 能力_最大給油能力。 当該施設が補給施設の場合に記述する。 |
| uro:people     [ FishingPortCapacity ] | xs::integer [0..1] | 能力_最大収容可能人数。 当該施設が厚生施設の場合に記述する。 |
| uro:other     [ FishingPortCapacity ] | xs::string [0..1] | 能力_その他。 当該施設が係留施設、水域施設、漁船漁具保全施設、補給施設、漁港厚生施設の場合に必要に応じて記述する。 |

