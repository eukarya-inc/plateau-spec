# O.2.4.4.2 施設識別属性（uro:FacilityIdAttribute）

## O.2.4.4.2 施設識別属性（uro:FacilityIdAttribute）

uro:FacilityIdAttributeは、施設の位置を特定する情報及び施設を識別する情報を記述するためのデータ型である。uro:FacilityIdAttributeは、施設を識別するための情報として、識別子（uro:id）や正式な名称以外の呼称（uro:alternativeName）に加え、施設の位置を示すための、都道府県（uro:prefecture）、市区町村（uro:city）及び開始位置の経緯度（uro:startLat、uro:startLong）を属性としてもつ。また、鉄道上や道路上の施設については、路線や距離標での位置特定のための属性（uro:route、uro:startPost、uro:endPost）を使用できる。

なお、河川管理施設の場合は、uro:FacilityIdAttributeを継承するuro:RiverFacilityIdAttributeを使用する。これにより、左右岸上での位置の情報を記述できる。

