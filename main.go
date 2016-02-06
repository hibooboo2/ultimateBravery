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
	"strings"
	"net/url"
)

var s1 = InitTemplates()

func main() {
	start := time.Now()
	lolapi.Init()
	fmt.Printf("Total to init: %v", time.Since(start))
	go func() {
		for {
			s1 = InitTemplates()
			time.Sleep(time.Second * 2)
		}
	}()
	mux := mux.NewRouter().StrictSlash(true)
	myMux := Router{
		TheRouter: mux,
	}
	myMux.TheRouter.NotFoundHandler = http.HandlerFunc(notFound)
	myMux.Middle("/", templateAttempt).Name("Root")
	myMux.Middle("/build/{id:[0-9A-Za-z]+}", build).Name("build")
	myMux.TheRouter.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",http.FileServer(http.Dir("./static/")))).Name("static")
	myMux.Middle("/items", allItems).Name("items")
	myMux.Middle("/items/{id:[0-9]+}", itemById).Name("itemById")
	http.Handle("/", mux)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		println(err.Error())
	}
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
	theMap := lolapi.AllMaps[0]
	build := lolapi.RandomBraveryBuild(theMap)
	s1.ExecuteTemplate(w, "header", build.Items[3])
	s1.ExecuteTemplate(w, "build", build)
	s1.ExecuteTemplate(w, "footer", nil)
}

func allItems(w http.ResponseWriter, r *http.Request) {
	s1.ExecuteTemplate(w, "header", lolapi.RandomItem(nil))
	s1.ExecuteTemplate(w, "content", lolapi.AllItems)
	s1.ExecuteTemplate(w, "footer", nil)
}

func itemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := lolapi.GetItemByIdString(vars["id"])
	s1.ExecuteTemplate(w, "header", item)
	s1.ExecuteTemplate(w, "item", item)
	s1.ExecuteTemplate(w, "footer", nil)
}

func build(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%##v", vars)
	s1.ExecuteTemplate(w, "header", lolapi.RandomItem(nil))
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

type RequestJson struct {
	Method string
	URL *url.URL
	Proto string
	Host string
	Vars map[string]string
}

func notFound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestProxy := &RequestJson{
		Method:r.Method,
		URL:r.URL,
		Proto:r.Proto,
		Host:r.Host,
		Vars: vars,
	}
	fmt.Println(vars["host"])
	s1.ExecuteTemplate(w, "header", nil)
	s1.ExecuteTemplate(w, "404", lolapi.Pretty(requestProxy))
	s1.ExecuteTemplate(w, "footer", nil)
}

func process(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter,r *http.Request) {
		defer next (w, r)
		if !strings.Contains(r.URL.Path, ".") {
			route := mux.CurrentRoute(r)
			if route != nil {
				fmt.Println(route.GetName())
			}
		}
	}

}

type Router struct {
	TheRouter *mux.Router
}

func (r *Router) Middle(path string, f func(http.ResponseWriter,
*http.Request)) *mux.Route {
	return r.TheRouter.NewRoute().Path(path).HandlerFunc(process(f))
}
