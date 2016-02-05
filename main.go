package main

import (
	"fmt"
	"github.com/hibooboo2/ultimateBravery/lolapi"
	"github.com/gorilla/mux"
	uuid "github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"
)

var s1 = InitTemplates()

func main() {
	go func() {
		for {
			s1 = InitTemplates()
			time.Sleep(time.Second * 2)
		}
	}()
	lolapi.Init()
	mux := mux.NewRouter()
	mux.HandleFunc("/*", templateAttempt)
	mux.HandleFunc("/build/{id:[0-9A-Za-z]+}", build)
	mux.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
	mux.HandleFunc("/items", allItems)
	mux.HandleFunc("/items/", allItems)
	mux.HandleFunc("/items/{id:[0-9]+}", itemById)
	http.Handle("/", mux)
	http.ListenAndServe(":8000", nil)
}

func templateAttempt(w http.ResponseWriter, r *http.Request) {
	session, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(503)
		io.WriteString(w, "Errored")
	}

	_, err = r.Cookie("UB_UUID")
	if err != nil {
		cookie := http.Cookie{
			Name:    "UB_UUID",
			Value:   session.String(),
			Path:    "/",
			Expires: time.Now().Add(time.Hour),
		}
		http.SetCookie(w, &cookie)
	}

	item := lolapi.AllItems[lolapi.RandomNumber(len(lolapi.AllItems)-1)]

	s1.ExecuteTemplate(w, "header", item)
	build := lolapi.RandomBuild()
	build.Init()
	s1.ExecuteTemplate(w, "build", build)
	s1.ExecuteTemplate(w, "footer", nil)
}

func allItems(w http.ResponseWriter, r *http.Request) {
	item := lolapi.AllItems[lolapi.RandomNumber(len(lolapi.AllItems)-1)]
	s1.ExecuteTemplate(w, "header", item)
	s1.ExecuteTemplate(w, "content", lolapi.AllItems)
	s1.ExecuteTemplate(w, "footer", nil)
}

func itemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s1.ExecuteTemplate(w, "error", err)
	}
	item := lolapi.GetItemById(id)
	s1.ExecuteTemplate(w, "header", item)
	s1.ExecuteTemplate(w, "item", item)
	s1.ExecuteTemplate(w, "footer", nil)
}

func build(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s1.ExecuteTemplate(w, "header", nil)
	build := lolapi.BuildFromLink(vars["buildLink"])
	s1.ExecuteTemplate(w, "build", *build)
	s1.ExecuteTemplate(w, "footer", nil)

}

func InitTemplates() *template.Template {
	files, _ := ioutil.ReadDir("./templates")
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, fmt.Sprint("./templates/", f.Name()))
	}

	s1, err := template.ParseFiles(fileNames...)
	if err != nil {
		panic(err)
	}
	return s1
}
