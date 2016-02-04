package lolapi

import "os"

const REALMS = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/realm?`
const ITEMS = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/item?itemListData=all&`
const ITEMS_BY_ID = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/item/%v?itemData=all&`
const ITEM_PICTURE = `http://ddragon.leagueoflegends.com/cdn/6.2.1/img/item/`
const CHAMPIONS = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/champion?champData=all&`
const CHAMPIONS_BY_ID = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/champion/%v?champData=all&`
const CHAMPION_PICTURE = `http://ddragon.leagueoflegends.com/cdn/6.2.1/img/champion/`
const CHAMPION_LOADING = `http://ddragon.leagueoflegends.com/cdn/img/champion/loading/%v_%v.jpg`
const CHAMPION_SPLASH = `http://ddragon.leagueoflegends.com/cdn/img/champion/splash/%v_%v.jpg`
var API_KEY = os.Getenv("RIOT_API_KEY")
const ADD_KEY = "api_key="

