# 4.11.3.1.1 brid:Bridge

## 4.11.3.1.1 brid:Bridge

**表 4-504**

| 型の定義 | 橋梁。橋梁とは、道路、鉄道、水路等の輸送路において、輸送の障害となる河川、渓谷、湖沼、海峡あるいは他の道路、鉄道、水路等の上方にこれらを横断するために建設される構造物の総称である。市街地において効率的な土地利用の観点から、道路上あるいは河川上の空間に連続して建設される高架橋も橋梁の一形態である。[土木工学ハンドブック] 標準製品仕様書では以下を対象とする。道路法（昭和27年法律第180号）第２条第１項に規定する「橋」。橋長２.０m以上を対象とし、高架橋及び桟道橋を含む[参考 [中部地方整備局 道路施設台帳作成マニュアル](https://www.cbr.mlit.go.jp/architecture/kensetsugijutsu/download/index.htm)]。桟道橋：斜面を通過する道路で、橋桁の一側が斜面に接し、反対側が橋脚になっている部分 [[作業規程の準則　付録7　公共測量標準図式](https://www.gsi.go.jp/gijyutukanri/gijyutukanri41018.html)]。鉄道事業者法施行規則別表第一に定める鉄道施設の「橋りょう」。鉄道事業者法施行規則別表第一に定める鉄道施設の「こ線橋」。こ（跨）線橋：駅構内の鉄道を横断するために構築された橋 [[作業規程の準則　付録7　公共測量標準図式](https://www.gsi.go.jp/gijyutukanri/gijyutukanri41018.html)]。道路法第30条第1項第10号に定める「横断歩道橋」。横断歩道橋：車道を横断する歩行者を車道から分離するための通路が、道路の上方に設置される道路横断施設[道路基盤地図情報]。また、標準製品仕様書では、建築基準法第44条第1項第4号において示された公共用歩廊で、道路上に設けられた高架の歩行者専用道路（ペデストリアンデッキ、スカイウェイ、スカイウォークなどと呼ばれる）を含む。高架橋のように延長の長い橋梁は、管理区間及び上部工の境界（伸縮装置の設置部）で区切ることができる。 | ← |
| --- | --- | --- |
| 上位の型 | brid:_AbstractBridge | ← |
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
| brid:class     [ _AbstractBridge ] | gml::CodeType [0..1] | 橋梁の形態による区分。コードリスト（Bridge_class.xml）より選択する。多重度は任意となっているが、運用上必須とする。 |
| brid:function     [ _AbstractBridge ] | gml::CodeType [0..*] | 橋梁の主たる機能による区分。コードリスト（Bridge_function.xml）より選択する。多重度は任意となっているが、運用上必須とする。 |
| brid:usage     [ _AbstractBridge ] | gml::CodeType [0..*] | 橋梁の用途。brid:functionで指定された機能と異なる場合に実際の用途を示すためにこの属性を用いる。標準製品仕様書では使用しない。 |
| brid:yearOfConstruction     [ _AbstractBridge ] | xs::gYear [0..1] | 橋梁が建築された年度。完成した年度とする。多重度は任意となっているが、運用上必須とする。 |
| brid:yearOfDemolition     [ _AbstractBridge ] | xs::gYear [0..1] | 橋梁が解体された年度。 |
| brid:isMovable     [ _AbstractBridge ] | xs::boolean [0..1] | 可動橋か否かの別。 1：可動橋である 0：可動橋ではない |
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
| brid:address             [ _AbstractBridge ] | core:Address [0..*] | 橋梁に紐づけられた住所。標準製品仕様書では使用しない。 |
| brid:boundedBy             [ _AbstractBridge ] | brid:_BoundarySurface [0..*] | 橋梁を構成する外壁面（brid:WallSurface）や屋外床面（brid:OuterFloorSurface）等の境界面。 |
| brid:consistsOfBridgePart             [ _AbstractBridge ] | brid:BridgePart [0..*] | 橋梁の部分（brid:BridgePart）。橋梁の一部（brid:BridgePart）の場合は使用しない。 |
| brid:interiorBridgeInstallation             [ _AbstractBridge ] | brid:IntBridgeInstallation [0..*] | 橋梁の内部空間において、その外観を特徴付ける部分のうち、管理用通路・階段のような、橋梁の構造上不可欠ではない付属物（brid:IntBridgeInstallation）。 |
| brid:interiorBridgeRoom             [ _AbstractBridge ] | brid:BridgeRoom [0..*] | 橋梁の内部空間（brid:InteriorBridgeRoom）。 |
| brid:lod1MultiSurface             [ _AbstractBridge ] | gml:MultiSurface [0..1] | 橋梁の外形を構成する面の集まり。標準製品仕様書では使用しない。 |
| brid:lod1Solid             [ _AbstractBridge ] | gml:Solid [0..1] | 橋梁の外周に一律の高さを与えた立体。一律の高さは、水面（陸上に設置されている場合は地表面）から橋梁の最高高さとする。橋梁の起点側と終点側の間での高さの変化（橋梁中央部の高さ）は表現しない。 |
| brid:lod1TerrainIntersection             [ _AbstractBridge ] | gml:MultiCurve [0..1] | LOD1における橋梁と地形との交線。標準製品仕様書では使用しない。 |
| brid:lod2MultiCurve             [ _AbstractBridge ] | gml:MultiCurve [0..1] | LOD2において橋梁の立体表現に加え、線状の表現を行う場合に使用する。標準製品仕様書では使用しない。 |
| brid:lod2MultiSurface             [ _AbstractBridge ] | gml:MultiSurface [0..1] | 橋梁の主要構造の外形を構成する面の集まり。面の集まりが立体の境界としての要件を満たすことが出来ない場合に、LOD2の幾何オブジェクトとして使用する。橋梁をbrid:BridgePartの集まりとして記述する場合、brid:Bridgeの空間属性は空となる。 |
| brid:lod2Solid             [ _AbstractBridge ] | gml:Solid [0..1] | 橋梁の主要構造の外形を示す立体。この時の立体は外壁等の、境界面により構成される。橋梁をbrid:BridgePartの集まりとして記述する場合、brid:Bridgeの空間属性は空となる。 |
| brid:lod2TerrainIntersection             [ _AbstractBridge ] | gml:MultiCurve [0..1] | LOD2における橋梁と地形との交線。標準製品仕様書では使用しない。 |
| brid:lod3MultiCurve             [ _AbstractBridge ] | gml:MultiCurve [0..1] | LOD3において橋梁の立体表現に加え、線状の表現を行う場合に使用する。橋梁の立体表現に加え、線状の表現を行う場合に使用する。これらの表現に意味を持たせる場合は、brid:BridgeInstallationやbrid:BridgeConstructionElementを使用する。標準製品仕様書では使用しない。 |
| brid:lod3MultiSurface             [ _AbstractBridge ] | gml:MultiSurface [0..1] | 橋梁の主要構造の外形を構成する面の集まり。面の集まりが立体の境界としての要件を満たすことが出来ない場合に、LOD3の幾何オブジェクトとして使用する。 |
| brid:lod3Solid             [ _AbstractBridge ] | gml:Solid [0..1] | 橋梁の主要構造の外形を示す立体。この時の立体は、外壁等を区分する境界面及び開口部の面（境界面の内空として作成されている場合）により構成される。 |
| brid:lod3TerrainIntersection             [ _AbstractBridge ] | gml:MultiCurve [0..1] | LOD3における橋梁と地形との交線。標準製品仕様書では使用しない。 |
| brid:lod4MultiCurve             [ _AbstractBridge ] | gml:MultiCurve [0..1] | LOD4において橋梁の立体表現に加え、線状の表現を行う場合に使用する。標準製品仕様書では使用しない。 |
| brid:lod4MultiSurface             [ _AbstractBridge ] | gml:MultiSurface [0..1] | 橋梁の外形を構成する面の集まり。面の集まりが立体の境界としての要件を満たすことが出来ない場合に、LOD4の幾何オブジェクトとして使用する。 |
| brid:lod4Solid             [ _AbstractBridge ] | gml:Solid [0..1] | 橋梁の詳細構造の外形を示す立体。この時の立体は、外壁等を区分する境界面及び開口部の面（境界面の内空として作成されている場合）により構成される。 |
| brid:lod4TerrainIntersection             [ _AbstractBridge ] | gml:MultiCurve [0..1] | LOD4における橋梁と地形との交線。標準製品仕様書では使用しない。 |
| brid:outerBridgeConstruction             [ _AbstractBridge ] | brid:BridgeConstructionElement [0..*] | 橋梁を構成する部分のうち、橋脚、橋台のような構造上不可欠な部分（brid:BridgeConstructionElement）。 |
| brid:outerBridgeInstallation             [ _AbstractBridge ] | brid:BridgeInstallation [0..*] | 橋梁の外観を特徴付ける部分のうち、アンテナや航空障害灯、管理用通路・階段のような、橋梁の構造上不可欠ではない付属物（brid:BridgeInstallation）。 |
| uro:bridBaseAttribute             [ _AbstractBridge ] | uro:ConstructionBaseAttribute [0..1] | 橋梁の管理に関する基本的な情報。 |
| uro:bridDataQualityAttribute             [ _AbstractBridge ] | uro:DataQualityAttribute [0..1] | 作成するデータの品質に関する情報。brid:Bridgeの場合は、原則必須とする。ただし、brid:BridgePartが品質属性をもつ場合は、省略できる。 |
| uro:bridDisasterRiskAttribute             [ _AbstractBridge ] | uro:DisasterRiskAttribute [0..*] | 橋梁の災害リスクに関する情報。uro:DisasterRiskAttributeの下位型を使用して記述する。 brid:BridgePartには作成しない。（brid:Bridgeのみに作成する。） |
| uro:bridDmAttribute             [ _AbstractBridge ] | uro:DmAttribute [0..*] | LOD0の幾何形状。 brid:BridgePartには作成しない。（brid:Bridgeにのみ作成する。） |
| uro:bridFacilityAttribute             [ _AbstractBridge ] | uro:FacilityAttribute [0..*] | uro:bridFacilityTypeAttribute.classによって指定された分野における施設管理情報。 brid:BridgePartには作成しない。（brid:Bridgeにのみ作成する。） |
| uro:bridFacilityIdAttribute             [ _AbstractBridge ] | uro:FacilityIdAttribute [0..1] | uro:bridFacilityTypeAttribute.classによって指定された分野における施設の識別情報。 brid:BridgePartには作成しない。（brid:Bridgeにのみ作成する。） |
| uro:bridFacilityTypeAttribute             [ _AbstractBridge ] | uro:FacilityTypeAttribute [0..*] | 特定分野における施設の分類情報。 brid:BridgePartには作成しない。（brid:Bridgeにのみ作成する。） |
| uro:bridFunctionalAttribute             [ _AbstractBridge ] | uro:BridgeFunctionalAttribute [0..1] | 橋梁の機能に関する情報。 |
| uro:bridKeyValuePairAttribute             [ _AbstractBridge ] | uro:KeyValuePairAttribute [0..*] | 属性を拡張するための仕組み。コ－ド値以外の属性を拡張する場合は、gen:_GenericAttributeの下位型を使用する。 brid:BridgePartには作成しない。（brid:Bridgeにのみ作成する。） |
| uro:bridRiskAssessmentAttribute             [ _AbstractBridge ] | uro:ConstructionRiskAssessmentAttribute [0..1] | 橋梁の損傷に関する情報。 |
| uro:bridStructureAttribute             [ _AbstractBridge ] | uro:BridgeStructureAttribute [0..1] | 橋梁の構造に関する情報。 |

