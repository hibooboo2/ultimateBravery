package lolapi

import (
	"encoding/base64"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

type SummonerSpell struct {
}

var TotalResourceCalls = 0

type RiotError struct {
	Url    string
	Status int
}

func (riotError *RiotError) Error() string {
	return fmt.Sprintf("Status %v Resource: %v", riotError.Status, strings.Split(riotError.Url, API_KEY)[0])
}

var canHitRiot = true
var timeLeft = time.NewTimer(time.Second)

func getResource(resourceUrl string, addKey bool) (interface{}, error) {
	if !canHitRiot {
		return nil, fmt.Errorf("Rate limited exceeded try again after %v", <-timeLeft.C)
	}
	TotalResourceCalls = TotalResourceCalls + 1
	if addKey {
		resourceUrl = resourceUrl + ADD_KEY + API_KEY
	}
	response, err := http.Get(resourceUrl)
	if err != nil || response.StatusCode >= 400 {
		if response.StatusCode == 429 {
			canHitRiot = false
			timeLeft = time.AfterFunc(time.Second*600, func() { canHitRiot = true })
		}
		return nil, &RiotError{Url: resourceUrl, Status: response.StatusCode}
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
	return data, nil
}

func RandomNumber(max int) int {
	if max <= 0 {
		return 0
	}
	return rand.Intn(max)
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

func idStringToID(idString string) int {
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

type ReaderWriter struct {
	data []byte
}
