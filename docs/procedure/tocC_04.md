# C.4 bldg:BuildingPart

## C.4 bldg:BuildingPart

bldg:BuildingPartは、一棟の建築物を複数の部分に分割して記述するために使用する地物型である。例えば、一棟の建築物の中に階数又は用途若しくは構造が異なる場所があり、それぞれを分けて記述したい場合に使用できる。

標準製品仕様では、bldg:BuildingPartはLOD2、LOD3及びLOD4において使用可能としている。

![Image](images/tocC_04_img001.png)

bldg:BuildingPartを使用して建築物を記述する場合、bldg:Buildingはbldg:BuildginPartを束ねるコンテナとして機能する。

**表 C-29 — C-18: bldg:BuildingPartの妥当性条件**

| ID | `/req/bldg/18` |
| --- | --- |
| 主題 | 妥当な建築物データセットの要件 |
| 分類 | 要件分類C-1: 妥当な建築物オブジェクト作成 |
| 条文 | 妥当なbldg:BuildingPartは、以下を満たさなければならない。 |
| A | 唯一のbldg:Buildingからbldg:consistsOfBuildingPartにより参照される。（複数のbldg:Buildingから参照されていてはならない） |
| B | 別のbldg:BuildingPartと接している。（一棟のbldg:Buildingは二つ以上のbldg:BuildingPartから構成され、それらは連続している） |

一棟の建築物を、建築物部分や建築物付属物に分けて表現する場合には、必ず、一体的な建築物として作成しなければならない。建築物部分をbldg:BuildingPartを用いて記述する場合は、一棟のbldg:Buildingを構成するbldg:BuildingPartが互いに接していなければならない。
 bldg:BuildingPartもbldg:Buildingと同様に、空間属性や主題属性をもつことができる。例えば、高層階と低層階を組み合わせた建築物において、都市計画基礎調査が、高層階・低層階それぞれで実施されている場合には、それぞれの調査結果をbldg:BuildingPartに記述する。

例えば、図C-21のように、一棟の建築物が商業施設、業務施設、共同住宅からなる複合施設であり、それぞれの用途に対して都市計画基礎調査が実施されていたとする。この場合、商業施設、業務施設、共同住宅をそれぞれbldg:BuildingPartとして記述し、bldg:BuildingPartごとに都市計画基礎調査で得られた情報を属性として付与できる。

このとき、それぞれのbldg:BuildingPartが幾何オブジェクト（gml:Solid）をもち、bldg:Buildingは、幾何オブジェクトをもたない。また、bldg:Buildingにはbldg:BuildingPartに共通となる属性を格納する。
ただし、bldg:usageのように複数の値を列挙できる属性に関して、一つでも異なる値がある場合には、当該属性の全ての値をbldg:BuildingPartに記述する。

![Image](images/tocC_04_img002.png)

ただし、ユースケースで必要ない場合には、bldg:BuildingPartを使用しなくてもよい。
例えば、 図C-22のように、複合施設であっても主たる用途についてのみ都市計画基礎調査が実施されている場合で、用途ごとに棟を分けて取得する必要が無い場合には、bldg:BuildingPartを使用せず、bldg:Buildingのみを使用して記述する。

![Image](images/tocC_04_img003.png)

bldg:BuildingPartは空間属性ももつため、bldg:BuildingPartを記述する場合は、これの集まりとなるbldg:Buildingに空間属性を記述しない。

**表 C-30 — C-19: bldg:BuildingPartの参照条件**

| ID | `/req/bldg/19` |
| --- | --- |
| 主題 | 妥当な建築物データセットの要件 |
| 分類 | 要件分類C-1: 妥当な建築物オブジェクト作成 |
| 条文 | bldg:BuildingPartを使用する場合には、これを参照するbldg:Buildingの空間属性は空とする。 |

なお、一棟の建築物をbldg:Buildingとbldg:BuildingPartとに分割して記述する場合においても、本来は一つの建築物である。そこで、共通となる属性については、全体となるbldg:Buildingの属性として記述する。

**表 C-31 — C-20: 建築物の属性付与基準**

| ID | `/req/bldg/20` |
| --- | --- |
| 主題 | 妥当な建築物データセットの要件 |
| 分類 | 要件分類C-1: 妥当な建築物オブジェクト作成 |
| 条文 | 一棟の建築物に付与されるべき属性は、全体となるbldg:Buildingにのみ付与する。 |

高層階と低層階を組み合わせた建築物があったときに、この建築物を一棟として都市計画基礎調査が実施されている場合には、この都市計画基礎調査の結果は、bldg:Buildingにのみ地物属性として付与する。
また、LOD0及びLOD1の空間属性も、全体となるbldg:Buildingに記述する。

