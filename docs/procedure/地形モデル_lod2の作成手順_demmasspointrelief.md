# S.3.1.3.3 地形モデル（LOD2）の作成手順（dem:MassPointRelief）

## S.3.1.3.3 地形モデル（LOD2）の作成手順（dem:MassPointRelief）

- 航空写真から作成した点群や航空写真レーザ点群をデータ変換しdem:MassPointReliefを作成する。
- 地形の外形を多角形で取得する。

ランダム点群の場合、点の集合から地形モデルの範囲を正確に取得できない。そのため地形の外側の境界（dem:extentのexterior）を必ず作成する。地形の内空の境界（dem:extentのinterior）は任意で取得する。

