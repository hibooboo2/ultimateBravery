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
	myMux.TheRouter.NotFoundHandler = http.HandlerFunc(process(notFound))
	myMux.Middle("/", templateAttempt).Name("Root")
	myMux.Middle("/build/{id:[0-9A-Za-z]+}", build).Name("build")
	staticFiles := http.StripPrefix("/static/",http.FileServer(http.Dir("./static/")))
	myMux.TheRouter.PathPrefix("/static/").Handler(staticFiles).Name("static")
	myMux.TheRouter.Handle("/favicon.ico", staticFiles)
	myMux.Middle("/items", allItems).Name("items")
	myMux.Middle("/items/{id:[0-9]+}", itemById).Name("itemById")
	myMux.Middle("/champion", allChamps).Name("Champs")
	myMux.Middle("/champions/{id:[0-9]+}", champById).Name("ChampById")
	myMux.TheRouter.HandleFunc("/json/build", json).Name("Json")
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
	build := lolapi.RandomBraveryBuild(theMap, lolapi.RandomChampion())
	s1.ExecuteTemplate(w, "build", build)
}

func allItems(w http.ResponseWriter, r *http.Request) {
	s1.ExecuteTemplate(w, "items", lolapi.AllItems)
}

func itemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := lolapi.GetItemByIdString(vars["id"])
	item.Init()
	s1.ExecuteTemplate(w, "item", item)
}

func allChamps(w http.ResponseWriter, r *http.Request) {
	s1.ExecuteTemplate(w, "champs", lolapi.AllChampions)
}

func champById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := lolapi.GetChampionByIdString(vars["id"])
	s1.ExecuteTemplate(w, "champion", item)
}

func build(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%##v", vars)
	build := lolapi.BuildFromLink(vars["buildLink"])
	s1.ExecuteTemplate(w, "build", *build)

}

func json(w http.ResponseWriter, r *http.Request) {
	theMap := lolapi.AllMaps[0]
	build := lolapi.RandomBraveryBuild(theMap, lolapi.RandomChampion())
	for _, item := range build.Items {
		for _, item := range item.IntoItems {
			item.IntoItems = nil
			item.FromItems = nil
		}
		for _, item := range item.FromItems {
			item.FromItems = nil
			item.IntoItems = nil
		}
	}
	s1.ExecuteTemplate(w, "raw", lolapi.Pretty(build))
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
	s1.ExecuteTemplate(w, "404", lolapi.Pretty(requestProxy))
}

func process(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter,r *http.Request) {
		defer s1.ExecuteTemplate(w, "footer", nil)
		defer r.Body.Close()
		if !strings.Contains(r.URL.Path, ".") {
			route := mux.CurrentRoute(r)
			if route != nil {
				fmt.Println(route.GetName())
			}
		}
		s1.ExecuteTemplate(w, "header", nil)
		next (w, r)

	}

}

type Router struct {
	TheRouter *mux.Router
}

func (r *Router) Middle(path string, f func(http.ResponseWriter,
*http.Request)) *mux.Route {
	return r.TheRouter.NewRoute().Path(path).HandlerFunc(process(f))
}
