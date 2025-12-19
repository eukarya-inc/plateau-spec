# 4.2.3.5.20 uro:IfcPsetSpaceCommon

## 4.2.3.5.20 uro:IfcPsetSpaceCommon

**表 4-77**

| 型の定義 | IFCで記述された部屋に共通の属性。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:reference     [ IfcPsetSpaceCommon ] | xs::string [0..1] | このプロジェクトのための参照記号。 |
| uro:category     [ IfcPsetSpaceCommon ] | xs::string [0..1] | この部屋の用途。 |
| uro:floorCovering     [ IfcPsetSpaceCommon ] | xs::string [0..1] | この部屋の床材の材質又は仕上げ。 |
| uro:wallCovering     [ IfcPsetSpaceCommon ] | xs::string [0..1] | この部屋の壁材の材質又は仕上げ。 |
| uro:ceilingCovering     [ IfcPsetSpaceCommon ] | xs::string [0..1] | この部屋の天井カバーの材質又は仕上げ。 |
| uro:skirtingBoard     [ IfcPsetSpaceCommon ] | xs::string [0..1] | この部屋の幅木ボードの素材又は構造。 |
| uro:grossPlannedArea     [ IfcPsetSpaceCommon ] | gml::MeasureType [0..1] | 総計画面積。単位はm2とする。 |
| uro:netPlannedArea     [ IfcPsetSpaceCommon ] | gml::MeasureType [0..1] | 正味計画面積。単位はm2とする。 |
| uro:publiclyAccessible     [ IfcPsetSpaceCommon ] | xs::boolean [0..1] | この部屋（トイレなどの場合）が公衆の用に供するよう公的にアクセス可能な部屋として設計されているか。1：されている0：されていない |
| uro:handicapAccessible     [ IfcPsetSpaceCommon ] | xs::boolean [0..1] | この部屋（トイレなどの場合）が障害者用に供するような部屋として設計されているか。1：されている0：されていない |
| uro:concealedFlooring     [ IfcPsetSpaceCommon ] | xs::boolean [0..1] | この部屋が隠し床として定義されているか。隠し床は、通常上げ床の下のスペースを指す。1：されている0：されていない |
| uro:concealedCeiling     [ IfcPsetSpaceCommon ] | xs::boolean [0..1] | この部屋が隠し天井として定義されているか。隠し天井は、通常スラブと吊り天井の間のスペースを指す。1：されている0：されていない |

