# 4.2.3.5.17 uro:IfcPsetDoorCommon

## 4.2.3.5.17 uro:IfcPsetDoorCommon

**表 4-74**

| 型の定義 | IFCで記述された扉に共通の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:reference     [ IfcPsetDoorCommon ] | xs::string [0..1] | このプロジェクトのための参照記号。 |
| uro:acousticRating     [ IfcPsetDoorCommon ] | xs::string [0..1] | 遮音等級情報。関連する建築基準法を参照。 |
| uro:firerating     [ IfcPsetDoorCommon ] | xs::string [0..1] | 主要な耐火等級。関連する建築基準法、消防法などの国家基準を参照。 |
| uro:securityRating     [ IfcPsetDoorCommon ] | xs::string [0..1] | 防犯等級情報。関連する基準を参照。 |
| uro:isExternal     [ IfcPsetDoorCommon ] | xs::boolean [0..1] | 外部の部材かどうかを示すブーリアン値。1：外部の部材で建物の外側に面している0：そうではない |
| uro:infiltration     [ IfcPsetDoorCommon ] | xs::double [0..1] | 隙間風の流量値。 |
| uro:thermalTransmittance     [ IfcPsetDoorCommon ] | xs::double [0..1] | 熱貫流率U値。ここでは扉を通した熱移動の方向における全体の熱還流率を示す。 |
| uro:glazingAreaFraction     [ IfcPsetDoorCommon ] | xs::double [0..1] | 外壁の総面積に対するガラスの面積の比率。 ガラスの面積が外壁に含まれる全てのパネルと分離されていないときに、使用される。 |
| uro:handicapAccessible     [ IfcPsetDoorCommon ] | xs::boolean [0..1] | 障碍者にアクセスできるように設計されているか否か。1：されている0：されていない |
| uro:fireExit     [ IfcPsetDoorCommon ] | xs::boolean [0..1] | 火災時の出口として使用されるように設計されているか否か1：されている0：されていない |
| uro:selfClosing     [ IfcPsetDoorCommon ] | xs::boolean [0..1] | 扉が使用後に自動で閉まるか否か。1：閉まる0：閉まらない |
| uro:smokeStop     [ IfcPsetDoorCommon ] | xs::boolean [0..1] | オブジェクトが煙止めを提供するように設計されているか否か。1：されている0：されていない |

