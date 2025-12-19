# 4.3.3.2.4 uro:RoadStructureAttribute

## 4.3.3.2.4 uro:RoadStructureAttribute

**表 4-156**

| 型の定義 | 路線、同等以上の道路との交差点又は道路構造の変化点（トンネル、橋梁）で変化する場所で道路を区切った区間における、道路の構造。 | ← |
| --- | --- | --- |
| 上位の型 | uro:RoadAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:widthType     [ RoadStructureAttribute ] | gml::CodeType [0..1] | 幅員の区分。コードリスト（RoadStructureAttribute_widthType.xml）より選択する。都市計画基礎調査で収集されている場合にのみ作成する。 |
| uro:width     [ RoadStructureAttribute ] | gml::LengthType [0..1] | 中央帯、車道、路肩、植樹帯、歩道等及び環境施設帯（環境施設帯の中の路肩、植樹帯、歩道等の部分を除いた部分）の幅員を合計した幅員。単位はm（uom=”m”）とする。 |
| uro:numberOfLanes     [ RoadStructureAttribute ] | xs::integer [0..1] | 上下線の合計（一方通行区間の場合を除く）の車線数。道路構造令第 2 条第 7 号の登坂車線、同第 2 条第 6 号にいう付加追越車線、同第 2 条 8 号の屈折車線、同第 2 条第 9 号の変速車線及び同第 2 条第 14 号の停車帯、及びゆずり車線は車線数には含めない。交差点付近において、右左折のための車線が設けられている場合はこの数を含まない。  「1 車線道路」は道路構造令第 5 条 1 項ただし書きによって、車線により構成されない車道を持つ道路であるが、ここでは車線数＝1とする。「1車線道路」は車道幅員が5.5m未満の場合とする。 道路構造が「交差部」の場合、この属性は作成しない。 |
| uro:sectionType     [ RoadStructureAttribute ] | gml::CodeType [0..1] | 道路構造の種別。コードリスト（RoadStructureAttribute_sectionType.xml）より選択する。 |

