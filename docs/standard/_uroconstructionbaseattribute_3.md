# 4.13.3.1.9 uro:ConstructionBaseAttribute

## 4.13.3.1.9 uro:ConstructionBaseAttribute

**表 4-626**

| 型の定義 | 構造物の管理に必要となる基本的な情報及び、構造物の規模に関する基本的な情報を記述するためのデータ型。構造物が橋梁、トンネル、護岸、床止め、堰、水門・閘門・陸閘、樋門・樋管、伏せ越し及び水制の場合に本データ型を用いて記述する。 構造物がダム及び堤防の場合は、これを継承するデータ型を用いて記述する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:adminType     [ ConstructionBaseAttribute ] | gml::CodeType [0..1] | 構造物の管理者の区分。コードリスト（ConstructionBaseAttribute_adminType.xml）から選択する。 |
| uro:administrator     [ ConstructionBaseAttribute ] | xs::string [0..1] | 構造物の管理者の名称。 |
| uro:adminOffice     [ ConstructionBaseAttribute ] | xs::string [0..1] | 構造物の管理事務所所在地。 |
| uro:operatorType     [ ConstructionBaseAttribute ] | gml::CodeType [0..1] | 構造物の運用者の区分。コードリストから選択する。この属性を使用する場合は、コードリスト（ConstructionBaseAttribute_operatorType.xml）を作成すること。 |
| uro:installerType     [ ConstructionBaseAttribute ] | gml::CodeType [0..1] | 構造物の設置者の区分。コードリスト（ConstructionBaseAttribute_installerType.xml）から選択する。 |
| uro:installer     [ ConstructionBaseAttribute ] | xs::string [0..1] | 構造物の設置者の名称。 |
| uro:structureOrdinance     [ ConstructionBaseAttribute ] | xs::string [0..1] | 構造物が適用された構造令の名称。 |
| uro:specification     [ ConstructionBaseAttribute ] | xs::string [0..1] | 構造物が適用された示方書の名称。 |
| uro:kana     [ ConstructionBaseAttribute ] | xs::string [0..1] | 構造物の名称のふりがな。 |
| uro:constructionStartYear     [ ConstructionBaseAttribute ] | xs::gYear [0..1] | 構造物の建設開始年度。 |
| uro:completionYear     [ ConstructionBaseAttribute ] | xs::gYear [0..1] | 構造物の完成年度。 |
| uro:facilityAge     [ ConstructionBaseAttribute ] | xs::integer [0..1] | 工事完成年度からの年数。 |
| uro:update     [ ConstructionBaseAttribute ] | xs::date [0..1] | 更新年月日。 |
| uro:purpose     [ ConstructionBaseAttribute ] | gml::CodeType [0..1] | 構造物の建設の目的。コードリスト（ConstructionBaseAttribute_purpose.xml）から選択する。コードリストに無い場合は文字列により記述する。 |

