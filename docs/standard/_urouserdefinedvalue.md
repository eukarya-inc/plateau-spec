# 4.2.3.5.37 uro:UserDefinedValue

## 4.2.3.5.37 uro:UserDefinedValue

**表 4-94**

| 型の定義 | 任意に追加するナビゲーション用の属性の値。いずれか一つの属性を選択する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:stringValue     [ UserDefinedValue ] | xs::string [0..1] | 文字列。 |
| uro:intValue     [ UserDefinedValue ] | xs::integer [0..1] | 整数。 |
| uro:doubleValue     [ UserDefinedValue ] | xs::double [0..1] | 実数。 |
| uro:codeValue     [ UserDefinedValue ] | gml::CodeType [0..1] | コード。コードリスト（UserDefinedValue_codeValue`[番号]`.xml）から選択する。`[番号]`は任意の半角数字とする。この属性を使用する場合は、コードリストを作成する。 |
| uro:dateValue     [ UserDefinedValue ] | xs::date [0..1] | 日付。 |
| uro:uriValue     [ UserDefinedValue ] | xs::anyURI [0..1] | URI。 |
| uro:measuredValue     [ UserDefinedValue ] | gml::MeasureType [0..1] | 単位付き数値。 |

