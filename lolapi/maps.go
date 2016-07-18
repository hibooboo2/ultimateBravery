package lolapi

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
)

type Map struct {
	MapID   int
	Image   Image
	MapName string
}

var AllMaps = []*Map{}

var allMapsMap = make(map[int]*Map)

func init() {
	logrus.Infoln("Maps init ran.")
	maps, err := getResource(MAPS, true)
	if err != nil {
		panic(err)
	}
	gotMaps := maps.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotMaps {
		var aMap Map
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &aMap)
		if aMap.MapID == 11 { // || aMap.MapID == 12 || aMap.MapID == 8 {
			AllMaps = append(AllMaps, &aMap)
			allMapsMap[aMap.MapID] = &aMap
		}
	}
	for _, val := range AllMaps {
		val.Init()
	}
	failedToVerify := []string{}
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
	logrus.Debug("%##v \n", theMap)
}

func RandomMap() *Map {
	return AllMaps[RandomNumber(len(AllMaps)-1)]
}
