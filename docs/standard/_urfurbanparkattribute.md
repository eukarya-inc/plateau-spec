# 4.20.3.1.3 urf:UrbanParkAttribute

## 4.20.3.1.3 urf:UrbanParkAttribute

**表 4-817**

| 型の定義 | 都市公園法第２条第１項で定義する都市公園に関する情報を定義したデータ型。都市公園は、上位の型がもつ属性「gml:name」、「urf:function」、「urf:nominalArea」、「urf:location」及び当該型に定義する属性を記述する。 なお、都市公園のうち、国又は地方公共団体が設置する都市計画施設である公園又は緑地に該当するときは、「urf:OpenSpaceForPublicUse」を必要に応じ別途作成する。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| urf:parkCode     [ UrbanParkAttribute ] | gml::CodeType [1..1] | 公園を一意に識別するコード。コードリスト（Common_parkCode.xml）より選択する。 |
| urf:startFrom     [ UrbanParkAttribute ] | xs::date [1..1] | 公園の設置の年月日（供用開始の日付）。 |
| urf:breakdownOfNominalArea     [ UrbanParkAttribute ] | BreakdownOfNominalArea [0..*] | 都市公園の敷地面積の内訳。 |

