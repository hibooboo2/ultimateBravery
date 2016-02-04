package lolapi

type Spell struct {

}

type LOLBuild struct {
	Name string
	Item1 Item
	Item2 Item
	Item3 Item
	Item4 Item
	Item5 Item
	Item6 Item
	Summoner1 SummonerSpell
	Summoner2 SummonerSpell
	Champion  Champion
	SpellToMax Spell
}

func (theBuild *LOLBuild) TotalCost() int{
	return theBuild.Item1.Gold.Total +
	theBuild.Item2.Gold.Total +
	theBuild.Item3.Gold.Total +
	theBuild.Item4.Gold.Total +
	theBuild.Item5.Gold.Total +
	theBuild.Item6.Gold.Total
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
	}
}

func RandomBuildName() string {
	return "Some name"
}
