package lolapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var AllItems = []*Item{}
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
	PermLink             string
	Maps                 map[string]bool
	Group               string
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

func (theItem *Item) CanUseOnMap(theMap *Map) bool {
	canUse := theItem.Maps[strconv.Itoa(theMap.MapId)]
	return canUse
}

func (theItem *Item) CanUseInBuild(theMap *Map, otherItems []*Item, champ *Champion) bool {
	if !theItem.CanUseOnMap(theMap) {
		return false
	}
	for _, otherItem := range otherItems {
		if otherItem.Id == theItem.Id {
			return false
		}
	}
	if theItem.CanUpgrade() || !theItem.IsAnUpgrade() {
		return false
	}
	if !champ.CanUseItem(theItem) {
		return false
	}
	return true
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

func (theItem *Item) Init() *Item {
	theItem.partialInit()
	if  len(theItem.FromItems) == 0 {
		theItem.FromItems = []*Item {}
		for _, val := range theItem.From {
			gotItem := GetItemByIdString(val)
			theItem.FromItems = append(theItem.FromItems, gotItem)
		}
		theItem.IntoItems = []*Item {}
		for _, val := range theItem.Into {
			gotItem := GetItemByIdString(val)
			theItem.IntoItems = append(theItem.IntoItems, gotItem)
		}

	}
	if theItem.Verify() == nil {
		return theItem
	}
	return nil
}

func (theItem *Item) partialInit() *Item {
	if theItem == nil {
		return &Item{
			Name: "Item non existent",
			Id: -111,
			Image: Image{
				Full: "3751.png",
			},
		}
	}
	theItem.Picture = ITEM_PICTURE + theItem.Image.Full
	theItem.PermLink = fmt.Sprintf("/items/%v",theItem.Id)
	return theItem
}

func (theItem *Item) PrintSimple() {
	theItem.Init()
	Pretty(theItem)
	for _, item := range theItem.FromItems {
		Pretty(item)
	}
	for _, item := range theItem.IntoItems {
		Pretty(item)
	}
}

func (theItem *Item) Verify() error {
	if theItem == nil {
		return fmt.Errorf("No Item Got Nil Pointer.")
	}
	if len(theItem.Into) > 0 {
		for _, idString := range theItem.Into {
			matched := false
			for _, item := range theItem.IntoItems {
				if item != nil && item.Id == idStringToId(idString) {
					matched = true
					break
				}
			}
			if !matched {
				return fmt.Errorf("Id for items don't match. ITEM: %##v", theItem)
			}
		}
	}
	if len(theItem.From) > 0 {
		for _, idString := range theItem.From {
			matched := false
			for _, item := range theItem.FromItems {
				if item != nil && item.Id == idStringToId(idString) {
					matched = true
					break
				}
			}
			if !matched {
				return fmt.Errorf("Id for items don't match. ITEM: %##v", theItem)
			}
		}
	}
	if theItem.PermLink == "" {
		return fmt.Errorf("No PermLink. Item: %v", theItem.Name)
	}
	return nil
}

func initializeItemsSlice() {
	items := getResource(ITEMS)
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems {
		var item Item
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &item)
		AllItems = append(AllItems, &item)
		allItemsMap[item.Id] = &item
	}
	for _, item := range AllItems {
		item.Init()
	}
	errs := []error{}
	for _, item := range AllItems {
		err := item.Verify()
		if err != nil {
			errs = append(errs, err)
		}
	}
	fmt.Printf("%#v \n", errs)
}

func RandomItem(itemsToUse []*Item) *Item {
	if itemsToUse == nil {
		itemsToUse = AllItems
	}
	return itemsToUse[RandomNumber(len(itemsToUse)-1)]
}

func RandomItemFromMap(theMap *Map, otherItems []*Item, champ * Champion) *Item {

	item := AllItems[RandomNumber(len(AllItems)-1)]
	for !item.CanUseInBuild(theMap, otherItems, champ) {
		item = AllItems[RandomNumber(len(AllItems)-1)]
	}
	item.Init()
	return item
}

func RandomItemsFromMap(howMany int, theMap *Map, champ *Champion) []*Item {
	if theMap == nil {
		theMap = RandomMap()
	}
	total := 0
	items := []*Item {}
	for total < 6 {
		item := RandomItemFromMap(theMap, items, champ)
		items = append(items, item)
		total++
	}
	return items
}

func GetItemById(id int) *Item {
	item := allItemsMap[id]
	item.partialInit()
	return item
}

func GetItemByIdString(idString string) *Item {
	item := allItemsMap[idStringToId(idString)]
	item.partialInit()
	return item
}

func GetItemFromRiot(idString string) *Item {
	gotItem := getResource(fmt.Sprintf(ITEMS_BY_ID, idString))
	var item Item
	jsonItem, err := json.Marshal(gotItem)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonItem, &item)
	item.partialInit()
	return &item
}

func FullItems() []*Item {
	filteredItems := []*Item{}
	for _, val := range AllItems {
		if val.CantUpgrade() && val.IsAnUpgrade() {
			filteredItems = append(filteredItems, val)
		}
	}
	return filteredItems
}

func MapItems(theMap *Map) []*Item {
	filteredItems := []*Item{}
	for _, val := range AllItems {
		if val.CantUpgrade() && val.IsAnUpgrade() && val.Maps[strconv.Itoa(theMap.MapId)] {
			filteredItems = append(filteredItems, val)
		}
	}
	return filteredItems
}
