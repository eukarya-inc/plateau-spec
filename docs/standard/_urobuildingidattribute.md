# 4.2.3.2.1 uro:BuildingIDAttribute

## 4.2.3.2.1 uro:BuildingIDAttribute

**表 4-45**

| 型の定義 | 建築物を識別するための情報。 | ← |
| --- | --- | --- |
| 上位の型 | uro:BuildingAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:buildingID     [ BuildingIDAttribute ] | xs::string [1..1] | 主たる建築物を識別するための番号。必須とする。`[市区町村コード]-[接頭辞]-[オブジェクト連番]` とする。`[市区町村コード]` は、当該地物が存在する市区町村に該当するコードとする。先頭の0は省略せず、5桁で記述する。`[接頭辞]`は地物型の区分を示す3桁又は4桁のコードとする。建築物の場合は、bldgとする。`[オブジェクト連番]`は半角数字の連番とする。 |
| uro:branchID     [ BuildingIDAttribute ] | xs::integer [0..1] | 主たる建築物に対して付帯する建築物を識別するための番号。 |
| uro:partID     [ BuildingIDAttribute ] | xs::integer [0..1] | 主たる建築物を複数のbldg:BuildingPartに分けて記述する場合の、建築物部分を識別するための番号。bldg:BuildingPartには必須とする。 |
| uro:prefecture     [ BuildingIDAttribute ] | gml::CodeType [0..1] | 建築物が所在する都道府県の都道府県コ－ド。JIS X0401に定義される2桁の半角数字。コードリスト（Common_localPublicAuthorities.xml）より選択する。 |
| uro:city     [ BuildingIDAttribute ] | gml::CodeType [1..1] | 建築物が所在する市区町村の市区町村コ－ド。JIS X0401に定義される2桁の半角数字とJIS X0402に定義される3桁の半角数字とを組み合わせた5桁の半角数字。政令市の場合は、区の市区町村コードとする。コードリスト（Common_localPublicAuthorities.xml）より選択する。i-URでは多重度が[0..1]となっているが、建築物の位置の把握に使用するため、標準製品仕様書では必須とする。 |

