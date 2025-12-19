# 4.15.3.1.15 uro:LengthAttribute

## 4.15.3.1.15 uro:LengthAttribute

**表 4-711**

| 型の定義 | 地下埋設物の実長及び亘長を表すデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:length     [ LengthAttribute ] | gml::LengthType [0..1] | 延長。単位はｍ。 |
| uro:measureType     [ LengthAttribute ] | gml::CodeType [0..1] | 延長の計測方法。コードリスト（LengthAttribute_measureType.xml）より選択する。 |
| uro:phaseType     [ LengthAttribute ] | gml::CodeType [0..1] | 延長が計測された段階。コードリスト（LengthAttribute_phaseType.xml）より選択する。 |

