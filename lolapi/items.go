package lolapi

import (
	"encoding/json"
	"math/rand"
	"time"
)

var AllItems = []Item {}

type Gold struct {
	Base int
	Purchasable bool
	Sell int
	Total int
}

type Item struct {
	Name string
	SanitizedDescription string
	Image Image
	Description string
	Plaintext string
	from []string
	Gold Gold
	Id int
	From []string
	Into []string
	Depth int
	Stats interface{}
}


type Image struct {
	Full string
	Group string
}


func (theItem *Item) CanUpgrade() bool {
	if len(theItem.Into) == 0 {
		return false;
	}
	for _, v := range theItem.Into {
		if v != "" {
			println(theItem.Name + " " + v)
			return true
		}
	}
	return false
}


func (theItem *Item) CantUpgrade() bool {
	if len(theItem.Into) >= 1 {
		return false;
	}
	return true
}

func (theItem *Item) IsAnUpgrade() bool {

	if len(theItem.From) == 0 {
		return false;
	}
	for _, v := range theItem.From {
		if v != "" {
			return true
		}
	}
	return false
}



func InitializeItemsSlice() []Item {
	rand.Seed(time.Now().Unix())

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

func RandomItem() Item {
	return 	AllItems[RandomNumber(len(AllItems) - 1)]
}
