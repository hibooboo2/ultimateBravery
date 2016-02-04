package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"time"
	"github.com/hibooboo2/ultimateBravery/lolapi"
	"html/template"
	"io/ioutil"
	"fmt"
)

var s1 = InitTemplates()

func main() {
	lolapi.Init()
	mux := http.NewServeMux()
	mux.HandleFunc("/", templateAttempt)
	mux.HandleFunc("/build/", build)
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

	item := lolapi.AllItems[lolapi.RandomNumber(len(lolapi.AllItems) - 1)]

	s1.ExecuteTemplate(w, "header", item)
	build  := lolapi.RandomBuild()
	build.Init()
	s1.ExecuteTemplate(w, "build", build)
	println(build.PermaLink)
	s1.ExecuteTemplate(w, "footer", nil)
}

func build(w http.ResponseWriter, r *http.Request) {
	println(r.URL)

}

func InitTemplates() *template.Template {
	files, _ := ioutil.ReadDir("./templates")
	fileNames := []string {}
	for _, f := range files {
		fileNames = append(fileNames, fmt.Sprint("./templates/", f.Name()))
	}

	s1, err := template.ParseFiles(fileNames...)
	if err != nil {
		panic(err)
	}
	return s1
}
