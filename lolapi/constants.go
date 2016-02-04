package lolapi

import "os"

const REALMS = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/realm?`
const ITEMS = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/item?itemListData=all&`
const ITEMS_BY_ID = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/item/%v?itemData=all&`
var API_KEY = os.Getenv("RIOT_API_KEY")
const ADD_KEY = "api_key="

