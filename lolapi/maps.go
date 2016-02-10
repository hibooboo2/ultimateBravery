package lolapi

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
)

type Map struct {
	MapId int
	Image Image
	MapName string
}

var AllMaps = []*Map {}

var allMapsMap = make(map[int]*Map)

func initializeMaps() {
	items, err := getResource(MAPS)
	if err != nil {
		panic(err)
	}
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems {
		var aMap Map
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &aMap)
		if aMap.MapId == 11 { // || aMap.MapId == 12 || aMap.MapId == 8 {
			AllMaps = append(AllMaps, &aMap)
			allMapsMap[aMap.MapId] = &aMap
		}
	}
	for _, val := range AllMaps {
		val.Init()
	}
	failedToVerify := []string {}
	for _, item := range AllMaps {
		err := item.Verify()
		if err != nil {
			failedToVerify = append(failedToVerify, item.MapName)
		}
	}
	logrus.Warnf("%##v \n", failedToVerify)
}

func (theMap *Map) Verify() error {
	return nil
}

func (theMap *Map) Init() {
	logrus.Debug("%##v \n",theMap)
}

func RandomMap() *Map {
	return AllMaps[RandomNumber(len(AllMaps)-1)]
}
