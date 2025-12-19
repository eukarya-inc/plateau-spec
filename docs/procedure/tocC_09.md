# C.9 bldg:IntBuildingInstallation

## C.9 bldg:IntBuildingInstallation

bldg:IntBuildingInstallationは、階段、スロープ、てすり、輸送設備（エレベータ、エスカレータ、動く歩道）のような、建築物の内側に設置された付属的な設備であり、かつ永続的に設置されている固定的な設備である。

## C.9.1 建築物屋内付属物の空間属性

bldg:BuildingInstallationの空間属性bldg:lod4Geomertyの型は、gml:_Geometryである。gml:_Geometryは、幾何オブジェクトの基底要素であり、実装においては、この下位型のうち、具象となる幾何オブジェクトを、地物の形状に合わせて選定する。

**表 C-43 — C-30: bldg:BuildingInstallationの空間属性タイプ**

| ID | `/req/bldg/30` |
| --- | --- |
| 主題 | 妥当な建築物データセットの要件 |
| 分類 | 要件分類C-1: 妥当な建築物オブジェクト作成 |
| 条文 | bldg:BuildingInstallationの空間属性の型には、gml:MultiSurfaceを使用することを原則とする。 |

ある程度の広がりをもつ地物は、曲面の集まり又は立体を使用する。CityGMLでは、曲面や立体だけではなく、点や曲線も使用可能である。しかしながら3D都市モデルでは、3次元での利用を想定し、屋内付属物の形状に使用可能な幾何オブジェクトを曲面又は立体に限定する。原則として面の集まり（gml:MultiSurface）を使用するが、ユースケースの必要に応じてgml:Solidを使用できる。

## C.9.2 建築物屋内付属物の主題属性

bldg:IntBuildingInstallationは、建築物の内部に付属する恒久的に設置された設備の総称となる。そのため、それが何であるか（例：階段、エレベータ）は、主題属性により区分する。

**表 C-44 — C-31: bldg:IntBuildingInstallationの設備区分方法**

| ID | `/req/bldg/31` |
| --- | --- |
| 主題 | 妥当な建築物データセットの要件 |
| 分類 | 要件分類C-1: 妥当な建築物オブジェクト作成 |
| 条文 | bldg:IntBuildingInstallationは、bldg:class及びbldg:functionによりその設備を区分する。 |

bldg:class及びbldg:functionが取りうる値は、コードリストにより定義されている。

標準製品仕様書では、bldg:classの定義域として表C-45に示す区分を定義している。この定義域は、BIMモデルの国際標準であるIFCでのクラスの区分に基づく。建築物モデル（LOD4）の場合、BIMモデルからの変換により作成されるケースが多いと考えられるためである。また、bldg:functionの定義域は、Uniclass2015に従う。Uniclassは、建築生産で必要とする建築部位、工法及び設備などの名称を体系的に整理したものである。

**表 C-45 — IntBuildingInstallation bldg:classの区分**

| ファイル名 | IntBuildingInstallation_class.xml |
| --- | --- |
| ファイルURL | [https://www.geospatial.jp/iur/codelists/3.2/IntBuildingInstallation_class.xml](https://www.geospatial.jp/iur/codelists/3.2/IntBuildingInstallation_class.xml) |
| コード | 説明 |
| BE_01 | IfcBeam |
| BE_02 | IfcColumn |
| BE_05 | IfcPlate |
| BE_06 | IfcRailing |
| BE_07 | IfcRamp |
| BE_08 | IfcRampFlight |
| BE_11 | IfcStair |
| BE_12 | IfcStairFlight |
| BE_16 | IfcBuildingElementProxy |
| BE_17 | IfcTransportElement |
| 出典： [IFC2x3 CV2.0](https://standards.buildingsmart.org/IFC/RELEASE/IFC2x3/FINAL/HTML/) | ← |

