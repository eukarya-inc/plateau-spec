# 4.27.3.1.1 core:CityModel

## 4.27.3.1.1 core:CityModel

**表 4-935**

| 型の定義 | 3次元都市モデルのための地物集合。 全ての都市オブジェクト及びその幾何形状等はこの地物型の中に含める。 | ← |
| --- | --- | --- |
| 上位の型 | gml:_FeatureCollection | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| app:appearanceMember             [ CityModel ] | app:Appearance [0..*] | 3D都市モデルに紐づけられたテクスチャ及びマテリアルの情報。 |
| core:cityObjectMember             [ CityModel ] | core:_CityObject [0..*] | 都市モデルに含まれる都市オブジェクト。 core:CityModelがcore:cityObjectMemberにより直接含む都市オブジェクトは、各応用スキーマにおいて、全体となる都市オブジェクトのみである。 それ以外の都市オブジェクトは、全体となる都市オブジェクトの部品として出現する。 全体となる都市オブジェクトとは、以下である。  bldg:Building  brid:Bridge  frn:CityFurniture  grp:CityObjectGroup  luse:LandUse  tran:Road  tran:Railway  tran:Track  tran:Square  tun:Tunnel  veg:SolitaryVegetationObject  veg:PlantCover  wtr:WaterBody  dem:ReliefFeature  uro:OtherConstruction  uro:UndergroundBuilding  uro:UtilityNetworkElementを継承する都市オブジェクト uro:Waterway  urf:Zone及びこれを継承する都市オブジェクト |

