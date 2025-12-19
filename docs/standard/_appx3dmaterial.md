# 4.22.3.1.4 app:X3DMaterial

## 4.22.3.1.4 app:X3DMaterial

**表 4-836**

| 型の定義 | 地物インスタンスの表面色を指定するための型。![Image](images/_appx3dmaterial_img001.png)bldg:Buildingの外観として色を指定した例 | ← |
| --- | --- | --- |
| 上位の型 | app:_SurfaceData | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| app:isFront     [ _SurfaceData ] | xs::boolean [0..1] | テクスチャを貼る面の指定。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| app:ambientIntensity     [ X3DMaterial ] | xs::double [0..1] | 環境光（光源からの光の当たっていない部分の明るさ）に対する反射率を指定。これが 0 だと陰の部分が真っ暗になる。 |
| app:diffuseColor     [ X3DMaterial ] | gml::doubleList [0..1] | 拡散反射率を指定（rgb、いずれも 0～1）。これが物体の色になる。 |
| app:emissiveColor     [ X3DMaterial ] | gml::doubleList [0..1] | 発光（rgb、いずれも 0～1）を指定。 0 0 0 より大きくすると、この物体自体がその色で発光する。 |
| app:specularColor     [ X3DMaterial ] | gml::doubleList [0..1] | 鏡面反射率を指定（rgb、いずれも 0～1）で指定。これは光源の光が物体表面でそのまま反射して見える部分（ハイライト）の色になる。 |
| app:shininess     [ X3DMaterial ] | xs::double [0..1] | ハイライトの「強さ」を、 0～１ の間で指定。この値が大きくなるにつれてハイライトが鋭くなり、輝き感が増す。 |
| app:transparency     [ X3DMaterial ] | xs::double [0..1] | 透明度を指定。 0 で不透明になり、1 で透明になる。デフォルトは 0。 |
| app:isSmooth     [ X3DMaterial ] | core::doubleBetween0and1 [0..1] | 陰影付のための補間方法を指定。trueの場合、グーロー法(による陰影付け)となる。 |
| app:target     [ X3DMaterial ] | xs::anyURI [0..*] | 色を設定する幾何への参照。 |

