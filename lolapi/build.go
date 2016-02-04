package lolapi

type Spell struct {

}

type LOLBuild struct {
	Name string `json:"omitempty"`
	Item1 Item `json:"omitempty"`
	Item2 Item `json:"omitempty"`
	Item3 Item `json:"omitempty"`
	Item4 Item `json:"omitempty"`
	Item5 Item `json:"omitempty"`
	Item6 Item `json:"omitempty"`
	Summoner1 SummonerSpell `json:"omitempty"`
	Summoner2 SummonerSpell `json:"omitempty"`
	Champion  Champion `json:"omitempty"`
	SpellToMax Spell `json:"omitempty"`
	PermaLink string `json:"omitempty"`
}

func (theBuild *LOLBuild) TotalCost() int{
	return theBuild.Item1.Gold.Total +
	theBuild.Item2.Gold.Total +
	theBuild.Item3.Gold.Total +
	theBuild.Item4.Gold.Total +
	theBuild.Item5.Gold.Total +
	theBuild.Item6.Gold.Total
}

func (theBuild *LOLBuild) Init(){
	theBuild.PermaLink = MakeLink(theBuild)
}

func RandomBuild() LOLBuild {
	return LOLBuild{
		Name: RandomBuildName(),
		Item1: RandomItem(),
		Item2: RandomItem(),
		Item3: RandomItem(),
		Item4: RandomItem(),
		Item5: RandomItem(),
		Item6: RandomItem(),
		Champion: RandomChampion(),
	}
}

func RandomBuildName() string {
	return "Some name"
}
