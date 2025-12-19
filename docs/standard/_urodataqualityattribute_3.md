# 4.4.3.2.2 uro:DataQualityAttribute

## 4.4.3.2.2 uro:DataQualityAttribute

**表 4-197**

| 型の定義 | 都市オブジェクトの品質を記述するためのデータ型。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:geometrySrcDescLod0     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD0の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択する。拡張製品仕様書でLOD0の幾何オブジェクトが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD0の幾何オブジェクトを含んでいない場合でも、「未作成」を示すコード「999」を選択すること（例えば、建築物モデルについて、一部の範囲のみLOD0の幾何オブジェクトが作成され、対象とする都市オブジェクトにはLOD1の幾何オブジェクトのみが含まれているような場合でも、その都市オブジェクトに関する本属性の値は「999」となる。）。 |
| uro:geometrySrcDescLod1     [ DataQualityAttribute ] | gml::CodeType [1..*] | LOD1の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択する。具体的な都市オブジェクトがLOD1の幾何オブジェクトを含んでいない場合でも、「未作成」を示すコード「999」を選択すること。 |
| uro:geometrySrcDescLod2     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD2の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択する。拡張製品仕様書でLOD2の幾何オブジェクトが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD2の幾何オブジェクトを含んでいない場合でも、「未作成」を示すコード「999」を選択すること（例えば、建築物モデルについて、一部の範囲のみLOD0の幾何オブジェクトが作成され、対象とする都市オブジェクトにはLOD1の幾何オブジェクトのみが含まれているような場合でも、その都市オブジェクトに関する本属性の値は「999」となる。）。 |
| uro:geometrySrcDescLod3     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD3の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択する。拡張製品仕様書でLOD3の幾何オブジェクトが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD3の幾何オブジェクトを含んでいない場合でも、「未作成」を示すコード「999」を選択すること（例えば、建築物モデルについて、一部の範囲のみLOD0の幾何オブジェクトが作成され、対象とする都市オブジェクトにはLOD1の幾何オブジェクトのみが含まれているような場合でも、その都市オブジェクトに関する本属性の値は「999」となる。）。 |
| uro:geometrySrcDescLod4     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD4の幾何オブジェクトの作成に使用した原典資料の種類。コードリスト（ DataQualityAttribute_geometrySrcDesc.xml）より選択する。拡張製品仕様書でLOD4の幾何オブジェクトが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD4の幾何オブジェクトを含んでいない場合でも、「未作成」を示すコード「999」を選択すること（例えば、建築物モデルについて、一部の範囲のみLOD0の幾何オブジェクトが作成され、対象とする都市オブジェクトにはLOD1の幾何オブジェクトのみが含まれているような場合でも、その都市オブジェクトに関する本属性の値は「999」となる。）。 |
| uro:thematicSrcDesc     [ DataQualityAttribute ] | gml::CodeType [0..*] | 主題属性の作成に使用した原典資料の種類。コードリスト（ DataQualityAttribute_thematicSrcDesc.xml）より選択する。主題属性が作成対象となっている場合は必須とする。 |
| uro:appearanceSrcDescLod0     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD0の幾何オブジェクトのアピアランスに使用した原典資料の種類。コードリスト（ DataQualityAttribute_appearanceSrcDesc.xml）より選択する。拡張製品仕様書でLOD0の幾何オブジェクトのアピアランスが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD0の幾何オブジェクトのアピアランスを含んでいない場合でも、「未作成」を示すコード「999」を選択すること。 |
| uro:appearanceSrcDescLod1     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD1の幾何オブジェクトのアピアランスに使用した原典資料の種類。コードリスト（ DataQualityAttribute_appearanceSrcDesc.xml）より選択する。拡張製品仕様書LOD1の幾何オブジェクトのアピアランスが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD1の幾何オブジェクトのアピアランスを含んでいない場合でも、「未作成」を示すコード「999」を選択すること。 |
| uro:appearanceSrcDescLod2     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD2の幾何オブジェクトのアピアランスに使用した原典資料の種類。コードリスト（ DataQualityAttribute_appearanceSrcDesc.xml）より選択する。拡張製品仕様書でLOD2の幾何オブジェクトのアピアランスが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD2の幾何オブジェクトのアピアランスを含んでいない場合でも、「未作成」を示すコード「999」を選択すること。 |
| uro:appearanceSrcDescLod3     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD3の幾何オブジェクトのアピアランスに使用した原典資料の種類。コードリスト（ DataQualityAttribute_appearanceSrcDesc.xml）より選択する。拡張製品仕様書でLOD3の幾何オブジェクトのアピアランスが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD3の幾何オブジェクトのアピアランスを含んでいない場合でも、「未作成」を示すコード「999」を選択すること。 |
| uro:appearanceSrcDescLod4     [ DataQualityAttribute ] | gml::CodeType [0..*] | LOD4の幾何オブジェクトのアピアランスに使用した原典資料の種類。コードリスト（ DataQualityAttribute_appearanceSrcDesc.xml）より選択する。拡張製品仕様書でLOD4の幾何オブジェクトのアピアランスが作成対象となっている場合は必須とする。この場合、具体的な都市オブジェクトがLOD4の幾何オブジェクトのアピアランスを含んでいない場合でも、「未作成」を示すコード「999」を選択すること。 |
| uro:lodType     [ DataQualityAttribute ] | gml::CodeType [0..*] | 幾何オブジェクトに適用されたLODの詳細な区分。建築物の場合コードリスト（ Building_lodType.xml）より選択する。 LOD2以上の幾何オブジェクトを作成する場合は必須とする。道路、徒歩道又は広場の場合コードリスト（ Road_lodType.xml）より選択する。 LOD3の幾何オブジェクトを作成する場合は必須とする。鉄道の場合コードリスト（ Railway_lodType.xml）より選択する。 LOD3の幾何オブジェクトを作成する場合は必須とする。LODの細分が定義されていない地物型の場合この属性を使用しない。 |
| uro:lod1HeightType     [ DataQualityAttribute ] | gml::CodeType [0..1] | LOD1の立体図形を作成する際に使用した高さの算出方法。コードリスト（DataQualityAttribute_lod1HeightType.xml）より選択する。立体の幾何オブジェクトを作成する場合は必須とする。 |
| uro:tranDataAcquisition     [ DataQualityAttribute ] | xs::string [0..1] | [道路基盤地図情報（整備促進版）製品仕様書（案）（平成27年5月）](https://www.nilim.go.jp/lab/bcg/siryou/tnn/tnn0848pdf/ks084811.pdf)に定める「取得レベル(level)」を記述するための属性。道路の場合にのみ使用する。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:publicSurveyDataQualityAttribute             [ DataQualityAttribute ] | uro:PublicSurveyDataQualityAttribute [0..1] | 使用した公共測量成果の地図情報レベルと種類。各LODの幾何オブジェクトの作成に使用した原典資料の種類に関する属性（uro:geometrySrcDescLod0等）のコード値（コードリスト（DataQualityAttribute_geometrySrcDesc.xml）より選択される）が公共測量成果（コード「000」）となっている場合は、必須とする。 |

