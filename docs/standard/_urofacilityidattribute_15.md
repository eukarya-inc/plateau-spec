# 4.26.3.1.1 uro:FacilityIdAttribute

## 4.26.3.1.1 uro:FacilityIdAttribute

**表 4-871**

| 型の定義 | 管理施設の識別に関する情報を定義したデータ型。 河川管理施設の場合にはこれを継承する下位型を用いて記述する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:id     [ FacilityIdAttribute ] | xs::string [0..1] | 構造物の識別子。 |
| uro:partId     [ FacilityIdAttribute ] | xs::string [0..1] | 構造物を部分（Part）に分けて記述する場合の各部分を識別するための番号。河川管理施設の場合は使用しない。 |
| uro:branchId     [ FacilityIdAttribute ] | xs::string [0..1] | 枝番。同一の構造物ではないが、一連の構造物として管理したい場合に、それぞれを識別するための番号。 |
| uro:prefecture     [ FacilityIdAttribute ] | gml::CodeType [0..*] | 構造物が所在する都道府県の都道府県コ－ド。JIS X0401に定義される2桁の半角数字。コードリスト（Common_localPublicAuthorities.xml）より選択する。複数の都道府県に跨って存在する場合は、複数の都道府県コードを記述する。 |
| uro:city     [ FacilityIdAttribute ] | gml::CodeType [0..*] | 構造物が所在する市区町村の市区町村コ－ド。JIS X0401に定義される2桁の半角数字とJIS X0402に定義される3桁の半角数字とを組み合わせた5桁の半角数字。政令市の場合は、区の市区町村コードとする。コードリスト（Common_localPublicAuthorities.xml）より選択する。 複数の市区町村に跨って存在する場合は、複数の市区町村コードを記述する。 |
| uro:route     [ FacilityIdAttribute ] | xs::string [0..1] | 構造物が存在する路線名。河川管理施設の場合は使用しない。 |
| uro:startPost     [ FacilityIdAttribute ] | xs::string [0..1] | 構造物の開始位置の距離標。河川管理施設の場合は使用しない。 |
| uro:endPost     [ FacilityIdAttribute ] | xs::string [0..1] | 構造物の終了位置の距離標。河川管理施設の場合は使用しない。 |
| uro:startLat     [ FacilityIdAttribute ] | xs::double [0..1] | 構造物の開始位置の緯度（北緯）。10進数により記述する。河川管理施設の場合は使用しない。 |
| uro:startLong     [ FacilityIdAttribute ] | xs::double [0..1] | 構造物の開始位置の経度（東経）。10進数により記述する。河川管理施設の場合は使用しない。 |
| uro:alternativeName     [ FacilityIdAttribute ] | xs::string [0..*] | 別名。gml:nameで記述する正式な名称以外に、一般に普及している名称がある場合に記述する。 |

