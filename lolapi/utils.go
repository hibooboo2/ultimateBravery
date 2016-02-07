package lolapi

import (
	"encoding/base64"
	"encoding/json"

	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"strconv"
	"github.com/Sirupsen/logrus"
)

type SummonerSpell struct {
}

var TotalResourceCalls = 0

func getResource(resourceUrl string) interface{} {
	TotalResourceCalls = TotalResourceCalls + 1
	resourceUrl = resourceUrl + ADD_KEY + API_KEY
	response, err := http.Get(resourceUrl)
	if err != nil || response.StatusCode >= 400 {
		logrus.Errorf("Response from riot: %v \n%v",response.Status ,resourceUrl)
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func RandomNumber(max int) int {
	if max <= 0 {
		return 0
	}
	return rand.Intn(max)
}

func Init() {
	defer func (){
		logrus.Debugf("Made %#v requests to riot.", TotalResourceCalls)
	}()
	rand.Seed(time.Now().Unix())
	initializeChampionsSlice()
	initializeItemsSlice()
	initializeItemsFromDataSlice()
	initializeMaps()
}

func MakeLink(object interface{}) string {
	data, err := json.Marshal(object)
	if err != nil {
		println("Failed to make a link.")
		return ""
	}
	str := base64.URLEncoding.EncodeToString(data)
	return str
}

func FromLink(object string, objectType interface{}) interface{} {

	data, err := base64.URLEncoding.DecodeString(object)
	if err != nil {
		logrus.Println("Decode error:", err)
		return nil
	}
	err = json.Unmarshal(data, objectType)
	if err != nil {
		logrus.Println("Unmarshal Error:", err)
		return nil
	}
	logrus.Printf("%##v \n", objectType)
	return objectType
}

func Pretty(object interface{}) string {
	jsonObject, err := json.MarshalIndent(object, "", "    ")
	if err == nil {
		return string(jsonObject)
	}
	logrus.Println(err.Error())
	return ""
}

func idStringToId(idString string) int {
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(idString + err.Error())
	}
	return id
}

func totalInBuild(items []*Item, group string) int {
	total := 0
	for _, value := range items {
		if value.Group == group {
			total++
		}
	}
	return total
}


type ReaderWriter struct{
	data []byte
}
