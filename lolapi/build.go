package lolapi

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"github.com/Sirupsen/logrus"
)

var pastBuilds = make(map[string]string)

type LOLBuild struct {
	Name       string
	Items      []*Item
	Summoner1  *SummonerSpell
	Summoner2  *SummonerSpell
	Champion   *Champion
	Map        *Map
	SpellToMax *Spell
	PermLink   string
	TotalCost  int
}

type LOLBuildLink struct {
	Name       string
	Items      []int
	Summoner1  string
	Summoner2  string
	Champion   int
	Map        string
	SpellToMax string
}

func (theLink *LOLBuildLink) getBuild() LOLBuild {
	items := []*Item{}
	for _, value := range theLink.Items {
		item := GetItemByID(value)
		item.Init()
		items = append(items, item)
	}
	build := LOLBuild{
		Name:     theLink.Name,
		Items:    items,
		Champion: GetChampionByID(theLink.Champion),
		Map:      allMapsMap[idStringToID(theLink.Map)],
	}
	build.init()
	return build
}

func (theLink *LOLBuild) getBuildLink() LOLBuildLink {
	items := []int{}
	for _, value := range theLink.Items {
		items = append(items, value.ID)
	}
	return LOLBuildLink{
		Name:     theLink.Name,
		Items:    items,
		Champion: theLink.Champion.ID,
		Map:      fmt.Sprint(theLink.Map.MapID),
	}
}

func (theBuild *LOLBuild) getLink() string {
	return MakeLink(theBuild.getBuildLink())
}

func (theBuild *LOLBuild) CalcTotalCost() int {
	sum := 0
	for _, item := range theBuild.Items {
		sum += item.Gold.Total
	}
	logrus.Debugf("Total Cost of Build: %v", sum)
	return sum
}

func (theBuild *LOLBuild) init() {
	for _, val := range theBuild.Items {
		val.Init()
	}
	buildLink := theBuild.getBuildLink()
	theBuild.PermLink = "/build/" + MakeLink(buildLink)
	theBuild.Champion.Init()
	theBuild.TotalCost = theBuild.CalcTotalCost()
}

func (theBuild *LOLBuild) Hash() string {
	itemsStrings := []string{}
	for _, value := range theBuild.Items {
		itemsStrings = append(itemsStrings, value.Name)
	}
	sort.Strings(itemsStrings)
	return theBuild.Champion.Name + ";" + strings.Join(itemsStrings, ";")
}

func RandomBuild() LOLBuild {
	theMap := RandomMap()
	champ := RandomChampion()
	build := LOLBuild{
		Name:     RandomBuildName(),
		Items:    RandomItemsFromMap(6, theMap, champ),
		Champion: champ,
		Map:      theMap,
	}
	build.init()
	return build
}

func RandomBraveryBuild(theMap *Map, champ *Champion) LOLBuild {
	used := true
	var build LOLBuild
	var hash string
	for used {
		build = LOLBuild{
			Name:     RandomBuildName(),
			Items:    RandomItemsFromMap(6, theMap, champ),
			Champion: champ,
			Map:      theMap,
		}
		hash = build.Hash()
		_, used = pastBuilds[hash]
		if used {
			logrus.Infof("Generated  a duplicate build. Rebuilding \n       < %v >", hash)
		}
		logrus.Debugf("Used %v Build : %v\n", used, hash)
	}
	build.init()
	pastBuilds[hash] = hash
	return build
}

func RandomBraveBuild() LOLBuild {
	return RandomBraveryBuild(RandomMap(), RandomChampion())
}

func RandomBuildName() string {
	return randomdata.SillyName()
}

func BuildFromLink(link string) *LOLBuild {
	x := FromLink(link, &LOLBuildLink{})
	if x != nil {
		build := x.(*LOLBuildLink).getBuild()
		build.init()
		return &build
	}
	return &LOLBuild{Name: "No Build"}
}
