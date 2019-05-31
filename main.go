package main

import (
	//"fmt"
	"net/http"
	// jsoniter
	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	"google.golang.org/appengine"

	log "github.com/sirupsen/logrus"
	//"github.com/ant0ine/go-json-rest/rest" // rest server
	//"github.com/sendgrid/rest" // rest client
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	r := mux.NewRouter()
	r.HandleFunc("/json", EchoHandler).Methods("POST")
	r.HandleFunc("/api/optin", EmailHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	// http.ListenAndServe(":3000",r)	// set specific port instead of default gcp port
	http.Handle("/", r) // use gcp default port

	appengine.Main()
}
