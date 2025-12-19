# 4.2.3.5.21 uro:IfcPsetWindowCommon

## 4.2.3.5.21 uro:IfcPsetWindowCommon

**表 4-78**

| 型の定義 | IFCで記述された窓に共通の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:reference     [ IfcPsetWindowCommon ] | xs::string [0..1] | このプロジェクトのための参照記号。 |
| uro:acousticRating     [ IfcPsetWindowCommon ] | xs::string [0..1] | 遮音等級情報。関連する建築基準法を参照。 |
| uro:fireRating     [ IfcPsetWindowCommon ] | xs::string [0..1] | 主要な耐火等級。関連する建築基準法、消防法などの国家基準を参照。 |
| uro:securityRating     [ IfcPsetWindowCommon ] | xs::string [0..1] | 防犯等級情報。関連する基準を参照。 |
| uro:isExternal     [ IfcPsetWindowCommon ] | xs::boolean [0..1] | 外部の部材かどうかを示すブーリアン値。1：外部の部材で建物の外側に面している0：そうではない |
| uro:infiltration     [ IfcPsetWindowCommon ] | xs::double [0..1] | 隙間風の流量値。 |
| uro:thermalTransmittance     [ IfcPsetWindowCommon ] | xs::double [0..1] | 熱貫流率U値。ここでは窓を通した熱移動の方向における全体の熱還流率を示す。 |
| uro:glazingAreaFraction     [ IfcPsetWindowCommon ] | xs::double [0..1] | 外壁の総面積に対するガラスの面積の比率。 ガラスの面積が外壁に含まれる全てのパネルと分離されていないときに、使用される。 |
| uro:smokeStop     [ IfcPsetWindowCommon ] | xs::boolean [0..1] | オブジェクトが煙止めを提供するように設計されているか否か。1：されている0：されていない |

