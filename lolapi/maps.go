package lolapi

import (
	"encoding/json"
	"fmt"
)

type Map struct {
	MapId int
	Image Image
	MapName string
}

var AllMaps = []*Map {}

var allMapsMap = make(map[int]*Map)

func initializeMaps() {
	items := getResource(MAPS)
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems {
		var aMap Map
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &aMap)
		if aMap.MapId == 11 { //|| aMap.MapId == 12 || aMap.MapId == 8 {
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
	fmt.Printf("%##v \n", failedToVerify)
}

func (theMap *Map) Verify() error {
	return nil
}

func (theMap *Map) Init() {
	fmt.Printf("%##v \n",theMap)
}

func RandomMap() *Map {
	return AllMaps[RandomNumber(len(AllMaps)-1)]
}
