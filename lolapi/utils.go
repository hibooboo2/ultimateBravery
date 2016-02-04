package lolapi

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"math/rand"
	"time"
	"encoding/base64"
)

type SummonerSpell struct {

}

func getResource(resourceUrl string) interface{} {
	response, err :=http.Get(resourceUrl + ADD_KEY + API_KEY)
	if err != nil || response.StatusCode >= 400 {
		fmt.Fprintln(os.Stderr, "Response from riot: " + response.Status)
		os.Exit(3)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if  err != nil {
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
	InitializeChampionsSlice()
	InitializeItemsSlice()
}

func MakeLink(object interface{}) (string){
	data, err := json.Marshal(object)
	if err != nil {
		return ""
	}
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)
	return str
}

func FromLink(object string,  objectType interface{}) interface{} {

	data, err := base64.StdEncoding.DecodeString(object)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	err = json.Unmarshal(data, objectType)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return objectType
}
