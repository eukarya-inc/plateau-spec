# 4.2.3.2.13 uro:RealEstateIDAttribute

## 4.2.3.2.13 uro:RealEstateIDAttribute

**表 4-57**

| 型の定義 | 建築物が存在する土地及び建築物に紐づく不動産IDを、建築物の属性として付与するためのデータ型。不動産IDとは、「不動産を一意に特定することができる、各不動産の共通コード」である。（不動産IDルールガイドライン） | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:realEstateIDOfBuilding     [ RealEstateIDAttribute ] | xs::string [1..1] | 建築物の「建築物全体」に付された不動産ID。不動産IDルールガイドライン（国土交通省）に基づく「不動産番号13桁＋”－”（ハイフン）＋4桁」で記載される。 |
| uro:numberOfBuildingUnitOwnership     [ RealEstateIDAttribute ] | xs::integer [0..1] | 当該建築物が区分所有の場合の、当該建築物の区分所有の数量。 |
| uro:realEstateIDOfBuildingUnitOwnership     [ RealEstateIDAttribute ] | xs::string [0..*] | 当該建築物が区分所有の場合の、当該建築物の各専有部分の不動産IDの一覧。不動産IDルールガイドライン（国土交通省）に基づく「不動産番号13桁＋”－”（ハイフン）＋4桁」で記載される。 |
| uro:numberOfRealEstateIDOfLand     [ RealEstateIDAttribute ] | xs::integer [0..1] | 当該建築物のある土地（筆）の数量。 |
| uro:realEstateIDOfLand     [ RealEstateIDAttribute ] | xs::string [0..*] | 当該建築物のある土地の不動産ID。不動産IDルールガイドライン（国土交通省）に基づく「不動産番号13桁＋”－”（ハイフン）＋4桁」で記載される。複数の土地にまたがる場合、建築物に紐づけられた登記簿の「所在」欄に記載されている地番の順番に従う。 |
| uro:matchingScore     [ RealEstateIDAttribute ] | xs::integer [1..1] | 建築物に不動産IDを付与する際に、Project PLATEAUが2023年度に開発した「dt23-03 3D都市モデル・不動産IDマッチングシステム」を用いる場合は、確率的なマッチング手法を採用していることから、当該建築物と登記簿データが付与された筆との空間属性及び主題属性の一致の程度が点数化して表される。uro:matchingScoreはその値を表す。 「dt23-03 3D都市モデル・不動産IDマッチングシステム」を用いる場合は必須とする。  uro:matchingScoreは、以下に示す①から④の項目のうち、「①、②及び④」又は「①、③及び④」の3項目の各点数の合計点（300点満点）の大きい方を100点満点に換算して算出する。  　  ①重なりスコア：建築物に紐づけられた登記簿の「所在」欄に記載されている土地と建築物を土地に投影した二次元空間属性の重なり割合（％）![Image](images/_urorealestateidattribute_img001.png)　 ②階数スコア：建築物が地上階数をbldg:storeysBelowGround属性に保持している場合、建築物に紐づけられた登記簿の「床面積」欄から算出した地上階数と建築物の地上階数の一致度地上階数が一致した場合は100点とする。一致しない場合は0点とする。  　  ③建築物高さスコア：建築物に紐づけられた登記簿の「床面積」欄から算出した地上階数から換算した高さと建築物の計測高さ(bldg:measuredHeight)の一致度![Image](images/_urorealestateidattribute_img002.png)　 ABS（X）は、Xの絶対値とする。このときXは数値でなければならない。式の値が負の場合は0とする。  　  ④床面積スコア：建築物に紐づけられた登記簿の「床面積」欄から算出した各階ごとの床面積のうち最大となる床面積と建築物の図形面積の一致度建築物がuro:buildingFootprintArea属性を保持している場合、登記簿から算出した面積と比較し、m2単位で一致していた場合は100点とする。属性がない場合またはm2単位で一致していなければ以下を算出する。![Image](images/_urorealestateidattribute_img003.png)　 ABS（X）は、Xの絶対値とする。このときXは数値でなければならない。式の値が負の場合は0とする。 建築物の図形面積は、次の建築物の各空間属性のうち、最初に存在する空間属性の水平投影面積とする。  bldg:lod0RoofEdge、bldg:lod1Solid、bldg:lod2Solid.RoofSurface、bldg:lod3Solid.RoofSurface |

