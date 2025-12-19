# 4.2.3.2.12 uro:PublicSurveyDataQualityAttribute

## 4.2.3.2.12 uro:PublicSurveyDataQualityAttribute

**表 4-56**

| 型の定義 | 都市オブジェクトのデータ作成情報を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:srcScaleLod0     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..1] | LOD0の幾何オブジェクトの作成に使用した原典資料の地図情報レベル。コードリスト（ PublicSurveyDataQualityAttribute_srcScale.xml）より選択する。 「LOD0の幾何オブジェクトの作成に使用した原典資料の種類に関する属性」（uro:geometrySrcDescLod0）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。 |
| uro:srcScaleLod1     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..1] | LOD1の幾何オブジェクトの作成に使用した原典資料の地図情報レベル。コードリスト（ PublicSurveyDataQualityAttribute_srcScale.xml）より選択する。 「LOD1の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod1）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。 |
| uro:srcScaleLod2     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..1] | LOD2の幾何オブジェクトの作成に使用した原典資料の地図情報レベル。コードリスト（ PublicSurveyDataQualityAttribute_srcScale.xml）より選択する。 「LOD2の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod2）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の地図情報レベルが混在する場合は、最も低い地図情報レベルを記載する。例えば、壁面は地図情報レベル1000、屋根面は地図情報レベル2500となる場合は、地図情報レベル2500とする。 |
| uro:srcScaleLod3     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..1] | LOD3の幾何オブジェクトの作成に使用した原典資料の地図情報レベル。コードリスト（ PublicSurveyDataQualityAttribute_srcScale.xml）より選択する。 「LOD3の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod3）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の地図情報レベルが混在する場合は、最も低い地図情報レベルを記載する。例えば、壁面は地図情報レベル1000、屋根面は地図情報レベル2500となる場合は、地図情報レベル2500とする。 |
| uro:srcScaleLod4     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..1] | LOD4の幾何オブジェクトの作成に使用した原典資料の地図情報レベル。コードリスト（ PublicSurveyDataQualityAttribute_srcScale.xml）より選択する。 「LOD4の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod4）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の地図情報レベルが混在する場合は、最も低い地図情報レベルを記載する。例えば、外側の形状は地図情報レベル1000、屋内の形状は地図情報レベル500となる場合は地図情報レベル1000とする。 |
| uro:publicSurveySrcDescLod0     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..*] | LOD0の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（PublicSurveyDataQualityAttribute_publicSurveySrcDesc.xml）より選択する。 「LOD0の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod0）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の種類の原典資料を使用した場合は、それぞれを記述する。 |
| uro:publicSurveySrcDescLod1     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..*] | LOD1の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（PublicSurveyDataQualityAttribute_publicSurveySrcDesc.xml）より選択する。 「LOD1の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod1）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の種類の原典資料を使用した場合は、それぞれを記述する。 |
| uro:publicSurveySrcDescLod2     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..*] | LOD2の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（PublicSurveyDataQualityAttribute_publicSurveySrcDesc.xml）より選択する。 「LOD2の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod2）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の種類の原典資料を使用した場合は、それぞれを記述する。 |
| uro:publicSurveySrcDescLod3     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..*] | LOD3の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（PublicSurveyDataQualityAttribute_publicSurveySrcDesc.xml）より選択する。 「LOD3の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod3）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の種類の原典資料を使用した場合は、それぞれを記述する。 |
| uro:publicSurveySrcDescLod4     [ PublicSurveyDataQualityAttribute ] | gml::CodeType [0..*] | LOD4の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（PublicSurveyDataQualityAttribute_publicSurveySrcDesc.xml）より選択する。 「LOD4の幾何オブジェクトの作成に使用した原典資料の種類についての属性」（uro:geometrySrcDescLod4）のコード値（コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）のみの場合は、必須とする。複数の種類の原典資料を使用した場合は、それぞれを記述する。 |

