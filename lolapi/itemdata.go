package lolapi

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	Id string
	MaxGroupOwnable string
}

type ItemData struct {
	Type string
	Version string
	Basic  map[string]interface{}
	Data   map[string]Item
	Groups []Group
}

var theItemData ItemData

func initializeItemsFromDataSlice() {
	data := getResource(ITEMS_JSON)
	jsonItem, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonItem, &theItemData)
	for _, item := range theItemData.Data {
		item.Init()
	}
	errs := []error{}
	for _, item := range theItemData.Data {
		err := item.Verify()
		if err != nil {
			errs = append(errs, err)
		}
	}
	fmt.Printf("%v", theItemData.Groups)
}

