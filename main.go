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
	"github.com/Sirupsen/logrus"
)

var s1 = InitTemplates()
var totalBuilds = 0

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	start := time.Now()
	lolapi.Init()
	logrus.Debugf("Total to init: %v \n", time.Since(start))
	mux := mux.NewRouter().StrictSlash(true)
	myMux := Router{
		TheRouter: mux,
	}
	myMux.TheRouter.NotFoundHandler = http.HandlerFunc(process(notFound))
	myMux.Middle("/", generateBuildAndStore).Name("Root")
	myMux.Middle("/kill/server/", kill).Name("Kill")
	myMux.Middle("/build/{buildLink}", build).Name("build")
	staticFiles := http.StripPrefix("/static/",http.FileServer(http.Dir("./static/")))
	myMux.TheRouter.PathPrefix("/static/").Handler(staticFiles).Name("static")
	myMux.TheRouter.Handle("/favicon.ico", staticFiles)
	myMux.Middle("/items", allItems).Name("items")
	myMux.Middle("/items/{id:[0-9]+}", itemById).Name("itemById")
	myMux.Middle("/champions", allChamps).Name("Champs")
	myMux.Middle("/champions/{id:[0-9]+}", champById).Name("ChampById")
	myMux.TheRouter.HandleFunc("/json/build", json).Name("Json")
	http.Handle("/", mux)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		println(err.Error())
	}
}

func generateBuildAndStore(w http.ResponseWriter, r *http.Request) {
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
	for x := 0; x < 1; x++ {
		s1.ExecuteTemplate(w, "build", lolapi.RandomBraveBuild())
	}
	totalBuilds++
}

func allItems(w http.ResponseWriter, r *http.Request) {
	s1.ExecuteTemplate(w, "items", lolapi.AllItems)
}

func kill(w http.ResponseWriter, r *http.Request) {
	logrus.Fatalf("Closing total Request since start: %#v", totalBuilds)
}

func itemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := lolapi.GetItemByIdString(vars["id"])
	item.Init()
	s1.ExecuteTemplate(w, "itemById", item)
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
	logrus.Debugf("%##v", vars)
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
	s1.ExecuteTemplate(w, "404", lolapi.Pretty(requestProxy))
}

func process(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter,r *http.Request) {
		thePage := &PageInfo{}
		defer s1.ExecuteTemplate(w, "footer", thePage)
		defer r.Body.Close()
		handleCookies(w, r, thePage)
		routeName := "Home Page"
		if !strings.Contains(r.URL.Path, ".") {
			route := mux.CurrentRoute(r)
			if route != nil {
				logrus.Debugf("Route: %v %v",route.GetName(), mux.Vars(r))
				routeName = route.GetName()
			}
		}
		thePage.Name = routeName
		s1.ExecuteTemplate(w, "header", thePage)
		next (w, r)

	}

}

func handleCookies(w http.ResponseWriter, r *http.Request, thePage *PageInfo) {
	cookie, err := r.Cookie("IS_DEV")
	logrus.Debugf("%#v %v", cookie, err)
	dev := false
	if err != http.ErrNoCookie && cookie.Value == "true" {
		dev = true
	}
	query := r.URL.Query()
	val, ok := query["dev"]
	logrus.Debugf("Query : %v", query)
	if ok && val[0] == "true" {
		cookie := http.Cookie{
			Name:    "IS_DEV",
			Value:   "true",
			Path:    "/",
			Expires: time.Now().Add(time.Second * 3600 * 24),
		}
		http.SetCookie(w, &cookie)
		dev = true
	} else if ok && val[0] == "false" {
		cookie := http.Cookie{
			Name:    "IS_DEV",
			Value:   "",
			Path:    "/",
			Expires: time.Now().Add(time.Second * 3600 * 24),
		}
		http.SetCookie(w, &cookie)
		dev = false
	}
	val, ok = query["UI"]
	logrus.Debugf("Query : %v", query)
	theme := ""
	if ok && val[0] != "" {
		cookie := http.Cookie{
			Name:    "UI",
			Value:   val[0],
			Path:    "/",
			Expires: time.Now().Add(time.Second * 3600 * 24 * 365 * 100),
		}
		http.SetCookie(w, &cookie)
		theme = val[0]
	}

	cookie, err = r.Cookie("UI")
	if err != http.ErrNoCookie || ok {
		if theme == "" {
			theme = cookie.Value
		}
		logrus.Debugf("Theme is: %v", theme)
		switch theme {
			case "DARK":
				thePage.Dark = true
			case "LIGHT":
				thePage.Light = true
			default:
				thePage.Light = true
		}
	}
	thePage.Dev = dev
}

type PageInfo struct {
	Dev bool
	Name string
	Dark bool
	Light bool
}

type Router struct {
	TheRouter *mux.Router
}

func (r *Router) Middle(path string, f func(http.ResponseWriter,
*http.Request)) *mux.Route {
	return r.TheRouter.NewRoute().Path(path).HandlerFunc(process(f))
}
