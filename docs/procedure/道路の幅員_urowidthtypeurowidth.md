# D.3.2.4 道路の幅員（uro:widthType、uro:width）

## D.3.2.4 道路の幅員（uro:widthType、uro:width）

uro:widthTypeは、道路の幅員を一定の長さで区分した属性である（表D-13）。

**表 D-13 — 幅員の区分**

| ファイル名 | RoadStructureAttribute_widthType.xml |
| --- | --- |
| ファイルURL | [https://www.geospatial.jp/iur/codelists/3.2/RoadStructureAttribute_widthType.xml](https://www.geospatial.jp/iur/codelists/3.2/RoadStructureAttribute_widthType.xml) |
| コード | 説明 |
| 1 | 15m以上 |
| 2 | 6m以上15m未満 |
| 3 | 4m以上6m未満 |
| 4 | 4m未満 |
| 出典： [都市計画基礎調査実施要領](https://www.mlit.go.jp/toshi/tosiko/kisotyousa001.html) | ← |

都市計画基礎調査において「道路の状況」として道路幅員が収集されている場合に、都市計画基礎調査のデータを用いてuro:widthTypeを作成する。都市計画基礎調査で収集されていない場合にはuro:widthTypeは作成しない。また、都市計画基礎調査で収集されたデータが表D-13に示す区分と異なる区分の場合は、拡張製品仕様書において収集されたデータに合わせたコードリストを作成し、この属性を作成する。

uro:widthは、道路の幅員値（実数値）である。都市計画基礎調査において、前述の幅員の区分だけではなく、幅員値が収集されている場合にはこれを用いてこの属性を作成する。都市計画基礎調査で収集されていない場合には、他の資料から収集又は道路の図形から算出する。想定される取得方法を以下に示す。

- 全国道路・街路交通量情勢調査（一般交通量調査）の値を採用する。
- 道路台帳に記載された幅員値を採用する。
- 道路の図形を用いてGIS上で平均幅員を算出する。

