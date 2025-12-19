# 4.2.3.5.11 uro:IfcGeometricRepresentationContext

## 4.2.3.5.11 uro:IfcGeometricRepresentationContext

**表 4-68**

| 型の定義 | プロジェクト内のIfcProduct オブジェクトの形状表現に適用されるコンテキストを定義する、3D のモデル表現形式のためのデータ型。コンテキスト情報とは、形状表現が定義されるコンテキストのタイプと、このコンテキストで定義される形状表現項目に適用される数値精度を定義、さらに、uro:worldCoordinateSystem 属性を使用して、グローバルな原点からプロジェクト座標系をオフセットする情報となる。uro:worldCoordinateSystem のy 軸が真北を指していない場合、uro:trueNorth 属性を指定することができる。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:contextIdentifier     [ IfcGeometricRepresentationContext ] | xs::string [0..1] | 識別子。 |
| uro:contextType     [ IfcGeometricRepresentationContext ] | xs::string [0..1] | “Model”とする。 |
| uro:coordinateSpaceDimension     [ IfcGeometricRepresentationContext ] | xs::integer [0..1] | 次元数。3とする。 |
| uro:precision     [ IfcGeometricRepresentationContext ] | xs::double [0..1] | 精度。通常は、1E-5 から 1E-8 の値を設定する。 |
| uro:worldCoordinateSystem     [ IfcGeometricRepresentationContext ] | IfcAxis2Placement3D [1..1] | プロジェクトで使用される全ての表現コンテキストのエンジニアリング座標系。 |
| uro:trueNorth     [ IfcGeometricRepresentationContext ] | gml::doubleList [0..1] | 北方角との差を2 次元ベクトルで設定する。角度表現のラジアン又は度の設定は、MVD-IfcProject.UnitsInContext（短径設定情報）を参照。北が0 時の方向であれば値は(0,1)。![Image](images/_uroifcgeometricrepresentationcontext_img001.png) |

