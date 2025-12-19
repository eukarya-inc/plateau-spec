# 4.12.3.1.16 tun:IntTunnelInstallation

## 4.12.3.1.16 tun:IntTunnelInstallation

**表 4-577**

| 型の定義 | トンネルの内部に設置された、恒久的に存在する固定的な設備（内部付属物）。内部付属物は、トンネルの付帯的な設備であり、主要な部分であってはならない。また、内部付属物は、トンネル（tun:Tunnel又はtun:TunnelPart又は内空（tun:HollowSpace）と接していなければならない。 避難連絡坑、受変電設備、予備電源設備、排水設備、換気設備、排水設備、監視制御設備、通報装置、非常用警報装置、消火設備、避難誘導設備、その他。 ただし、ユースケースの要求に応じて、取得対象とする内部付属物を限定してもよく、また、内部付属物として取得せず建築物の一部として取得してもよい。 | ← |
| --- | --- | --- |
| 上位の型 | core:_CityObject | ← |
| ステレオタイプ | << FeatureType >> | ← |
| 継承する属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| gml:description     [ _Feature ] | gml:StringOrRefType [0..1] | 都市オブジェクトの概要。 |
| gml:name     [ _Feature ] | gml::CodeType [0..1] | 都市オブジェクトを識別する名称。文字列とする。 |
| gml:boundedBy     [ _Feature ] | gml::Envelope [0..1] | 都市オブジェクトの範囲及び適用される空間参照系。 CityModelの場合のみ必須とする。 |
| core:creationDate     [ _CityObject ] | xs::date [0..1] | データが作成された日。運用上必須とする。 |
| core:terminationDate     [ _CityObject ] | xs::date [0..1] | データが削除された日。 |
| core:relativeToTerrain     [ _CityObject ] | core::RelativeToTerrainType [0..1] | 地表面との相対的な位置関係。標準製品仕様書では使用しない。 |
| core:relativeToWater     [ _CityObject ] | core::RelativeToWaterType [0..1] | 水面との相対的な位置関係。標準製品仕様書では使用しない。 |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| tun:class     [ IntTunnelInstallation ] | gml::CodeType [0..1] | 内部付属物の形態による区分。標準製品仕様書では使用しない。 |
| tun:function     [ IntTunnelInstallation ] | gml::CodeType [0..*] | 内部付属物の主たる働き。コードリスト（TunnelInstallation_function.xml）より選択する。 |
| tun:usage     [ IntTunnelInstallation ] | gml::CodeType [0..*] | 内部付属物の主な使い道。標準製品仕様書では使用しない。 |
| 継承する関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| core:externalReference             [ _CityObject ] | core:ExternalReference [0..*] | 外部への参照。標準製品仕様書では使用しない。 |
| gen:dateAttribute             [ _CityObject ] | gen:dateAttribute [0..*] | 日付型属性。属性を追加したい場合に使用する。 |
| gen:doubleAttribute             [ _CityObject ] | gen:doubleAttribute [0..*] | 実数型属性。属性を追加したい場合に使用する。 |
| gen:genericAttributeSet             [ _CityObject ] | gen:genericAttributeSet [0..*] | 汎用属性のセット。属性を追加したい場合に使用する。 |
| gen:intAttribute             [ _CityObject ] | gen:intAttribute [0..*] | 整数型属性。属性を追加したい場合に使用する。 |
| gen:measureAttribute             [ _CityObject ] | gen:measureAttribute [0..*] | 単位付き数値型属性。属性を追加したい場合に使用する。 |
| gen:stringAttribute             [ _CityObject ] | gen:stringAttribute [0..*] | 文字列型属性。属性を追加したい場合に使用する。 |
| gen:uriAttribute             [ _CityObject ] | gen:uriAttribute [0..*] | URI型属性。属性を追加したい場合に使用する。 |
| uro:pointCloud             [ _CityObject ] | uro:AbstractPointCloud [0..1] | ポイントクラウドへの参照。 |
| 自身に定義された関連役割 | ← | ← |
| 関連役割名 | 関連役割の型及び多重度 | 定義 |
| tun:boundedBy             [ IntTunnelInstallation ] | tun:_BoundarySurface [0..*] | 内部付属物を構成する内壁、天井等の境界面への参照。付属物の境界面が内部空間（tun:HollowSpace）の境界面となる場合にのみ作成する。ただし、空間から内部付属物を除く必要が無い場合は、内部付属物の形状を構成する面を、境界面（tun:_BoundarySurface）にする必要はない。 |
| tun:lod4Geometry             [ IntTunnelInstallation ] | gml:_Geometry [0..1] | 内部付属物のLOD4の形状。gml:MultiSurfaceにより記述することを基本とする。容積の算出等ユースケースで必要な場合は、gml:Solidを使用する。 |

