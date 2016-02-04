package lolapi

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Item struct {
	Name string
	SanitizedDescription string
	Image string
	Description string
	Plaintext string
	from []string
	Gold map[string]interface{}
	Id int
}

func getResource(resourceUrl string) {
	response, err :=http.Get(resourceUrl + ADD_KEY + API_KEY)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if  err != nil {
		panic(err)
	}
	//fmt.Print(string(body))
	var items interface{}
	err = json.Unmarshal(body, &items)
	if err != nil {
		panic(err)
	}
	keys := []string {}
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for key, value := range gotItems{
		keys = append(keys, key)
		var item Item
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &item)
		fmt.Printf("%v\n", item)
	}
}
