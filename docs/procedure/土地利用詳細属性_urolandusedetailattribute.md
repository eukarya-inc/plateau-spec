# I.2.4.2 土地利用詳細属性（uro:landUseDetailAttribute）

## I.2.4.2 土地利用詳細属性（uro:landUseDetailAttribute）

都市計画基礎調査の土地利用現況として記録された情報を格納するための属性である。属性の型であるuro:LandUseDetailAttributeには都市独自の土地利用の区分（uro:orgLandUse）の他、面積（uro:nominalArea）や建ぺい率（uro:buildingCoverageRate）など、様々な属性の箱が用意されている。都市計画基礎調査でこれらの情報が収集されている場合には、定義済みの属性を使用する。また、該当する属性が無い場合には、汎用属性を使用して拡張できる。

なお、都市独自の土地利用の区分（uro:orgLandUse）を作成する場合には、区分を示すコードリスト（LandUseDetailAttribute_orgLandUse.xml）を作成しなければならない。

