package lolapi

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
)

var AllChampions = []*Champion{}
var allChampsMap = make(map[int]*Champion)

type Skin struct {
	Id   int
	Name string
	Num  int
}

func (champ *Champion) getSkinPic(skinNumber int) string {
	if skinNumber > len(champ.Skins) {
		return ""
	}
	return fmt.Sprintf(CHAMPION_LOADING, champ.Key, skinNumber)
}

func (champ *Champion) CanUseItem(theItem *Item) bool {
	//TODO: Implement this logic.
	if theItem.RequiredChampion != "" && theItem.RequiredChampion != champ.Name {
		return false
	}
	return true
}

func (champ *Champion) init() *Champion {
	champ.Picture = CHAMPION_PICTURE + champ.Image.Full
	champ.PermLink = fmt.Sprintf("/champions/%v", champ.Id)
	return champ
}

type Champion struct {
	Id    int
	Key   string
	Name  string
	Title string
	Image Image
	Skins []Skin
	Picture string
	PermLink string
}

func RandomChampion() *Champion {
	return AllChampions[RandomNumber(len(AllChampions)-1)]

}

func initializeChampionsSlice() {
	items := getResource(CHAMPIONS)
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems {
		var champ Champion
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &champ)
		champ.init()
		AllChampions = append(AllChampions, &champ)
		allChampsMap[champ.Id] = &champ
	}
	logrus.Debugf("Total champions: %v", len(AllChampions))
}

func GetChampionById(id int) *Champion {
	champ := allChampsMap[id]
	champ.init()
	return champ
}

func GetChampionByIdString(id string) *Champion {
	return GetChampionById(idStringToId(id))
}
