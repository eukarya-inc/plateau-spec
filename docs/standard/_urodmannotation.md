# 4.25.3.1.3 uro:DmAnnotation

## 4.25.3.1.3 uro:DmAnnotation

**表 4-861**

| 型の定義 | 都市計画基本図として、注記を表現するために必要な情報のデータ型。地物（ステレオタイプがFeatureType）は、関連役割uro:dmAttributeにより、uro:DmAnnotationを保持できる。このとき、uro:DmAnnotationは地物に付属する情報となる。 | ← |
| --- | --- | --- |
| 上位の型 | uro:DmAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:dmCode     [ DmAttribute ] | gml::CodeType [1..1] | DMの図式分類コード。レイヤ番号（2桁）とデータ項目（2桁）からなる4桁の半角数字の列。コードリスト（ Common_dmCode.xml）より選択する。 |
| uro:meshCode     [ DmAttribute ] | gml::CodeType [0..1] | 数値地形図データが含まれる国土基本図の図郭識別番号。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:geometryType     [ DmAnnotation ] | gml::CodeType [1..1] | レコードタイプ。コードリスト（ Common_geometryType.xml）より選択する。注記の場合はE7となる。 |
| uro:shapeType     [ DmAnnotation ] | gml::CodeType [1..1] | 図形区分。コードリスト（ Common_shapeType.xml）より選択する。注記の場合は0となる。 |
| uro:label     [ DmAnnotation ] | xs::string [1..1] | 注記の文字列。 |
| uro:isVertical     [ DmAnnotation ] | xs::boolean [1..1] | 縦書きか否か。 1 ：縦書き 0 ：横書き |
| uro:size     [ DmAnnotation ] | xs::integer [1..1] | 字の大きさ。単位は10分の1ミリメートル。 |
| uro:orientation     [ DmAnnotation ] | xs::integer [1..1] | 注記の表示方向を示す角度。単位は度、範囲は縦書きの場合は－135度～－45度、横書きの場合は－45度～＋45度とする。 |
| uro:linewidth     [ DmAnnotation ] | xs::integer [1..1] | 注記の線の太さ。線号の号数を記述する。 |
| uro:spacing     [ DmAnnotation ] | xs::integer [1..1] | 字の間隔。単位は10分の1ミリメートル。全角・半角が混在する場合には、全角を基準とする。 |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:dmElement             [ DmAttribute ] | uro:DmElement [0..1] | 数値地形図データファイル仕様に基づく要素レコードの情報。数値地形図データファイルの要素レコード情報を保持したい場合に必須とする。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| uro:lod0anchorPoint             [ DmAnnotation ] | gml:_Geometry [1..1] | 注記を配置する位置。 点（gml:Point）を使用して記述する。 |

