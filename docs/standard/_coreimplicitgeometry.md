# 9.8.5.1 core:ImplicitGeometry

## 9.8.5.1 core:ImplicitGeometry

**表 9-11**

| 型の定義 | 繰り返しオブジェクト。 地物毎に幾何オブジェクトを作成する代替として、一つのプロトタイプモデルを複数の地物が参照する仕組み。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << Type >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| core:libraryObject     [ ImplicitGeometry ] | xs::anyURI [0..1] | 繰り返しオブジェクトで使用するプロトタイプモデルの所在を示すURI。 この属性が記述されていない場合、core:relativeGMLGeometryを必須とする。 |
| core:mimeType     [ ImplicitGeometry ] | gml::CodeType [0..1] | 繰り返しオブジェクトで使用するプロトタイプモデルのファイル種類。コードリスト（ImplicitGeometry_mimeType.xml）より選択する。 |
| core:transformationMatrix     [ ImplicitGeometry ] | TransformationMatrix4x4 [0..1] | 繰り返しオブジェクトで使用するプロトタイプモデルの変形パラメータ。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| core:referencePoint             [ ImplicitGeometry ] | gml:Point [1..1] | 繰り返しオブジェクトの原点（0,0,0）を配置する参照点。3D都市モデルに適用される測地座標で記述する。 |
| core:relativeGMLGeometry             [ ImplicitGeometry ] | gml:_Geometry [0..1] | 繰り返しオブジェクトで使用するプロトタイプモデル。GML形式で記述する場合に必須とする。 この関連役割が記述されていない場合、core:libraryObjectを必須とする。 |

