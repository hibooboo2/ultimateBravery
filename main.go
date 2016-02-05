package main

import (
	"fmt"
	"github.com/hibooboo2/ultimateBravery/lolapi"
	uuid "github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var s1 = InitTemplates()

func main() {
	go func(x *template.Template) {
		for {
			x = InitTemplates()
			time.Sleep(time.Second * 2)
		}
	}(s1)
	lolapi.Init()
	mux := http.NewServeMux()
	mux.HandleFunc("/", templateAttempt)
	mux.HandleFunc("/build/", build)
	mux.HandleFunc("/items/", allItems)
	http.ListenAndServe(":8000", mux)
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
	println(build.PermaLink)
	s1.ExecuteTemplate(w, "footer", nil)
}

func allItems(w http.ResponseWriter, r *http.Request) {
	item := lolapi.AllItems[lolapi.RandomNumber(len(lolapi.AllItems)-1)]
	s1.ExecuteTemplate(w, "header", item)
	s1.ExecuteTemplate(w, "content", lolapi.AllItems)
	s1.ExecuteTemplate(w, "footer", nil)
}

func build(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	fmt.Printf("%##v \n", split[2])
	s1.ExecuteTemplate(w, "header", nil)
	build := lolapi.BuildFromLink(split[2])
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
