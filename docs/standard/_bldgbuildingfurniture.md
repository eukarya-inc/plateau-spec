# 4.2.3.1.17 bldg:BuildingFurniture

## 4.2.3.1.17 bldg:BuildingFurniture

**表 4-44**

| 型の定義 | 室内の移動できる備品（家具）。 bldg:IntBuildingInstallationが、建築物内部に設置された恒久的かつ固定的な設備であることと対照的に、bldg:BuildingFurnitureは椅子やテーブルのような、動かすことができる備品である。![Image](images/_bldgbuildingfurniture_img001.png)bldg:BuildingFurnitureの例（机、椅子）LOD4.2の場合にのみ取得する。ただし、ユースケースの要求に応じて、取得対象とする家具を限定してよい。 | ← |
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
| bldg:class     [ BuildingFurniture ] | gml::CodeType [0..1] | 家具の形態による区分。コードリスト（BuildingFurniture_class.xml）より選択する。 |
| bldg:function     [ BuildingFurniture ] | gml::CodeType [0..*] | 家具の主たる働き。コードリスト（BuildingFurniture_function.xml）より選択する。 |
| bldg:usage     [ BuildingFurniture ] | gml::CodeType [0..*] | 家具の主な使い道。標準製品仕様書では使用しない。 |
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
| bldg:lod4Geometry             [ BuildingFurniture ] | gml:_Geometry [0..1] | 家具の形状。家具の主要な構造について、それぞれの外形を構成する特徴点により作成した立体を平面に分割した面の集まりとして、表現する。gml:MultiSurfaceにより記述することを基本とする。容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。gml:MultiSurfaceにより記述することを基本とする。 |
| uro:ifcBuildingFurnitureAttribute             [ BuildingFurniture ] | uro:IfcAttribute [0..*] | IDM・MVDで定義されるIFCに含まれる情報。bldg:BuildingFurnitureに付与可能なデータ型は以下とする。uro:IfcFurnishingElement |
| uro:indoorFurnitureAttribute             [ BuildingFurniture ] | uro:IndoorAttribute [0..*] | 屋内ナビゲーションに必要な情報。bldg:BuildingFurnitureに付与可能なデータ型は以下とする。 uro:IndoorPublicTagAttribute  uro:IndoorZoneAttribute  uro:IndoorUserDefinedAttribute |

