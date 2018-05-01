package main

// main function
import (
	"net/http"
	// jsoniter
	// fasthttp
	"google.golang.org/appengine"

	handlers "hello2/tassel-static-web/handlers"

	log "github.com/sirupsen/logrus"
	//"github.com/ant0ine/go-json-rest/rest" // rest server
	//"github.com/sendgrid/rest" // rest client
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	// Serve static files from "static" directory.
	http.Handle("/", http.FileServer(http.Dir("public")))
	//http.Handle("/public/", http.FileServer(http.Dir(".")))

	// get json
	// http.HandleFunc("/json", jsonStreamHandler) // https://cloud.google.com/appengine/docs/standard/go/issue-requests
	http.HandleFunc("/json", handlers.EchoHandler)
	// send mail
	http.HandleFunc("/email", handlers.EmailHandler) // https://cloud.google.com/appengine/docs/standard/go/issue-requests

	// testing rest framework

	appengine.Main()
}
