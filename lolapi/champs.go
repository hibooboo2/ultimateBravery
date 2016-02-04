package lolapi

import (
	"encoding/json"
	"fmt"
)

var AllChampions =  []Champion {}

type Skin struct {
	Id int
	Name string
	Num int
}

func (champ *Champion) getSkinPic(skinNumber int) string {
	if (skinNumber > len(champ.Skins)){
		return ""
	}
	return fmt.Sprintf(CHAMPION_LOADING,champ.Key, skinNumber)
}

type Champion struct {
	Id int
	Key string
	Name string
	Title string
	Image Image
	Skins []Skin
}

func RandomChampion() Champion {
	return AllChampions[RandomNumber(len(AllChampions) - 1)]

}


func InitializeChampionsSlice() {
	items := getResource(CHAMPIONS)
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems{
		var champ Champion
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &champ)
		AllChampions = append(AllChampions, champ)
	}
}
