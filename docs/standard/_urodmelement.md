# 4.25.3.1.2 uro:DmElement

## 4.25.3.1.2 uro:DmElement

**表 4-860**

| 型の定義 | 数値地形図データの要素レコードの情報を保持するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:locationType     [ DmElement ] | gml::CodeType [0..1] | 地域分類。文字列とする。必要に応じて利用者が任意に定義するコード。 |
| uro:infoType     [ DmElement ] | gml::CodeType [0..1] | 情報分類。文字列とする。必要に応じて利用者が任意に定義するコード。 |
| uro:elementKey     [ DmElement ] | xs::string [0..1] | 要素識別番号。 |
| uro:hierarchyLevel     [ DmElement ] | xs::string [0..1] | 階層レベル。 |
| uro:dataType     [ DmElement ] | gml::CodeType [0..1] | 実データ区分。コードリスト（ DmElement_dataType.xml）から選択する。 |
| uro:annotationType     [ DmElement ] | gml::CodeType [0..1] | 注記区分。コードリスト（ DmElement_annotationType.xml）から選択する。 |
| uro:precisionType     [ DmElement ] | gml::CodeType [0..1] | 精度区分。コードリスト（ DmElement_precisionType.xml）から選択する。 |
| uro:dislocationType     [ DmElement ] | gml::CodeType [0..1] | 転位区分。コードリスト（ DmElement_dislocationType.xml）から選択する。 |
| uro:breakType     [ DmElement ] | gml::CodeType [0..1] | 間断区分。コードリスト（ DmElement_breakType.xml）から選択する。 |
| uro:attributeValue     [ DmElement ] | xs::string [0..1] | 属性数値。 |
| uro:attributeType     [ DmElement ] | gml::CodeType [0..1] | 属性区分。利用者が独自に設ける区分。文字列とする。 |
| uro:attributeValueType     [ DmElement ] | xs::string [0..1] | 属性データ書式。属性レコードを持つ場合の、そのレコードに記述されている内容の書式をFortran形式で記述する。 |
| uro:creationDate     [ DmElement ] | xs::gYearMonth [0..1] | 取得年月。 |
| uro:updateDate     [ DmElement ] | xs::gYearMonth [0..1] | 更新年月。 |
| uro:terminationDate     [ DmElement ] | xs::gYearMonth [0..1] | 削除年月。 |
| uro:freeSpace     [ DmElement ] | xs::string [0..1] | 空き領域。数値地形図データファイル形式で空き領域にデータが設定されている場合には，この属性を用いて保持する。 |

