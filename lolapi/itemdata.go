package lolapi

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
)

// Group ...
type Group struct {
	ID              string
	MaxGroupOwnable string
}

// ItemData ...
type ItemData struct {
	Type    string
	Version string
	Basic   map[string]interface{}
	Data    map[string]Item
	Groups  []Group
}

var theItemData ItemData

func init() {
	logrus.Infoln("Itemdata.go init ran.")
	data, err := getResource(ITEMS_JSON, false)
	if err != nil {
		panic(err)
	}
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
}
