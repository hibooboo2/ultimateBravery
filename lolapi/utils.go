package lolapi

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

var AllItems = []Item {}

type Image struct {
	Full string
	Group string
}

type Gold struct {
	Base int
	Purchasable bool
	Sell int
	Total int
}

type Item struct {
	Name string
	SanitizedDescription string
	Image Image
	Description string
	Plaintext string
	from []string
	Gold Gold
	Id int
	From []string
	Into []string
	Depth int
	Stats interface{}
}


func (theItem *Item) CanUpgrade() bool {
	if len(theItem.Into) == 0 {
		return false;
	}
	for _, v := range theItem.Into {
		if v != "" {
			return true
		}
	}
	return false
}

func (theItem *Item) IsUpgrade() bool {

	if len(theItem.From) == 0 {
		return false;
	}
	for _, v := range theItem.From {
		if v != "" {
			return true
		}
	}
	return false
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
