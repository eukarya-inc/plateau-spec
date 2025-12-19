# 4.2.3.5.25 uro:IfcSpaceBaseQuantity

## 4.2.3.5.25 uro:IfcSpaceBaseQuantity

**表 4-82**

| 型の定義 | IFCで記述されたSpaceの数量に関する属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:nominalHeight     [ IfcSpaceBaseQuantity ] | gml::LengthType [0..1] | スラブ上端から上階スラブ下端までの高さ（予備寸法）。単位はm。 |
| uro:clearHeight     [ IfcSpaceBaseQuantity ] | gml::LengthType [0..1] | 床面（仕上げを含む）と天井面（仕上げ、下地を含む）の高さ。単位はm。 |
| uro:finishCeilingHeight     [ IfcSpaceBaseQuantity ] | gml::LengthType [0..1] | 天井高。例：床仕上げの上部面から天井の下部面までの高さ。単位はm。 |
| uro:grossPerimeter     [ IfcSpaceBaseQuantity ] | gml::LengthType [0..1] | 床レベルでの総周辺長（開口部の外周部分を含む）。単位はm。 |
| uro:netPerimeter     [ IfcSpaceBaseQuantity ] | gml::LengthType [0..1] | 正味周囲長（開口部外周部分は含まない）。単位はm。 |
| uro:grossCeilingArea     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 天井面積。単位はm2。 |
| uro:grossFloorArea     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 延面積（通常、柱、内壁などの面積も含まれる）。単位はm2。 |
| uro:netCeilingArea     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 正味天井面積（通常、柱、床開口部などの面積は除く）。単位はm2。 |
| uro:netFloorArea     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 正味延面積（通常、柱、床開口などの面積は除く）。単位はm2。 |
| uro:grossWallArea     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 壁面積（ドア、窓などの開口部分も含む）。単位はm2。 |
| uro:netWallArea     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 正味壁面積（ドア、窓などの開口部分を除く）。単位はm2。 |
| uro:grossVolume     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 体積（通常空間内の建築要素の体積も含む）。単位はm3。 |
| uro:netVolume     [ IfcSpaceBaseQuantity ] | gml::MeasureType [0..1] | 正味体積（空間内の建築要素の体積は除く）。単位はm3。 |

