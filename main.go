package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"time"
	"github.com/hibooboo2/ultimateBravery/lolapi"
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
	<meta http-equiv="refresh" content="5; URL=/">Welcome to ultimateBravery
	`)
}

func shit(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Shit")
}

func main() {
	lolapi.GetItems()
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", hello)
	//mux.HandleFunc("/shit", shit)
	//http.ListenAndServe(":8000", mux)
}
