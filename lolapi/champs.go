package lolapi

import (
	"os"
	"encoding/json"
)

const REALMS = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/realm?`
const ITEMS = `https://global.api.pvp.net/api/lol/static-data/na/v1.2/item?itemListData=all&`
var API_KEY = os.Getenv("RIOT_API_KEY")
const ADD_KEY = "api_key="

func InitializeItemsSlice() []Item {
	items := getResource(ITEMS)
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems{
		var item Item
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &item)
		AllItems = append(AllItems, item)
	}
	return AllItems
}
