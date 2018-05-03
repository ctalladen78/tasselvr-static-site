package main

// main function
import (

	// jsoniter // json parser/decoder/encoder
	// fasthttp
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/appengine"

	log "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Static("/", "public")
	log.SetFormatter(&log.JSONFormatter{})

	// Serve static files from "static" directory.
	// http.Handle("/", http.FileServer(http.Dir("public")))
	//http.Handle("/public/", http.FileServer(http.Dir(".")))

	// get json
	// http.HandleFunc("/json", jsonStreamHandler) // https://cloud.google.com/appengine/docs/standard/go/issue-requests
	// http.HandleFunc("/json", handlers.EchoHandler)
	// send mail
	// http.HandleFunc("/email", handlers.EmailHandler) // https://cloud.google.com/appengine/docs/standard/go/issue-requests

	// testing rest framework
	// e := echo.New(":8080")
	http.Handle("/", e)

	appengine.Main()
}
