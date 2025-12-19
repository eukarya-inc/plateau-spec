# 4.2.3.5.16 uro:IfcPsetBuildingCommon

## 4.2.3.5.16 uro:IfcPsetBuildingCommon

**表 4-73**

| 型の定義 | IFCで記述された建築物に共通となる属性の集まり。 | ← |
| --- | --- | --- |
| 上位の型 | uro:IfcAttribute | ← |
| ステレオタイプ | << DataType >> | ← |
| 自身に定義された属性 | ← | ← |
| 属性名 | 属性の型及び多重度 | 定義 |
| uro:buildingId     [ IfcPsetBuildingCommon ] | xs::string [0..1] | 建築物に付与される固有の識別子。計画申請時に一時的な識別子が付与される。この一時的な識別子は、建物が法的な建物とプロパティのデータベースに登録されると、恒久的な識別子に変更される。 |
| uro:isPermanentId     [ IfcPsetBuildingCommon ] | xs::boolean [0..1] | 建物に割り当てられた識別子が永続的か一時的かを示す。 1：永続的 0：一時的 |
| uro:mainFireUse     [ IfcPsetBuildingCommon ] | xs::string [0..1] | 建築物の主な防災用途で、関連する国の建築基準法で定められた防災用途分類表から割り当てられるもの。 |
| uro:ancillaryFireUse     [ IfcPsetBuildingCommon ] | xs::string [0..1] | 付属的な防災用途で、関連する国家建築基準法の防災用途分類表から割り当てられたもの。 |
| uro:sprinklerProtection     [ IfcPsetBuildingCommon ] | xs::boolean [0..1] | スプリンクラーで保護されているか、されていないかを示す。 1：保護されている 0：保護されていない |
| uro:sprinklerProtectionAutomatic     [ IfcPsetBuildingCommon ] | xs::boolean [0..1] | 自動スプリンクラーで保護されているかどうかを示す。 1：保護されている 0：保護されていないこれは、プロパティ “SprinklerProtection” が 1（true） に設定されている場合にのみ、指定されるべきである。 |
| uro:occupancyType     [ IfcPsetBuildingCommon ] | gml::CodeType [0..1] | 入居者タイプ。国の建築基準法に従って定義される。コードリスト（IfcPsetBuildingCommon_occupancyType.xml）から選択する。この属性を使用する場合は、コードリストを作成する。 |
| uro:grossPlannedArea     [ IfcPsetBuildingCommon ] | gml::MeasureType [0..1] | 建築物の計画総面積。 |
| uro:numberOfStoreys     [ IfcPsetBuildingCommon ] | xs::integer [0..1] | 建築物内の階数。 |
| uro:yearOfConstruction     [ IfcPsetBuildingCommon ] | xs::gYear [0..1] | この建築物の建築年。 |
| uro:isLandmarked     [ IfcPsetBuildingCommon ] | xs::boolean [0..1] | この建築物が歴史的建造物として登録されているか否か。 1：されている 0：されていない |

