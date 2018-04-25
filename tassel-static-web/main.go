package main

// main function
import (
	"net/http"
	// jsoniter
	// fasthttp
	"google.golang.org/appengine"
)

func main() {
	// Serve static files from "static" directory.
	http.Handle("/", http.FileServer(http.Dir("public")))
	//http.Handle("/public/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/json", jsonHandler)
	appengine.Main()
}
