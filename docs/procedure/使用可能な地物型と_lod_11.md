# P.2.2 使用可能な地物型とLOD

## P.2.2 使用可能な地物型とLOD

地下埋設物モデルは、LODごとに、使用すべき地物型やその空間属性が異なる。

**表 P-2 — P-1: 地下埋設物記述のLOD対応条件**

| ID | `/req/unf/1` |
| --- | --- |
| 主題 | 妥当な地下埋設物データセットの要件 |
| 分類 | 要件分類P-1: 妥当な地下埋設物オブジェクト |
| 条文 | 地下埋設物の記述には、指定されたLODに対応する地物型及び空間属性を使用する。 |

地下埋設物の記述には、i-URのUrban Objectモジュールに定義された、uro:UtilityNetworkElementを継承する地物型を使用する。

uro:UtilityNetworkElementは複数のオフセット・デプス情報（uro:OffsetDepth）をもつことができる。uro:OffsetDepthは特定の位置（属性pos）での道路縁等からの離れ（属性offset）や土被りの深さ（属性depth）をもつ。土被りとは、地表面から、地下埋設物の上端までをいう。また、土被りの深さを範囲（最大値となるmaxDepth、最小値となるminDepth）でもつこともできる。

標準製品仕様書では、uro:UtilityNetworkElementを継承する三つの地物型（及びこれらを継承する地物型）を用いて地下埋設物のネットワークを表現する。

uro:UtilityLinkは、ネットワークのリンクであり、ケーブル（uro:Cable）、管路（uro:Pipe）及びトラフ（函渠）等の保護設備（uro:Duct）に区分される。また、ケーブルは通信ケーブル（uro:TelecommunicationsCable）、管路は水道管（WaterPipe）というように、ケーブルと管路には、さらに具体化された下位の地物型が定義されている。

uro:UtilityNodeは、ネットワークのノードに該当する。バルブやガバナ、制水栓等の弁栓類というようなリンクとリンクのつながりを制御する設備（uro:Appurtenance）を表す地物型である。この制御設備には、変圧器、整圧器又は開閉器といった地上に設置された設備を含む。

uro:UtilityNodeContainerは、uro:UtilityNodeを格納し、自らもノードとなることができる地物型である。uro:UtilityNodeContainerは、マンホール（uro:Manhole）及びハンドホール（uro:Handhole）の二つの地物型に区分される。

- 地下埋設物モデル（LOD0）では、地下埋設物の形状を点、線又は面として表現する。
- 地下埋設物モデル（LOD1）では、地下埋設物の形状を、地下埋設物の外周に地表からの一律の高さを下向きに与えた立体として表現する。
- 地下埋設物モデル（LOD2）では、地下埋設物の形状を、地下埋設物の外周に地下埋設物が埋まった深さから一律の高さを下向きに与えた立体として表現する。
- 地下埋設物モデル（LOD3）では、地下埋設物の詳細な形状を、面の集まりとして表現する。
- 地下埋設物モデル（LOD4）では、LOD3の外形に加えて、内部の形状を表現する。

地下埋設物の各LODにおいて使用可能な地物型と空間属性を、表P-3に示す。

**表 P-3 — 地下埋設物モデルの記述に使用する地物型と空間属性**

| 地物型 | 空間属性 | LOD0 | LOD1 | LOD2 | LOD3 | LOD4 | 適用 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| uro: Pipe及びこれを継承する地物型 |  | ● | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  |  | 原典資料の取得方法に従う。 |
| ↑ | frn:lod1Geometry |  | ● |  |  |  | Solidとする。 |
| ↑ | frn:lod2Geometry |  |  | ● |  |  | Solidとする。 |
| ↑ | frn:lod3Geometry |  |  |  | ● |  | CompositeSurface又はMultiSurfaceとする。 |
| ↑ | frn:lod4Geometry |  |  |  |  | ● | CompositeSurface又はMultiSurfaceとする。 |
| uro:Cable及びこれを継承する地物型 |  | ● | ● | ● | ● | 〇 a) |  |
| ↑ | uro:lod0Geometry | ● |  |  |  |  | 原典資料の取得方法に従う。 |
| ↑ | frn:lod1Geometry |  | ● |  |  |  | Solidとする。 |
| ↑ | frn:lod2Geometry |  |  | ● |  |  | Solidとする。 |
| ↑ | frn:lod3Geometry |  |  |  | ● |  | CompositeSurface又はMultiSurfaceとする。 |
| ↑ | frn:lod4Geometry |  |  |  |  | ■ | LOD4を作成する場合は必須とする。 |
| uro: Duct |  | ● | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  |  | 原典資料の取得方法に従う。 |
| ↑ | frn:lod1Geometry |  | ● |  |  |  | Solidとする。 |
| ↑ | frn:lod2Geometry |  |  | ● |  |  | Solidとする。 |
| ↑ | frn:lod3Geometry |  |  |  | ● |  | CompositeSurface又はMultiSurfaceとする。 |
| ↑ | frn:lod4Geometry |  |  |  |  | ● | CompositeSurface又はMultiSurfaceとする。 |
| uro:Appurtenance |  | ● | ● | ● | ● | 〇 a) |  |
| ↑ | uro:lod0Geometry | ● |  |  |  |  | 原典資料の取得方法に従う。 |
| ↑ | frn:lod1Geometry |  | ● |  |  |  | Solidとする。 |
| ↑ | frn:lod2Geometry |  |  | ● |  |  | Solidとする。 |
| ↑ | frn:lod3Geometry |  |  |  | ● |  | CompositeSurface又はMultiSurfaceとする。 |
| ↑ | frn:lod4Geometry |  |  |  |  | ■ | LOD4を作成する場合は必須とする。 |
| uro: Manhole |  | ● | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  |  | 原典資料の取得方法に従う。 |
| ↑ | frn:lod1Geometry |  | ● |  |  |  | Solidとする。 |
| ↑ | frn:lod2Geometry |  |  | ● |  |  | Solidとする。 |
| ↑ | frn:lod3Geometry |  |  |  | ● |  | CompositeSurface又はMultiSurfaceとする。 |
| ↑ | frn:lod4Geometry |  |  |  |  | ● | CompositeSurface又はMultiSurfaceとする。 |
| uro: Handhole |  | ● | ● | ● | ● | ● |  |
| ↑ | uro:lod0Geometry | ● |  |  |  |  | 原典資料の取得方法に従う。 |
| ↑ | frn:lod1Geometry |  | ● |  |  |  | Solidとする。 |
| ↑ | frn:lod2 Geometry |  |  | ● |  |  | Solidとする。 |
| ↑ | frn:lod3 Geometry |  |  |  | ● |  | CompositeSurface又はMultiSurfaceとする。 |
| ↑ | frn:lod4Geometry |  |  |  |  | ● | CompositeSurface又はMultiSurfaceとする。 |
| ●：必須■：条件付必須〇：任意（ユースケースに応じて要否を決定してよい） | ← | ← | ← | ← | ← | ← | ← |

地下埋設物モデル（LOD4）において、uro:Cable及びuro:Appurtenanceの形状表現は外部の形状のみとしてよい。このときの外部の形状は、地下埋設物モデル（LOD3）と一致する。内部の形状表現は、ユースケースに応じて要否を決定してよい。

