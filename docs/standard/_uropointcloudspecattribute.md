# 4.23.3.1.3 uro:PointCloudSpecAttribute

## 4.23.3.1.3 uro:PointCloudSpecAttribute

**表 4-842**

| 型の定義 | 三次元点群データの仕様に関する詳細情報。 | ← |
| --- | --- | --- |
| 上位の型 | — | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:class     [ PointCloudSpecAttribute ] | gml::CodeType [1..1] | 三次元点群データの種類。コードリスト（PointCloudSpecAttribute_class.xml）より選択する。 |
| uro:function     [ PointCloudSpecAttribute ] | gml::CodeType [0..*] | 三次元点群データの計測方法。コードリスト（PointCloudSpecAttribute_function.xml）より選択する。 |
| uro:srcScale     [ PointCloudSpecAttribute ] | gml::CodeType [0..1] | 原典の三次元点群データの地図情報レベル。コードリスト（PointCloudSpecAttribute_srcScale.xml）より選択する。 |
| uro:srcRequiredHorizontalAccuracy     [ PointCloudSpecAttribute ] | gml::LengthType [0..1] | 原典の三次元点群データの水平要求精度。単位はｍとする。 |
| uro:srcRequiredVerticalAccuracy     [ PointCloudSpecAttribute ] | gml::LengthType [0..1] | 原典の三次元点群データの垂直要求精度。単位はｍとする。 |
| uro:srcRequiredPointDensity     [ PointCloudSpecAttribute ] | gml::MeasureType [0..1] | 原典の三次元点群データの要求点密度。単位は点/m2とする。 |
| uro:srcFilteringCriteria     [ PointCloudSpecAttribute ] | gml::CodeType [0..*] | 原典の三次元点群データのフィルタリングの状況。コードリスト（PointCloudSpecAttribute_srcFilteringCriteria.xml）より選択する。 |
| uro:srcGridInterval     [ PointCloudSpecAttribute ] | gml::MeasureType [0..1] | 原典の三次元点群データの格子間隔。単位はm又はdegとする。 三次元点群データがグリッドデータの場合に記述する。 |
| uro:isIntensity     [ PointCloudSpecAttribute ] | xs::boolean [0..1] | 三次元点群データにIntensity（反射強度）の情報が含まれているかを示す。 1：含まれている 0：含まれていない |
| uro:spectralBandSet     [ PointCloudSpecAttribute ] | gml::CodeType [0..*] | 三次元点群データに含まれる波長スペクトル情報を示すコード。コードリスト（PointCloudSpecAttribute_spectralBandSet.xml）から選択する。 |
| uro:isReturnInfo     [ PointCloudSpecAttribute ] | xs::boolean [0..1] | 三次元点群データにReturnNumber及びNumber of Returns（反射番号、総反射回数）の情報が含まれているかを示す。 1：含まれている 0：含まれていない |
| uro:isStandardClassification     [ PointCloudSpecAttribute ] | xs::boolean [0..1] | 三次元点群データがASPRS（American Society for Photogrammetry and Remote Sensing）の定義したクラス分類に適合しているかを示す。 1：適合している 0：適合していないなお、具体的なクラス分類は「別表 ASPRSが定義したクラス分類」に示す。 |
| uro:customClassificationType     [ PointCloudSpecAttribute ] | gml::CodeType [0..*] | 独自のクラス分類コード。三次元点群データを独自にクラス分類している場合に使用しているクラス分類コードを入力する。コードリスト（PointCloudSpecAttribute_customClassificationType.xml）を作成し、選択する。 |
| uro:pointCloudAssociationMethod     [ PointCloudSpecAttribute ] | gml::CodeType [1..1] | 三次元点群データと地物の対応づけの方法。コードリスト（PointCloudSpecAttribute_pointCloudAssociationMethod.xml）より選択する。 |
| uro:pointSubsamplingInfo     [ PointCloudSpecAttribute ] | gml::CodeType [0..1] | 三次元点群データの間引きの情報。コードリスト（PointCloudSpecAttribute_pointSubsamplingInfo.xml）より選択する。 |

別表：ASPRSが定義したクラス分類

**表 4-843**

| クラス値 | 定義 |
| --- | --- |
| 0 | Created, Never Classified (未分類、作成時のまま) |
| 1 | Unclassified (未分類) |
| 2 | Ground (地面) |
| 3 | Low Vegetation (低植生) |
| 4 | Medium Vegetation (中植生) |
| 5 | High Vegetation (高植生) |
| 6 | Building (建物) |
| 7 | Low Point (Noise) (ノイズ) |
| 8 | Model Key-Point (Mass Point) [Format 0-5採用] |
| 9 | Water (水域) |
| 10 | Rail (レール) |
| 11 | Road Surface (道路表面) |
| 12 | Overlap Points (オーバーラップ点) [Format 0-5採用] |
| 13 | Wire – Guard (Shield) (ガード線) |
| 14 | Wire – Conductor (Phase) (送電線) |
| 15 | Transmission Tower (送電塔) |
| 16 | Wire-Structure Connector (電線-構造物接続部) |
| 17 | Bridge Deck (橋梁デッキ) |
| 18 | High Noise (高ノイズ) |
| 19 | Overhead Structure (頭上構造物) |
| 20 | Ignored Ground (無視された地面) |
| 21 | Snow (雪) |
| 22 | Temporal Exclusion (時間的除外) |
| 23-63 | Reserved for ASPRS Definition |
| 64-255 | User Definable (ユーザー定義可能) |

