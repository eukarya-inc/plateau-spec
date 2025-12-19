# L.7 brid:BridgeFurniture

## L.7 brid:BridgeFurniture

brid:BridgeFurnitureは、橋梁内部の空間に配置された、可動設備を記述するために用いる。brid:IntBridgeInstallationが、橋梁内部に設置された恒久的かつ固定的な設備であることと対照的に、brid:BridgeFurnitureは椅子や机のような動かすことができる（位置が固定されない）設備である。

ユースケースの要求に応じて、取得対象とする可動設備を決定してよい。

## L.7.1 可動設備の空間属性

brid:BridgeFurnitureの空間属性の型は、gml:_Geometryである。gml:_Geometryは、幾何オブジェクトの基底要素であり、実装においては、この下位型のうち具象となる幾何オブジェクトを地物の形状に合わせて選定する。

**表 L-15 — L-10: brid:BridgeFurnitureの空間属性タイプ**

| ID | `/req/brid/10` |
| --- | --- |
| 主題 | 妥当な橋梁データセットの要件 |
| 分類 | 要件分類L-1: 妥当な橋梁オブジェクト |
| 条文 | brid:BridgeFurnitureの空間属性の型には、gml:MultiSurfaceを使用することを原則とする。 |

CityGMLでは、曲面や立体だけではなく、点や曲線も使用可能である。しかしながら3D都市モデルでは、3次元での利用を想定し、付属物の形状に使用可能な幾何オブジェクトを曲面又は立体に限定する。原則として面の集まり（gml:MultiSurface）を使用するが、ユースケースの必要に応じてgml:Solidを使用できる。

