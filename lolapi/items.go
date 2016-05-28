package lolapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
)

var AllItems = []*Item{}
var allItemsMap = make(map[int]*Item)
var idsToIgnore = make(map[int]int)
var fullBoots = []*Item{}

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
	Group                string
	RequiredChampion     string
	HideFromAll          bool
	initialized          bool
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
	if theItem.HideFromAll {
		return false
	}
	if theItem.Group == "FlaskGroup" || theItem.Group == "JungleItems" {
		return false
	}
	if len(otherItems) > 0 && theItem.IsBoot() {
		for _, otherItem := range otherItems {
			if otherItem.IsBoot() {
				return false
			}
		}
	} else if len(otherItems) == 0 && !theItem.IsBoot() {
		return false
	}
	if !theItem.CanUseOnMap(theMap) {
		return false
	}
	for _, otherItem := range otherItems {
		if otherItem.Id == theItem.Id {
			return false
		}
	}
	if theItem.Group != "" {
		totalCanHave := 0
		for _, group := range theItemData.Groups {
			if theItem.Group == group.Id {
				if group.MaxGroupOwnable != "-1" && group.MaxGroupOwnable != "" {
					totalCanHave = idStringToId(group.MaxGroupOwnable)
				}
			}
		}
		if totalCanHave >= 1 && totalInBuild(otherItems, theItem.Group)+1 > totalCanHave {
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

func (theItem *Item) IsBoot() bool {
	return strings.Contains(theItem.Group, "Boot") || strings.Contains(theItem.Name, "Boots")
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
	if theItem == nil {
		panic("A nil Item.")
	}
	if theItem.initialized {
		return theItem
	}
	if theItem.partialInit() == nil {
		return nil
	}

	theItem.FromItems = []*Item{}
	for _, val := range theItem.From {
		id := idStringToId(val)
		gotItem := GetItemById(id)
		if gotItem.Id == DEFAULT_ITEM {
			temp, err := GetItemFromRiot(val)
			if err != nil {
				logrus.Error(err.Error())
				logrus.Errorf("From Item %v not found for item %#v", val, theItem)
			} else {
				gotItem = temp
			}
		}
		gotItem = gotItem.partialInit()
		if gotItem != nil {
			theItem.FromItems = append(theItem.FromItems, gotItem)
		} else {
			idsToIgnore[id] = id
		}
	}
	theItem.IntoItems = []*Item{}
	for _, val := range theItem.Into {
		id := idStringToId(val)
		gotItem := GetItemById(id)
		if gotItem.Id == DEFAULT_ITEM {
			temp, err := GetItemFromRiot(val)
			if err != nil {
				logrus.Error(err.Error())
				logrus.Errorf("Into Item %v not found for item %#v", val, theItem)
			} else {
				gotItem = temp
			}
		}
		gotItem.partialInit()
		if gotItem != nil {
			theItem.IntoItems = append(theItem.IntoItems, gotItem)
		} else {
			idsToIgnore[id] = id
		}
	}

	if theItem.Verify() == nil {
		theItem.initialized = true
		return theItem
	}
	return nil
}

func (theItem *Item) partialInit() *Item {
	if theItem == nil {
		logrus.Warn("Failed a partial Init")
		return nil
	}
	theItem.Picture = ITEM_PICTURE + theItem.Image.Full
	theItem.PermLink = fmt.Sprintf("/items/%v", theItem.Id)
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

var NIL_ITEM = fmt.Errorf("No Item Got Nil Pointer.")

func (theItem *Item) Verify() error {
	if theItem == nil {
		return NIL_ITEM
	}
	err := checkItems(theItem, theItem.Into, theItem.IntoItems)
	if err != nil {
		return err
	}
	err = checkItems(theItem, theItem.From, theItem.FromItems)
	if err != nil {
		return err
	}
	if theItem.PermLink == "" {
		return fmt.Errorf("No PermLink. Item: %v", theItem.Name)
	}
	return nil
}

func checkItems(theItem *Item, idStrings []string, items []*Item) error {
	if len(idStrings) > 0 {
		for _, idString := range idStrings {
			matched := false
			hasItem := false
			_, exists := idsToIgnore[idStringToId(idString)]
			if exists {
				break
			}
			for _, item := range items {
				if item != nil {
					hasItem = true
					if strconv.Itoa(item.Id) == idString {
						matched = true
						break
					}
				}
			}
			if hasItem && !matched {
				return fmt.Errorf("Id for items don't match. ITEM: %v %v \n\n %#v", theItem.Name, theItem.Id, items)
			}
		}
	}
	return nil
}
func init() {
	logrus.Infoln("Items.go init ran.")
	items, err := getResource(ITEMS, true)
	if err != nil {
		panic(err)
	}
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems {
		var item Item
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &item)
		//itemGot := GetItemFromRiot(strconv.Itoa(item.Id))
		//itemGot.partialInit()
		//item = *itemGot
		AllItems = append(AllItems, &item)
		allItemsMap[item.Id] = &item
	}
	for _, item := range AllItems {
		item.Init()
		err := item.Verify()
		if err == nil {
			allItemsMap[item.Id] = item
		} else {
			panic(err)
		}
	}
	AllItems = []*Item{}
	for _, item := range allItemsMap {
		AllItems = append(AllItems, item)
	}
	logrus.Debugf("Ignoring these id: %#v", idsToIgnore)
	for _, item := range AllItems {
		if item.IsBoot() && item.IsAnUpgrade() && item.CantUpgrade() {
			fullBoots = append(fullBoots, item)
		}
	}
	logrus.Debugf("There are %v full boots.", len(fullBoots))
}

func RandomItem(itemsToUse []*Item) *Item {
	if itemsToUse == nil {
		itemsToUse = AllItems
	}
	n := RandomNumber(len(itemsToUse) - 1)
	logrus.Printf("Len of items to use: %v\n", n)
	return itemsToUse[n]
}

func RandomItemFromMap(theMap *Map, otherItems []*Item, champ *Champion) *Item {

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
	total := 1
	items := []*Item{RandomItem(fullBoots)}
	for total < 6 {
		item := RandomItemFromMap(theMap, items, champ)
		items = append(items, item)
		total++
	}
	return items
}

func GetItemById(id int) *Item {
	item, keyExists := allItemsMap[id]
	if keyExists {
		return item.partialInit()
	}
	return &Item{
		Name:                 "Not Found",
		Id:                   id,
		SanitizedDescription: "Riots Api never returned this item.",
	}
}

const DEFAULT_ITEM = 99999

func GetItemByIdString(idString string) *Item {
	return GetItemById(idStringToId(idString))
}

func GetItemFromRiot(idString string) (*Item, error) {
	gotItem, err := getResource(fmt.Sprintf(ITEMS_BY_ID, idString), false)
	if gotItem == nil {
		id := idStringToId(idString)
		idsToIgnore[id] = id
		return nil, err
	}
	var item Item
	jsonItem, err := json.Marshal(gotItem)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonItem, &item)
	gotItem = item.partialInit()
	return &item, nil
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

func FullItemsSR() []*Item {
	filteredItems := []*Item{}
	for _, val := range AllItems {
		if val.CantUpgrade() && val.IsAnUpgrade() && val.Maps["11"] && val.RequiredChampion == "" {
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
