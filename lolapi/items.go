package lolapi

import (
	"encoding/json"
	"fmt"
)

var AllItems = []Item{}
var allItemsMap = make(map[int]*Item)

type Gold struct {
	Base        int
	Purchasable bool
	Sell        int
	Total       int
}

type Item struct {
	Name                 string
	SanitizedDescription string
	Image                Image
	Description          string
	Plaintext            string
	Gold                 Gold
	Id                   int
	From                 []string
	FromItems            []*Item
	Into                 []string
	IntoItems            []*Item
	Depth                int
	Stats                interface{}
	Picture              string
	PermaLink            string
}

func (theItem *Item) CanUpgrade() bool {
	if len(theItem.Into) == 0 {
		return false
	}
	for _, v := range theItem.Into {
		if v != "" {
			return true
		}
	}
	return false
}

func (theItem *Item) CantUpgrade() bool {
	if len(theItem.Into) >= 1 {
		return false
	}
	return true
}

func (theItem *Item) IsAnUpgrade() bool {

	if len(theItem.From) == 0 {
		return false
	}
	for _, v := range theItem.From {
		if v != "" {
			return true
		}
	}
	return false
}

func (theItem *Item) Init() {
	theItem.Picture = ITEM_PICTURE + theItem.Image.Full
	theItem.FromItems = []*Item {}
	//for _, val := range theItem.From {
	//	id, err := strconv.Atoi(val)
	//	if err != nil {
	//		break
	//	}
	//	gotItem := GetItemById(id)
	//	if gotItem != nil {
	//		theItem.FromItems = append(theItem.FromItems, gotItem)
	//	}
	//}
	//theItem.IntoItems = []*Item {}
	//for _, val := range theItem.Into {
	//	id, err := strconv.Atoi(val)
	//	if err != nil {
	//		break
	//	}
	//	gotItem := GetItemById(id)
	//	if gotItem != nil {
	//		theItem.IntoItems = append(theItem.IntoItems, gotItem)
	//	}
	//}
	theItem.PermaLink = fmt.Sprintf("/items/%v",theItem.Id)
}

func InitializeItemsSlice() {
	items := getResource(ITEMS)
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems {
		var item Item
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &item)
		item.Init()
		AllItems = append(AllItems, item)
		allItemsMap[item.Id] = &item
	}
}

func RandomItem() *Item {
	itemsToUse := FullItems()
	return &itemsToUse[RandomNumber(len(itemsToUse)-1)]
}

func GetItemById(id int) *Item {
	return allItemsMap[id]
}

func FullItems() []Item {
	filteredItems := []Item{}
	for _, val := range AllItems {
		if val.CantUpgrade() && val.IsAnUpgrade() {
			filteredItems = append(filteredItems, val)
		}
	}
	return filteredItems
}
