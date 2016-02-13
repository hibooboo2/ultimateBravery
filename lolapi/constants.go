package lolapi

import "os"

const DATA_DRAGON = `http://ddragon.leagueoflegends.com/cdn/`
const STATIC_DATA = `https://global.api.pvp.net/api/lol/static-data`
const VERSION = `6.3.1`

const REALMS = STATIC_DATA + `/na/v1.2/realm?`
const ITEMS = STATIC_DATA + `/na/v1.2/item?itemListData=all&`
const ITEMS_JSON = DATA_DRAGON + VERSION + `/data/en_US/item.json?`
const ITEMS_BY_ID = STATIC_DATA + `/na/v1.2/item/%v?itemData=all&`
const ITEM_PICTURE = DATA_DRAGON + VERSION + `/img/item/`
const CHAMPIONS = STATIC_DATA + `/na/v1.2/champion?champData=all&`
const CHAMPIONS_BY_ID = STATIC_DATA + `/na/v1.2/champion/%v?champData=all&`
const CHAMPION_PICTURE = DATA_DRAGON + VERSION + `/img/champion/`
const CHAMPION_LOADING = DATA_DRAGON + `img/champion/loading/%v_%v.jpg`
const CHAMPION_SPLASH = DATA_DRAGON + `img/champion/splash/%v_%v.jpg`

const SPELL_PICTURE = DATA_DRAGON + VERSION + `/img/spell/`

const MAPS = STATIC_DATA + `/na/v1.2/map?`
var API_KEY = os.Getenv("RIOT_API_KEY")

const ADD_KEY = "api_key="
