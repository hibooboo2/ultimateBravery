package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"time"
	"github.com/hibooboo2/ultimateBravery/lolapi"
	"encoding/json"
	"html/template"
)

func main() {
	lolapi.InitializeItemsSlice()
	mux := http.NewServeMux()
	mux.HandleFunc("/", templateAttempt)
	http.ListenAndServe(":8000", mux)
}


func templateAttempt(w http.ResponseWriter, r *http.Request) {
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

	s1, err := template.ParseFiles("header.tmpl","footer.tmpl","item.tmpl", "itemInBuild.tmpl","content.tmpl", "build.tmpl")
	if err != nil {
		panic(err)
	}
	item := lolapi.AllItems[lolapi.RandomNumber(len(lolapi.AllItems) - 1)]
	displayItems := []lolapi.Item {}
	for _, val := range lolapi.AllItems {
		if val.CantUpgrade() && val.IsAnUpgrade() {
			_, _ = json.MarshalIndent(val, "", "    ")
			displayItems = append(displayItems, val)
		}
	}

	s1.ExecuteTemplate(w, "header", item)
	s1.ExecuteTemplate(w, "build", lolapi.RandomBuild())
	s1.ExecuteTemplate(w, "footer", nil)
}
