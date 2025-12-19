# 4.21.3.1.8 gen:GenericCityObject

## 4.21.3.1.8 gen:GenericCityObject

**表 4-832**

| 型の定義 | CityGMLに定義されていない地物を定義するための汎用的な地物型。属性gml:nameは、汎用都市オブジェクトを識別する名称とする。コードリスト（GenericCityObject_name.xml）より選択する。 汎用都市オブジェクトを使用する場合は、コードリストを作成すること。 | ← |
| --- | --- | --- |
| 上位の型 | core:_CityObject | ← |
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
| gen:class     [ GenericCityObject ] | gml::CodeType [0..1] | 汎用都市オブジェクトの区分。コードリストより選択する。この属性を使用する場合は、コードリストを作成すること。 |
| gen:function     [ GenericCityObject ] | gml::CodeType [0..*] | 汎用都市オブジェクトの機能。コードリストより選択する。この属性を使用する場合は、コードリストを作成すること。 |
| gen:usage     [ GenericCityObject ] | gml::CodeType [0..*] | 汎用都市オブジェクトの用途。コードリストより選択する。この属性を使用する場合は、コードリストを作成すること。 |
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
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| gen:lod0Geometry             [ GenericCityObject ] | gml:_Geometry [0..1] | LOD0の形状。点、線又は平面により記述する。 |
| gen:lod0ImplicitRepresentation             [ GenericCityObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD0）への参照。標準製品仕様書では使用しない。 |
| gen:lod1Geometry             [ GenericCityObject ] | gml:_Geometry [0..1] | LOD1の形状。平面又は平面に一律の高さを与えた立体とする。 |
| gen:lod1ImplicitRepresentation             [ GenericCityObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD1）への参照。標準製品仕様書では使用しない。 |
| gen:lod2Geometry             [ GenericCityObject ] | gml:_Geometry [0..1] | LOD2の形状。主要な構造を単純な立体図形（球、円柱、円柱、角柱等）の組み合わせにより記述する。 |
| gen:lod2ImplicitRepresentation             [ GenericCityObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD2）への参照。標準製品仕様書では使用しない。 |
| gen:lod3Geometry             [ GenericCityObject ] | gml:_Geometry [0..1] | LOD3の形状。主構造及び詳細構造の外形により構成する。 |
| gen:lod3ImplicitRepresentation             [ GenericCityObject ] | core:ImplicitGeometry [0..1] | 繰り返しオブジェクト（LOD3）への参照。標準製品仕様書では使用しない。 |
| gen:lod4Geometry             [ GenericCityObject ] | gml:_Geometry [0..1] | LOD4の形状。主構造及び詳細構造の外形及びその内部形状により構成する。 |

