
# UltimateBravery
Make random builds for league of legends

Table of Contents

1. [Champ Viewing API](#champs)
1. [Item Viewing API](#items)

<a name="champs"></a>

## champs

| Specification | Value |
|-----|-----|
| Resource Path | /champs |
| API Version | 1.0.0 |
| BasePath for the API | http://ub.jhrb.us/ |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /champs/ | [GET](#getChamps) | Gets all of the champions. |



<a name="getChamps"></a>

#### API: /champs/ (GET)


Gets all of the champions.



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Champion](#github.com.hibooboo2.ultimateBravery.lolapi.Champion) |  |
| 404 | object | error | Champs not found |




### Models

<a name="github.com.hibooboo2.ultimateBravery.lolapi.ChampInfo"></a>

#### ChampInfo

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Attack | int |  |
| Defense | int |  |
| Difficulty | int |  |
| Magic | int |  |

<a name="github.com.hibooboo2.ultimateBravery.lolapi.ChampStats"></a>

#### ChampStats

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Armor | int |  |
| Armorperlevel | int |  |
| Attackdamage | int |  |
| Attackdamageperlevel | int |  |
| Attackrange | int |  |
| Attackspeedoffset | int |  |
| Attackspeedperlevel | int |  |
| Crit | int |  |
| Critperlevel | int |  |
| Hp | int |  |
| Hpperlevel | int |  |
| Hpregen | int |  |
| Hpregenperlevel | int |  |
| Movespeed | int |  |
| Mp | int |  |
| Mpperlevel | int |  |
| Mpregen | int |  |
| Mpregenperlevel | int |  |
| Spellblock | int |  |
| Spellblockperlevel | int |  |

<a name="github.com.hibooboo2.ultimateBravery.lolapi.Champion"></a>

#### Champion

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Allytips | array |  |
| Blurb | string |  |
| Enemytips | array |  |
| ID | int |  |
| Image | github.com.hibooboo2.ultimateBravery.lolapi.Image |  |
| Info | github.com.hibooboo2.ultimateBravery.lolapi.ChampInfo |  |
| Key | string |  |
| Lore | string |  |
| Name | string |  |
| Partype | string |  |
| Passive | github.com.hibooboo2.ultimateBravery.lolapi.Passive |  |
| PermLink | string |  |
| Picture | string |  |
| Skins | array |  |
| Stats | github.com.hibooboo2.ultimateBravery.lolapi.ChampStats |  |
| Tags | array |  |
| Title | string |  |

<a name="github.com.hibooboo2.ultimateBravery.lolapi.Image"></a>

#### Image

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Full | string |  |
| Group | string |  |
| Picture | string |  |

<a name="github.com.hibooboo2.ultimateBravery.lolapi.Passive"></a>

#### Passive

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Image | Image |  |
| Name | string |  |

<a name="github.com.hibooboo2.ultimateBravery.lolapi.Skin"></a>

#### Skin

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| ID | int |  |
| Name | string |  |
| Num | int |  |


