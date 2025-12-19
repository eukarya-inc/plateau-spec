# 4.5.3.2.4 uro:TrackAttribute

## 4.5.3.2.4 uro:TrackAttribute

**表 4-241**

| 型の定義 | 徒歩道に関する情報を定義したデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:alternativeName     [ TrackAttribute ] | xs::string [0..*] | 徒歩道の名称。複数の名称を格納したい場合に使用する。本属性を使用する場合、gml:nameは必須とする。 |
| uro:adminType     [ TrackAttribute ] | gml::CodeType [0..1] | 徒歩道の管理者の区分。コードリスト（TrackAttribute_adminType.xml）より選択する。 |
| uro:relativeLevel     [ TrackAttribute ] | xs::integer [0..1] | 階層順。道路及び鉄道の立体交差部や、道路の上に建物が建設されている場合などにおける、階層の相対順位を表す値。0が最も下層にあることを示す。 |
| uro:widthType     [ TrackAttribute ] | gml::CodeType [0..1] | 道路構造の区分。コードリスト（TrackAttribute_widthType.xml）より選択する。 |
| uro:structureType     [ TrackAttribute ] | gml::CodeType [0..1] | 道路構造の区分。コードリスト（TrackAttribute_structureType.xml）より選択する。 |
| uro:isTollRoad     [ TrackAttribute ] | xs::boolean [0..1] | 走行するために料金が徴収されるかどうかを示す。有料の場合に1とする。 |
| uro:separator     [ TrackAttribute ] | gml::LengthType [0..1] | 分離帯がある道路であることを示す。分離帯の幅が1m単位で設定される。単位はmとする。 |

