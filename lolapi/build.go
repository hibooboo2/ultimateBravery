package lolapi

import (
	"github.com/Pallinder/go-randomdata"
)

type Spell struct {
}

type LOLBuild struct {
	Name       string
	Item1      *Item
	Item2      *Item
	Item3      *Item
	Item4      *Item
	Item5      *Item
	Item6      *Item
	Summoner1  *SummonerSpell
	Summoner2  *SummonerSpell
	Champion   *Champion
	SpellToMax *Spell
	PermaLink  string
}

func (theBuild *LOLBuild) TotalCost() int {
	return theBuild.Item1.Gold.Total +
		theBuild.Item2.Gold.Total +
		theBuild.Item3.Gold.Total +
		theBuild.Item4.Gold.Total +
		theBuild.Item5.Gold.Total +
		theBuild.Item6.Gold.Total
}

func (theBuild *LOLBuild) Init() {
	theBuild.Name = theBuild.Name + " " + theBuild.Champion.Name
	buildLink := theBuild.getBuildLink()
	theBuild.PermaLink = "/build/" + MakeLink(buildLink)
}

func RandomBuild() LOLBuild {
	return LOLBuild{
		Name:     RandomBuildName(),
		Item1:    RandomItem(),
		Item2:    RandomItem(),
		Item3:    RandomItem(),
		Item4:    RandomItem(),
		Item5:    RandomItem(),
		Item6:    RandomItem(),
		Champion: RandomChampion(),
	}
}

type LOLBuildLink struct {
	Name       string
	Item1      int
	Item2      int
	Item3      int
	Item4      int
	Item5      int
	Item6      int
	Summoner1  string
	Summoner2  string
	Champion   int
	SpellToMax string
}

func (theLink *LOLBuildLink) getBuild() LOLBuild {
	return LOLBuild{
		Name:     theLink.Name,
		Item1:    GetItemById(theLink.Item1),
		Item2:    GetItemById(theLink.Item2),
		Item3:    GetItemById(theLink.Item3),
		Item4:    GetItemById(theLink.Item4),
		Item5:    GetItemById(theLink.Item5),
		Item6:    GetItemById(theLink.Item6),
		Champion: GetChampionById(theLink.Champion),
	}
}

func (theLink *LOLBuild) getBuildLink() LOLBuildLink {
	return LOLBuildLink{
		Name:     theLink.Name,
		Item1:    theLink.Item1.Id,
		Item2:    theLink.Item2.Id,
		Item3:    theLink.Item3.Id,
		Item4:    theLink.Item4.Id,
		Item5:    theLink.Item5.Id,
		Item6:    theLink.Item6.Id,
		Champion: theLink.Champion.Id,
	}
}

func RandomBuildName() string {
	return randomdata.SillyName()
}

func BuildFromLink(link string) *LOLBuild {
	x := FromLink(link, &LOLBuildLink{})
	build := x.(*LOLBuildLink).getBuild()
	return &build
}
