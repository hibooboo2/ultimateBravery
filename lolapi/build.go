package lolapi

import (
	"github.com/Pallinder/go-randomdata"
	"fmt"
)

type Spell struct {
}

type LOLBuild struct {
	Name       string
	Items      []Item
	Summoner1  *SummonerSpell
	Summoner2  *SummonerSpell
	Champion   *Champion
	Map        *Map
	SpellToMax *Spell
	PermLink  string
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
	items := []Item {}
	for _, value := range theLink.Items {
		items = append(items, GetItemById(value))
	}
	build := LOLBuild{
		Name:     theLink.Name,
		Items:    items,
		Champion: GetChampionById(theLink.Champion),
		Map: allMapsMap[idStringToId(theLink.Map)],
	}
	build.init()
	return build
}

func (theLink *LOLBuild) getBuildLink() LOLBuildLink {
	items := []int {}
	for _, value := range theLink.Items {
		items = append(items, value.Id)
	}
	return LOLBuildLink{
		Name:     theLink.Name,
		Items: items,
		Champion: theLink.Champion.Id,
		Map: fmt.Sprint(theLink.Map.MapId),
	}
}

func (theBuild *LOLBuild) TotalCost() int {
	sum := 0
	for _, item := range theBuild.Items {
		sum += item.Gold.Total
	}
	return sum
}

func (theBuild *LOLBuild) init() {
	theBuild.Name = theBuild.Name + " " + theBuild.Champion.Name
	for _, val := range theBuild.Items {
		val.Init()
	}
	buildLink := theBuild.getBuildLink()
	theBuild.PermLink = "/build/" + MakeLink(buildLink)
}

func RandomBuild() LOLBuild {
	theMap := RandomMap()
	build := LOLBuild{
		Name:     RandomBuildName(),
		Items: RandomItemsFromMap(6, theMap),
		Champion: RandomChampion(),
		Map: theMap,
	}
	build.init()
	return build
}

func RandomBraveryBuild(theMap *Map) LOLBuild {
	build := LOLBuild{
		Name:     RandomBuildName(),
		Items: RandomItemsFromMap(6, theMap),
		Champion: RandomChampion(),
		Map: theMap,

	}
	build.init()
	fmt.Printf("%##v", build)
	for _, value := range build.Items {
		fmt.Printf("%##v", value)
	}
	return build
}

func RandomBuildName() string {
	return randomdata.SillyName()
}

func BuildFromLink(link string) *LOLBuild {
	x := FromLink(link, &LOLBuildLink{})
	if x != nil {
		build := x.(*LOLBuildLink).getBuild()
		return &build
	}
	return &LOLBuild{}
}
