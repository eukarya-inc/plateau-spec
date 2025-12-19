# 4.2.3.2.4 uro:RiverFloodingRiskAttribute

## 4.2.3.2.4 uro:RiverFloodingRiskAttribute

**表 4-48**

| 型の定義 | 洪水浸水想定区域内に存在する建築物に、浸水想定区域がもつ属性を与えるための属性型。同一の浸水想定区域図において、複数の区域に建築物が跨って存在する場合は、同一浸水ランクを持つ浸水ランクのメッシュを一つの区域とし、その区域と建築物が重なる面積が最も大きい浸水ランクの値を採用する。（面積が等しい場合は、浸水ランクがより危険な区域を採用する） 浸水深は採用した浸水ランクを持つ浸水深のメッシュのうち、建築物と重なる面積が最も大きいメッシュの浸水深を採用する。（同じ浸水深を持つメッシュは面積算出の際に合算する） 浸水継続時間は採用した浸水深のメッシュと重なる浸水継続時間のメッシュの浸水継続時間を採用する。複数の浸水継続時間のメッシュが重なる場合は最も大きい浸水継続時間の値を採用する。 浸水深の有効桁数は、「浸水想定区域図データ電子化ガイドライン（第4版）」に従い、浸水深の有効桁数は、小数点以下 3 桁まで登録可能とするが、小数点以下 2 桁でもよいとする。 面積の有効桁数は、小数点2桁（3桁目で四捨五入）とする。![Image](images/_uroriverfloodingriskattribute_img001.png) | ← |
| --- | --- | --- |
| 上位の型 | uro:FloodingRiskAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:description     [ DisasterRiskAttribute ] | gml::CodeType [1..1] | 使用する下位のデータ型により以下のように定義する。洪水浸水想定区域の場合指定河川の名称。コードリスト（RiverFloodingRiskAttribute_description.xml）より選択する。都市ごとにコードリストを作成する。指定河川の名称には、水防法に基づき指定された洪水浸水想定区域図の対象となる洪水予報河川又は水位周知河川として示された、「水系名」及び「指定河川名」を用いることを基本とする。 一つの浸水想定区域図に複数の洪水予報河川又は水位周知河川が含まれている場合は、「指定河川名」を列挙する。指定河川名を列挙する場合の区切り文字は「・」（全角中点）を使用する。また、都道府県が独自に作成している浸水の区域図は、当該浸水想定区域の名称から、対象となる区域を指す名称を用いる。津波浸水想定の場合津波浸水想定の属性を付与する元となる図又はデ－タの名称。コードリスト（TsunamiRiskAttribute_description.xml）より選択する。都市ごとにコードリストを作成する。高潮浸水想定区域の場合高潮浸水想定区域の属性を付与する元となる図又はデ－タ集合の名称。コードリスト（HighTideRiskAttribute_description.xml）より選択する。都市ごとにコードリストを作成する。内水浸水想定区域の場合内水浸水想定区域の属性を付与する元となる図又はデ－タの名称。コードリスト（InlandFloodingRiskAttribute_description.xml）より選択する。都市ごとにコードリストを作成する。ため池ハザードマップの場合ため池ハザードマップの属性を付与する元となる図又はデ－タの名称。コードリスト（ReservoirFloodingRiskAttribute_description.xml）より選択する。都市ごとにコードリストを作成する。土砂災害警戒区域の場合発生が想定されている災害の種類。コードリスト（LandSlideRiskAttribute_description.xml）より選択する。 |
| uro:rank     [ FloodingRiskAttribute ] | gml::CodeType [0..1] | 浸水深（津波浸水想定の場合は「水位」）に応じた区分。uro:rank又はuro:rankOrgのいずれか一つをもつ。洪水浸水想定区域の場合コードリスト（RiverFloodingRiskAttribute_rank.xml）より選択する。津波浸水想定の場合コードリスト（TsunamiRiskAttribute_rank.xml）より選択する。高潮浸水想定区域の場合コードリスト（HighTideRiskAttribute_rank.xml）より選択する。内水浸水想定区域の場合コードリスト（InlandFloodingRiskAttribute_rank.xml）より選択する。ため池ハザードマップの場合コードリスト（ReservoirFloodingRiskAttribute_rank.xml）より選択する。なお、津波浸水想定の場合の水位は、「津波基準水位」がある場合はこれを採用し、ない場合は「津波浸水想定に定める水深に係る水位」とする。「津波基準水位」とは、「津波浸水想定に定める水深に係る水位に建築物等への衝突による津波の水位の上昇（せき上げ）を考慮して必要と認められる値を加えて定める水位」（『津波浸水想定の設定の手引き』参照）である。 |
| uro:rankOrg     [ FloodingRiskAttribute ] | gml::CodeType [0..1] | 市町村等が独自に設定した浸水深（津波浸水想定の場合は「水位」）に応じた区分。uro:rank又はuro:rankOrgのいずれか一つをもつ。洪水浸水想定区域の場合コードリスト（RiverFloodingRiskAttribute_rankOrg.xml）より選択する。津波浸水想定の場合コードリスト（TsunamiRiskAttribute_rankOrg.xml）より選択する。高潮浸水想定区域の場合コードリスト（HighTideRiskAttribute_rankOrg.xml）より選択する。内水浸水想定区域の場合コードリスト（InlandFloodingRiskAttribute_rankOrg.xml）より選択する。ため池ハザードマップの場合コードリスト（ReservoirFloodingRiskAttribute_rankOrg.xml）より選択する。なお、津波浸水想定の場合の水位は、「津波基準水位」がある場合はこれを採用し、ない場合は「津波浸水想定に定める水深に係る水位」とする。「津波基準水位」とは、「津波浸水想定に定める水深に係る水位に建築物等への衝突による津波の水位の上昇（せき上げ）を考慮して必要と認められる値を加えて定める水位」（『津波浸水想定の設定の手引き』参照）である。 |
| uro:depth     [ FloodingRiskAttribute ] | gml::LengthType [0..1] | 浸水の深さ。津波浸水想定、高潮浸水想定、内水浸水想定区域及びため池ハザードマップの場合は、陸上の各地点で水面が最も高い位置にきたときの地面から水面までの高さ。 単位はm（uom=”m”）とする。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:adminType     [ RiverFloodingRiskAttribute ] | gml::CodeType [1..1] | 洪水予報河川又は水位周知河川を指定した機関の別。 コードリスト（ RiverFloodingRiskAttribute_adminType.xml）より選択する。 |
| uro:scale     [ RiverFloodingRiskAttribute ] | gml::CodeType [1..1] | 想定最大規模降雨あるいは計画規模降雨のいずれにより作成されたかの区分。 コードリスト（ RiverFloodingRiskAttribute_scale.xml）より選択する。 |
| uro:duration     [ RiverFloodingRiskAttribute ] | gml::MeasureType [0..1] | 浸水が継続する時間。単位は時間（uom=”hour”）とする。 |

