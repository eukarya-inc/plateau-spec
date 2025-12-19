# 4.9.3.2.7 uro:KeyValuePairAttribute

## 4.9.3.2.7 uro:KeyValuePairAttribute

**表 4-320**

| 型の定義 | 都市オブジェクトに付与する追加情報。都市オブジェクトが継承する属性及び都市オブジェクトに定義された属性以外にコード型の属性を追加したい場合に使用する。属性名称と属性の値の対で構成される。コード値以外の属性を追加する場合は、gen:_GenericAttributeを使用すること。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:key     [ KeyValuePairAttribute ] | gml::CodeType [1..1] | 拡張する属性の名称。名称は、コ－ドリスト（KeyValuePairAttribute_key.xml）を作成し、選択する。 |
| uro:codeValue     [ KeyValuePairAttribute ] | gml::CodeType [1..1] | 拡張された属性の値。名称は、コ－ドリスト（KeyValuePairAttribute_key`[key]`.xml）を作成し、選択する。 `[key]`は、属性uro:keyの値に一致する。 |

