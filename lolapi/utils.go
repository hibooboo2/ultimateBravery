package lolapi

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var allItems  = []Item {}

type Item struct {
	Name string
	SanitizedDescription string
	Image interface{}
	Description string
	Plaintext string
	from []string
	Gold map[string]interface{}
	Id int
	From interface{}
	Into interface{}
	Depth int
}

func getResource(resourceUrl string) []Item {
	if  len(allItems) > 0{
		return allItems
	}
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
	gotItems := items.(map[string]interface{})["data"].(map[string]interface{})
	for _, value := range gotItems{
		var item Item
		jsonItem, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(jsonItem, &item)
		allItems = append(allItems, item)
	}
	return allItems
}
