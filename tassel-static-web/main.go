package main

// main function
import (
	"net/http"
	"log"
	"os"
	// jsoniter
	// fasthttp
	"google.golang.org/appengine"
	//"github.com/ant0ine/go-json-rest/rest" // rest server
	//"github.com/sendgrid/rest" // rest client
	
)

func main() {
	sendGridApiKey := os.Getenv("SENDGRID_API_KEY")
	// Serve static files from "static" directory.
	http.Handle("/", http.FileServer(http.Dir("public")))
	//http.Handle("/public/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/json", jsonHandler) 
	
	// testing rest framework
	log.Println(sendGridApiKey) 
	// send mail
	//emailHandler()
	appengine.Main()
}
