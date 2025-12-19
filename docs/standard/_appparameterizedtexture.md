# 4.22.3.1.2 app:ParameterizedTexture

## 4.22.3.1.2 app:ParameterizedTexture

**表 4-834**

| 型の定義 | 地物インスタンスに貼り付ける画像と貼り付け方をパラメータにより指定するための型。![Image](images/_appparameterizedtexture_img001.png)bldg:Buildingの外観としてテクスチャを指定した例 | ← |
| --- | --- | --- |
| 上位の型 | app:_Texture | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| app:isFront     [ _SurfaceData ] | xs::boolean [0..1] | テクスチャを貼る面の指定。 |
| app:imageURI     [ _Texture ] | xs::anyURI [1..1] | テクスチャの画像への参照URI。相対パスにより記述する。 |
| app:mimeType     [ _Texture ] | gml::CodeType [0..1] | テクスチャの画像フォーマット。コードリスト（Appearance_mimeType.xml）より選択する。画像フォーマットは、PNG又はJPEGとする。 |
| app:textureType     [ _Texture ] | TextureTypeType [0..1] | テクスチャの種類。 |
| app:wrapMode     [ _Texture ] | WrapModeType [0..1] | テクスチャの貼り付け方。 borderを選択した場合は、borderColorも指定すること。 |
| app:borderColor     [ _Texture ] | ColorPlusOpacity [0..1] | テクスチャの端の色表現の指定。RGBに加え、不透過率を0から1までの実数値で指定する。不透過率が0の時に透明となる。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| app:target             [ ParameterizedTexture ] | app:_TextureParameterization [0..*] | テクスチャの貼り付けに使用する幾何への参照。 app:TexCoordListを使用する。 |

