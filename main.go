package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"time"
	"github.com/hibooboo2/ultimateBravery/lolapi"
	"encoding/json"
	"strings"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	session, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(503)
		io.WriteString(w, "Errored")
	}

	_, err = r.Cookie("UB_UUID")
	if  err != nil {
		cookie := http.Cookie{
			Name:    "UB_UUID",
			Value:   session.String(),
			Path:    "/",
			Expires: time.Now().Add(time.Hour),
		}
		http.SetCookie(w, &cookie)
	}
	io.WriteString(w, `<!DOCTYPE html>
	 <meta http-equiv="refresh" content="5; URL=/">
	<body>
	Welcome to ultimateBravery!
	`)
	gotItems := lolapi.GetItems()
	io.WriteString(w, "<table>")
	for _, item := range gotItems {
		if item.Into == nil && item.From != nil && !strings.HasPrefix(item.Name, "Enchantment") {
			itemData, _ := json.MarshalIndent(item, "", "    ")
			io.WriteString(w, "<tr>")
			io.WriteString(w, "<td>")
			io.WriteString(w, fmt.Sprintf(`<img src="http://ddragon.leagueoflegends.com/cdn/6.2.1/img/item/%v" />`, item.Image.(map[string]interface{})["full"].(string)))
			io.WriteString(w, "</td>")
			io.WriteString(w, "<td>")
			io.WriteString(w, "<PRE>")
			io.WriteString(w, string(itemData))
			io.WriteString(w, "</PRE>")
			io.WriteString(w, "</td>")
			io.WriteString(w, "</tr>")
		}
	}
	io.WriteString(w, "</table>")
	io.WriteString(w, "</body>")
}

func shit(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Shit")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/shit", shit)
	http.ListenAndServe(":8000", mux)
}
