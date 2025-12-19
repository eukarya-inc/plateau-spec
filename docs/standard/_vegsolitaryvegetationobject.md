# 4.17.3.1.1 veg:SolitaryVegetationObject

## 4.17.3.1.1 veg:SolitaryVegetationObject

**表 4-753**

| 型の定義 | 独立した樹木。![Image](images/_vegsolitaryvegetationobject_img001.png)SolitaryVegetationObjectの例 | ← |
| --- | --- | --- |
| 上位の型 | veg:_VegetationObject | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |
| core:creationDate     [ _CityObject ] | xs::date [0..1] | データが作成された日。運用上必須とする。 |
| core:terminationDate     [ _CityObject ] | xs::date [0..1] | データが削除された日。 |
| core:relativeToTerrain     [ _CityObject ] | core::RelativeToTerrainType [0..1] | 地表面との相対的な位置関係。標準製品仕様書では使用しない。 |
| core:relativeToWater     [ _CityObject ] | core::RelativeToWaterType [0..1] | 水面との相対的な位置関係。標準製品仕様書では使用しない。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| veg:class     [ SolitaryVegetationObject ] | gml::CodeType [0..1] | 高木、中木、低木の別。コードリスト（SolitaryVegetationObject_class.xml）より選択する。 |
| veg:function     [ SolitaryVegetationObject ] | gml::CodeType [0..*] | 常緑又は落葉の区分及び針葉又は広葉の区分。コードリスト（SolitaryVegetationObject_function.xml）より選択する。 |
| veg:usage     [ SolitaryVegetationObject ] | gml::CodeType [0..*] | 樹木の用途。標準製品仕様書では使用しない。 |
| veg:species     [ SolitaryVegetationObject ] | gml::CodeType [0..1] | 樹木の樹種。標準製品仕様書では使用しない。 |
| veg:height     [ SolitaryVegetationObject ] | gml::LengthType [0..1] | 樹高。樹木の樹冠の頂端から根鉢の上端までの垂直高をいい、一部の突出した枝は含まない。単位はm（uom=”m”）とする。 |
| veg:trunkDiameter     [ SolitaryVegetationObject ] | gml::LengthType [0..1] | 樹径。幹周を3.14で除算した数値。なお、幹周とは樹木の幹の周長をいい、根鉢の上端より1.2m上りの位置を測定する。この部分に枝が分岐しているときは、その上部を測定する。幹が二本以上の樹木の場合においては、各々の周長の総和の70％をもって幹周とする。単位はm（uom=”m”）とする。 |
| veg:crownDiameter     [ SolitaryVegetationObject ] | gml::LengthType [0..1] | 樹木の四方面に伸長した枝(葉)の幅をいう。測定方向により幅に長短がある場合は、最長と最短の平均値とする。なお一部の突出した枝は含まない。単位はm（uom=”m”）とする。 |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| core:externalReference             [ _CityObject ] | core:ExternalReference [0..*] | 外部への参照。標準製品仕様書では使用しない。 |
| gen:dateAttribute             [ _CityObject ] | gen:dateAttribute [0..*] | 日付型属性。属性を追加したい場合に使用する。 |
| gen:doubleAttribute             [ _CityObject ] | gen:doubleAttribute [0..*] | 実数型属性。属性を追加したい場合に使用する。 |
| gen:genericAttributeSet             [ _CityObject ] | gen:genericAttributeSet [0..*] | 汎用属性のセット。属性を追加したい場合に使用する。 |
| gen:intAttribute             [ _CityObject ] | gen:intAttribute [0..*] | 整数型属性。属性を追加したい場合に使用する。 |
| gen:measureAttribute             [ _CityObject ] | gen:measureAttribute [0..*] | 単位付き数値型属性。属性を追加したい場合に使用する。 |
| gen:stringAttribute             [ _CityObject ] | gen:stringAttribute [0..*] | 文字列型属性。属性を追加したい場合に使用する。 |
| gen:uriAttribute             [ _CityObject ] | gen:uriAttribute [0..*] | URI型属性。属性を追加したい場合に使用する。 |
| uro:pointCloud             [ _CityObject ] | uro:AbstractPointCloud [0..1] | ポイントクラウドへの参照。 |
| uro:vegDataQualityAttribute             [ _VegetationObject ] | uro:DataQualityAttribute [1..1] | 作成されたデータの品質に関する情報。必須とする。関連役割uro:publicSurveyDataQualityAttributeは、属性uro:srcScaleLod2、uro:srcScaleLod3及びuro:publicSurveySrcDescLod2、uro:publicSurveySrcDescLod3を使用しない。 |
| uro:vegDmAttribute             [ _VegetationObject ] | uro:DmAttribute [0..*] | 公共測量標準図式による図形表現に必要な情報。 |
| uro:vegFacilityAttribute             [ _VegetationObject ] | uro:FacilityAttribute [0..*] | uro:vegFacilityTypeAttribute.class によって指定された分野における施設管理情報。 |
| uro:vegFacilityIdAttribute             [ _VegetationObject ] | uro:FacilityIdAttribute [0..1] | uro:vegFacilityTypeAttribute.class によって指定された分野における施設の識別情報。 |
| uro:vegFacilityTypeAttribute             [ _VegetationObject ] | uro:FacilityTypeAttribute [0..*] | 特定分野における施設の分類情報。 |
| uro:vegKeyValuePairAttribute             [ _VegetationObject ] | uro:KeyValuePairAttribute [0..*] | 属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttribute の下位型を使用する。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| veg:lod1Geometry             [ SolitaryVegetationObject ] | gml:_Geometry [0..1] | 単独木の形状をそれが占有している範囲（面）に一律の高さを与えた立体として表現する。樹冠の情報からの正射影の外周を樹高の高さで上向きに押し出した立体となる。![Image](images/_vegsolitaryvegetationobject_img002.png)LOD1のSolitaryVegetationObjectの例 |
| veg:lod1ImplicitRepresentation             [ SolitaryVegetationObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD1）への参照。標準製品仕様書では使用しない。 |
| veg:lod2Geometry             [ SolitaryVegetationObject ] | gml:_Geometry [0..1] | 単独木の形状を立体とし、その主要な部分の外形を面の集まり又は立体として区分して表現する。樹冠と樹幹をそれぞれ簡略化した立体（楕円体、球体、円錐、角錐、角柱、円柱などの単純な立体図形）を組み合わせた立体として表現する。![Image](images/_vegsolitaryvegetationobject_img003.png)LOD2のSolitaryVegetationObjectの例樹木量の算定等、容積が必要ではない場合には、gml:MultiSurfaceにより外形を構成する。 |
| veg:lod2ImplicitRepresentation             [ SolitaryVegetationObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD2）への参照。標準製品仕様書では使用しない。 |
| veg:lod3Geometry             [ SolitaryVegetationObject ] | gml:_Geometry [0..1] | 単独木の形状を立体とし、その主要な部分の外形を面の集まり又は立体として詳細に区分して表現する。一定高さごとに樹冠の横断面を作成し、この頂点を結び外形を構成する（樹冠内部の主枝等の表現は行わない）。![Image](images/_vegsolitaryvegetationobject_img004.png)LOD3のSolitaryVegetationObjectの例樹木量の算定等、容積が必要ではない場合には、gml:MultiSurfaceにより外形を構成する。 |
| veg:lod3ImplicitRepresentation             [ SolitaryVegetationObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD3）への参照。標準製品仕様書では使用しない。 |
| veg:lod4Geometry             [ SolitaryVegetationObject ] | gml:_Geometry [0..1] | LOD4の幾何形状。標準製品仕様書では使用しない。 |
| veg:lod4ImplicitRepresentation             [ SolitaryVegetationObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD4）への参照。標準製品仕様書では使用しない。 |

