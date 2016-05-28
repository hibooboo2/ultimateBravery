package lolapi

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
)

var AllChampions = []*Champion{}
var allChampsMap = make(map[int]*Champion)
var shuffle = []*Champion{}

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

func (champ *Champion) Init() *Champion {
	champ.Picture = CHAMPION_PICTURE + champ.Image.Full
	champ.PermLink = fmt.Sprintf("/champions/%v", champ.Id)
	for _, spell := range champ.Spells {
		spell.Image.Init()
		logrus.Debugf("Pic: %v", spell.Image.Picture)
	}
	return champ
}

type Champion struct {
	Id        int
	Key       string
	Name      string
	Title     string
	Image     Image
	Skins     []*Skin
	Lore      string
	Blurb     string
	Allytips  []string
	Enemytips []string
	Tags      []string
	Partype   string
	Info      ChampInfo
	Stats     ChampStats
	Spells    []*Spell
	Passive   Passive
	Picture   string
	PermLink  string
}

type Passive struct {
	Name, Description, SanitizedDescription string
	Image                                   Image
}

type Spell struct {
	Name                 string
	Description          string
	SanitizedDescription string
	Tooltip              string
	SanitizedTooltip     string
	Leveltip             string
	Image                Image
	Resource             string
	Maxrank              int
	Cost                 []int
	CostType             string
	CostBurn             string
	Cooldown             []int
	CooldownBurn         string
	Effect               [][]int
	EffectBurn           []string
	Range                []int
	RangeBurn            string
	key                  string
}

type ChampInfo struct {
	Attack     int
	Defense    int
	Magic      int
	Difficulty int
}

type ChampStats struct {
	Armor                int
	Armorperlevel        int
	Attackdamage         int
	Attackdamageperlevel int
	Attackrange          int
	Attackspeedoffset    int
	Attackspeedperlevel  int
	Crit                 int
	Critperlevel         int
	Hp                   int
	Hpperlevel           int
	Hpregen              int
	Hpregenperlevel      int
	Movespeed            int
	Mp                   int
	Mpperlevel           int
	Mpregen              int
	Mpregenperlevel      int
	Spellblock           int
	Spellblockperlevel   int
}

func RandomChampion() *Champion {
	index := RandomNumber(len(shuffle) - 1)
	champ := shuffle[index]
	logrus.Debugf("Left in shuffle: %v", len(shuffle))
	if len(shuffle) > 1 {
		shuffle = append(shuffle[:index], shuffle[index+1:]...)
	} else if len(shuffle) <= 1 {
		logrus.Debug("Resetting shuffle.")
		shuffle = []*Champion{}
		for _, val := range allChampsMap {
			shuffle = append(shuffle, val)
		}
	}
	champ.Init()
	return champ

}

func init() {
	logrus.Infoln("Champs init ran.")
	items, err := getResource(CHAMPIONS, true)
	if err != nil {
		logrus.Fatalf("Failed to get items: %#v", err)
	}
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems {
		var champ Champion
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &champ)
		champ.Init()
		AllChampions = append(AllChampions, &champ)
		allChampsMap[champ.Id] = &champ
	}
	logrus.Debugf("Total champions: %v", len(AllChampions))

	shuffle = []*Champion{}
	for _, val := range allChampsMap {
		shuffle = append(shuffle, val)
	}
}

func GetChampionById(id int) *Champion {
	champ := allChampsMap[id]
	champ.Init()
	return champ
}

func GetChampionByIdString(id string) *Champion {
	return GetChampionById(idStringToId(id))
}
