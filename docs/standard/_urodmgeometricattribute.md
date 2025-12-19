# 4.25.3.1.1 uro:DmGeometricAttribute

## 4.25.3.1.1 uro:DmGeometricAttribute

**表 4-859**

| 型の定義 | 都市計画基本図として、地物の形状を公共測量標準図式に従って表現するためのデータ型。 uro:DmGeometrictAttributeは、地物の実体を表す図形だけではなく、地物を図式に従って表現する際に必要な情報（例：建物記号、建物の棟割線、記号の方向、指示点）を含む。 地物（ステレオタイプがFeatureTypeとなるクラス）は、関連役割uro:dmAttributeにより、このuro:DmGeometrictAttributeを保持できる。 | ← |
| --- | --- | --- |
| 上位の型 | uro:DmAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:dmCode     [ DmAttribute ] | gml::CodeType [1..1] | DMの図式分類コード。レイヤ番号（2桁）とデータ項目（2桁）からなる4桁の半角数字の列。コードリスト（ Common_dmCode.xml）より選択する。 |
| uro:meshCode     [ DmAttribute ] | gml::CodeType [0..1] | 数値地形図データが含まれる国土基本図の図郭識別番号。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:geometryType     [ DmGeometricAttribute ] | gml::CodeType [1..1] | レコードタイプ。コードリスト（ Common_geometryType.xml）より選択する。 |
| uro:mapLevel     [ DmGeometricAttribute ] | gml::CodeType [1..1] | 地図情報レベル。コードリスト（ Common_MapLevel.xml）より選択する。都市計画基本図の場合は、2500となる。 |
| uro:shapeType     [ DmGeometricAttribute ] | gml::CodeType [1..1] | 図形区分。コードリスト（ Common_shapeType.xml）より選択する。 |
| uro:visibility     [ DmGeometricAttribute ] | xs::boolean [0..1] | 可視性。上空から見た場合に、他の地物に遮蔽されておらず、上空から見えている（地図上に表現される）ことを示すフラグ。仮想的に設定された付属図形で描画対象としない場合にもこのフラグを0にする。 1 ：他の地物に遮蔽されていない。（省略時値）  0 ：他の地物に遮蔽されている。 |
| uro:is3d     [ DmGeometricAttribute ] | xs::boolean [0..1] | この図形の座標値が３次元データであることを示すフラグ。 1 ：この図形の座標値は３次元データである。  0 ：この図形の座標値は２次元データである。（省略時値）  uro:is3d=0の場合、z値に示された”0”は意味を持たない。 |
| uro:isInstallation     [ DmGeometricAttribute ] | xs::boolean [0..1] | この図形が地物の付属図形であることを示すフラグ。付属図形とは、地物の本質的な実体ではないが、描画処理などで地物を図形表現する際に利用可能な図形である。  1 ：この図形は付属図形である。  0 ：この図形は付属図形ではない。（省略時値）  uro:isInstallationの値が1の場合、uro:dmShapeTypeの値は、0以外となる。 |
| uro:isEdited     [ DmGeometricAttribute ] | xs::boolean [0..1] | 個別の編集処理がおこなわれたことを示すフラグ。 1：編集処理が行われた。  0：編集処理が行われていない。（省略時値） |
| uro:isSupplementarySymbol     [ DmGeometricAttribute ] | xs::boolean [0..1] | この図形が地物の補助記号であることを示すフラグ。補助記号とは、公共測量標準図式において自動発生が可能とされる図形を指す。自動発生が不可能な場合で図形を作成する必要がある場合にのみ使用する。  1 ：この図形は補助記号である。  0 ：この図形は補助記号ではない。（省略時値） |
| uro:angle     [ DmGeometricAttribute ] | xs::double [0..1] | 図形の角度。真北を0とし、時計まわりを正とする。 uro:dmGeometryTypeの値がE7（方向）の場合に多重度は任意となっているが、必須とする。 |
| uro:elevation     [ DmGeometricAttribute ] | gml::LengthType [0..1] | この図形の標高。単位はmとする。 uro:dmCodeのレイヤ番号が73の場合は、必須とする。 |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:dmElement             [ DmAttribute ] | uro:DmElement [0..1] | 数値地形図データファイル仕様に基づく要素レコードの情報。数値地形図データファイルの要素レコード情報を保持したい場合に必須とする。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:lod0Geometry             [ DmGeometricAttribute ] | gml:_Geometry [0..1] | 地物の形状を示す幾何オブジェクト。参照先の幾何オブジェクトの型は、uro:dmCode の値に応じて、gml:MultiPoint、gml:MultiCurve又はgml:MultiSurfaceのいずれかとする。 いずれの幾何オブジェクトの型となるかは、数値地形図の取得方法に従う。 |

