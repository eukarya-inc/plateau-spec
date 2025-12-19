# 4.19.3.2.1 uro:WaterBodyDetailAttribute

## 4.19.3.2.1 uro:WaterBodyDetailAttribute

**表 4-801**

| 型の定義 | 水部の基盤的な情報。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:kana     [ WaterBodyDetailAttribute ] | xs::string [0..1] | 水部の名称のフリガナ。カタカナで記述する。 |
| uro:waterSystemCode     [ WaterBodyDetailAttribute ] | gml::CodeType [0..1] | 水系域コード。河川コード仕様書（国土交通省河川局）にて示された2桁の地域番号と4桁の水系番号からなる6桁の番号。 （水系域コード一覧：https://nlftp.mlit.go.jp/ksj/gml/codelist/WaterSystemCodeCd.html） 地域番号は、1級水系の場合は地方整備局等の番号、1級水系以外は都道府県の番号となる。また、水系番号は、一水系につき一つ与えられた番号であり、地域番号と併せて使用することで、水系の特定が可能となる。 コードリスト（WaterBodyDetailAttribute_SystemCode.xml）より選択する。 |
| uro:riverCode     [ WaterBodyDetailAttribute ] | gml::CodeType [0..1] | 河川コード。河川コード仕様書（国土交通省河川局）にて示された、１級河川、２級河川、準用河川、普通河川に該当する個別の河川を一意に特定するために付与された2桁の地域番号、4桁の水系番号、4桁の河川番号からなる10桁の番号。 （河川コード一覧：https://nlftp.mlit.go.jp/ksj/gml/codelist/RiverCodeCd.html） 河川番号、同一水系内において河川を特定するための番号であり、一河川につき一つの番号が付与される。 コードリスト（WaterBodyDetailAttribute_riverCode.xml）より選択する。 |
| uro:adminType     [ WaterBodyDetailAttribute ] | gml::CodeType [0..1] | 河川管理者が管理する区間種別。コードリスト（WaterBodyDetailAttribute_adminType.xml）より選択する。 |
| uro:flowDirection     [ WaterBodyDetailAttribute ] | xs::boolean [0..1] | 水部の流下方向の判明状況。 |
| uro:maximumDepth     [ WaterBodyDetailAttribute ] | gml::LengthType [0..1] | 最大水深。単位はｍ（uom=”m”）とする。 |
| uro:waterSurfaceElevation     [ WaterBodyDetailAttribute ] | gml::LengthType [0..1] | 水面標高。単位はｍ（uom=”m”）とする。 |
| uro:area     [ WaterBodyDetailAttribute ] | gml::MeasureType [0..1] | 水部の範囲の実測により取得した面積。単位はm2（uom=”m2”）とする。 |
| uro:measurementYearMonth     [ WaterBodyDetailAttribute ] | xs::gYearMonth [0..1] | 水部を測量した年月。 |
| uro:prefecture     [ WaterBodyDetailAttribute ] | gml::CodeType [0..*] | 水部が所在する都道府県の都道府県コ－ド。JIS X0401に定義される2桁の半角数字。コードリスト（Common_localPublicAuthorities.xml）より選択する。複数の都道府県に跨って存在する場合は、複数の都道府県コードを記述する。 |
| uro:city     [ WaterBodyDetailAttribute ] | gml::CodeType [0..*] | 水部が所在する市区町村の市区町村コ－ド。JIS X0401に定義される2桁の半角数字とJIS X0402に定義される3桁の半角数字とを組み合わせた5桁の半角数字。政令市の場合は、区の市区町村コードとする。コードリスト（Common_localPublicAuthorities.xml）より選択する。複数の市区町村に跨って存在する場合は、複数の市区町村コードを記述する。 |

