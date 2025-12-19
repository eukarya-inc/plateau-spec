# 4.17.3.1.2 veg:PlantCover

## 4.17.3.1.2 veg:PlantCover

**表 4-754**

| 型の定義 | 植被。芝生や茂みのように、植生に覆われた範囲を指し、個々の樹木を識別しない。![Image](images/_vegplantcover_img001.png)veg:PlantCoverの例 | ← |
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
| veg:class     [ PlantCover ] | gml::CodeType [0..1] | 植被を構成する主な樹種による分類。コードリスト（PlantCover_class.xml）より選択する。 |
| veg:function     [ PlantCover ] | gml::CodeType [0..*] | 植被の機能。標準製品仕様書では使用しない。 |
| veg:usage     [ PlantCover ] | gml::CodeType [0..*] | 植被の用途。標準製品仕様書では使用しない。 |
| veg:averageHeight     [ PlantCover ] | gml::LengthType [0..1] | 平均高さ。単位はm（uom=”m”）とする。 |
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
| veg:lod1MultiSolid             [ PlantCover ] | gml:MultiSolid [0..1] | 植被の形状をそれが占有している範囲（面）に一律の高さを与えた立体として表現する。植被の占有している範囲（面）に、その範囲内で中央値となる植被の高さで上向きに押し出した立体となる。![Image](images/_vegplantcover_img002.png)LOD1のPlantCoverの例 |
| veg:lod1MultiSurface             [ PlantCover ] | gml:MultiSurface [0..1] | LOD1はSolidにより表現するため、使用しない。 |
| veg:lod2MultiSolid             [ PlantCover ] | gml:MultiSolid [0..1] | 植被の形状を立体とし、その主要な部分の外形を立体として詳細に区分して表現する。植被の範囲内で比高3m 以上の場合にこれを区分した面を境界とする立体となる。![Image](images/_vegplantcover_img003.png)LOD2のPlantCoverの例 |
| veg:lod2MultiSurface             [ PlantCover ] | gml:MultiSurface [0..1] | 植被の形状を立体とし、その主要な部分の外形を面の集まり又は立体として区分して表現する。植被の範囲内で比高3m以上の場合にこれを区分した面の集まりとなる。![Image](images/_vegplantcover_img004.png)LOD2のPlantCoverの例 |
| veg:lod3MultiSolid             [ PlantCover ] | gml:MultiSolid [0..1] | 植被の形状を立体とし、その主要な部分の外形を立体として詳細に区分して表現する。植被の範囲内で比高1m 以上の場合にこれを区分した面を境界とする立体となる。![Image](images/_vegplantcover_img005.png)LOD3のPlantCoverの例 |
| veg:lod3MultiSurface             [ PlantCover ] | gml:MultiSurface [0..1] | 植被の形状を立体とし、その主要な部分の外形を面の集まりとして詳細に区分して表現する。植被の範囲内で比高1m以上の場合にこれを区分した面の集まりとなる。![Image](images/_vegplantcover_img006.png)LOD3のPlantCoverの例 |
| veg:lod4MultiSolid             [ PlantCover ] | gml:MultiSolid [0..1] | LOD4の幾何形状を立体により表現する。標準製品仕様書では使用しない。 |
| veg:lod4MultiSurface             [ PlantCover ] | gml:MultiSurface [0..1] | LOD4の幾何形状を面の集まりにより表現する。標準製品仕様書では使用しない。 |

