package main

// handler functions

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/sendgrid/sendgrid-go"
	//"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// turn this into json
const jsonRes = `<!doctype html>
<html>
<head>
  <title>Static Files</title>
  <link rel="stylesheet" type="text/css" href="main.css">
</head>
<body>
  <p>This could be json.</p>
</body>
</html>`

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, jsonRes)
}

func emailHandler(){
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/api_keys", "https://api.sendgrid.com")
	request.Method = "GET"

	response, err := sendgrid.API(request)
	if err != nil {
		log.Println("error", err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
