# 4.3.3.2.5 uro:TrafficVolumeAttribute

## 4.3.3.2.5 uro:TrafficVolumeAttribute

**表 4-157**

| 型の定義 | 道路の交通量に関する情報。全国道路・街路交通情勢調査一般交通量調査の対象となる高速自動車国道、都市高速道路、一般国道、主要地方道である都道府県道及び指定市の市道、一般都道府県道、指定市の一部の一般市道を対象とする。 | ← |
| --- | --- | --- |
| 上位の型 | uro:RoadAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:sectionID     [ TrafficVolumeAttribute ] | xs::string [0..1] | 交通量調査において、調査の単位となる交通調査基本区間に付与される番号。原則として「都道府県（2 桁）」＋「道路種別（1 桁）」＋「路線番号（4 桁）」＋「順番号（4 桁）」からなる 11 桁の番号。 |
| uro:routeName     [ TrafficVolumeAttribute ] | xs::string [0..1] | 路線名。 |
| uro:weekday12hourTrafficVolume     [ TrafficVolumeAttribute ] | xs::integer [0..1] | 平日7時～19時までに通過する車両台数。単位は台とする。 |
| uro:weekday24hourTrafficVolume     [ TrafficVolumeAttribute ] | xs::integer [0..1] | 平日7時～翌朝7時又は0時～翌日0時までに通過する車両台数。単位は台とする。 |
| uro:largeVehicleRate     [ TrafficVolumeAttribute ] | xs::double [0..1] | 自動車類交通量に対する大型車交通量の割合。単位は％とする。 |
| uro:congestionRate     [ TrafficVolumeAttribute ] | xs::double [0..1] | 交通調査基本区間の交通容量に対する交通量の比。単位は％とする。 |
| uro:averageTravelSpeedInCongestion     [ TrafficVolumeAttribute ] | xs::double [0..1] | 朝のラッシュ時間帯（7時～9時）又は夕方のラッシュ時間帯（17時～19時）において平均旅行速度を集計し、その遅い方の時間帯の旅行速度。都市計画基礎調査で収集されている場合にのみ作成する。単位はkm/hとする。 |
| uro:averageInboundTravelSpeedInCongestion     [ TrafficVolumeAttribute ] | xs::double [0..1] | 朝のラッシュ時間帯（7時～9時）又は夕方のラッシュ時間帯（17時～19時）において上り線における平均旅行速度を集計し、その遅い方の時間帯の旅行速度。単位はkm/hとする。 |
| uro:averageOutboundTravelSpeedInCongestion     [ TrafficVolumeAttribute ] | xs::double [0..1] | 朝のラッシュ時間帯（7時～9時）又は夕方のラッシュ時間帯（17時～19時）において下り線における平均旅行速度を集計し、その遅い方の時間帯の旅行速度。単位はkm/hとする。 |
| uro:averageInboundTravelSpeedNotCongestion     [ TrafficVolumeAttribute ] | xs::double [0..1] | 昼間非混雑時（9～17時）における上り線の平均旅行速度。 単位はkm/hとする。 |
| uro:averageOutboundTravelSpeedNotCongestion     [ TrafficVolumeAttribute ] | xs::double [0..1] | 昼間非混雑時（9～17時）における下り線平均旅行速度。 単位はkm/hとする。 |
| uro:observationPointName     [ TrafficVolumeAttribute ] | xs::string [0..1] | 交通量等を観測した地点の名称。 |
| uro:reference     [ TrafficVolumeAttribute ] | xs::string [0..1] | 対象となる道路の区間を図上で識別する番号。 |
| uro:surveyYear     [ TrafficVolumeAttribute ] | xs::gYear [0..1] | 調査が実施された年。本データ型を作成する場合は、必須とする。 |

