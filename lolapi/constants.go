package lolapi

import "os"

var VERSION = func() string {
	x, err := getResource(`https://ddragon.leagueoflegends.com/api/versions.json`, false)
	if err != nil {
		panic(err)
	}
	return x.([]interface{})[0].(string)
}()

const DATA_DRAGON = `http://ddragon.leagueoflegends.com/cdn/`
const STATIC_DATA = `https://global.api.pvp.net/api/lol/static-data`

const REALMS = STATIC_DATA + `/na/v1.2/realm?`
const ITEMS = STATIC_DATA + `/na/v1.2/item?itemListData=all&`
const ITEMS_BY_ID = STATIC_DATA + `/na/v1.2/item/%v?itemData=all&`
const CHAMPIONS = STATIC_DATA + `/na/v1.2/champion?champData=all&`
const CHAMPIONS_BY_ID = STATIC_DATA + `/na/v1.2/champion/%v?champData=all&`
const CHAMPION_LOADING = DATA_DRAGON + `img/champion/loading/%v_%v.jpg`
const CHAMPION_SPLASH = DATA_DRAGON + `img/champion/splash/%v_%v.jpg`
const MAPS = STATIC_DATA + `/na/v1.2/map?`

var CHAMPION_PICTURE = DATA_DRAGON + VERSION + `/img/champion/`
var ITEMS_JSON = DATA_DRAGON + VERSION + `/data/en_US/item.json?`
var ITEM_PICTURE = DATA_DRAGON + VERSION + `/img/item/`
var SPELL_PICTURE = DATA_DRAGON + VERSION + `/img/spell/`

var API_KEY = os.Getenv("RIOT_API_KEY")

const ADD_KEY = "api_key="
