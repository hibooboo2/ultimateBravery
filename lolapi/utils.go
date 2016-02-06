package lolapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
	"strconv"
)

type SummonerSpell struct {
}

func getResource(resourceUrl string) interface{} {
	resourceUrl = resourceUrl + ADD_KEY + API_KEY
	response, err := http.Get(resourceUrl)
	if err != nil || response.StatusCode >= 400 {
		fmt.Fprintln(os.Stderr, "Response from riot: "+response.Status)
		fmt.Fprintln(os.Stderr, resourceUrl)
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
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func FromLink(object string, objectType interface{}) interface{} {

	data, err := base64.StdEncoding.DecodeString(object)
	if err != nil {
		fmt.Println("Decode error:", err)
		return nil
	}
	err = json.Unmarshal(data, objectType)
	if err != nil {
		fmt.Println("Unmarshal Error:", err)
		return nil
	}
	fmt.Printf("%##v \n", objectType)
	return objectType
}

func Pretty(object interface{}) string {
	jsonObject, err := json.MarshalIndent(object, "", "    ")
	if err == nil {
		return string(jsonObject)
	}
	fmt.Println(err.Error())
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
