package main

// handler functions

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
	//"github.com/sendgrid/sendgrid-go"
	//"github.com/sendgrid/sendgrid-go/helpers/mail"
	//"google.golang.org/appengine"
    //    "google.golang.org/appengine/urlfetch" 
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
	/*
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/api_keys", "https://api.sendgrid.com")
	request.Method = "GET"

	response, err := sendgrid.API(request) // instead of using sendgrid use the appengine urlfetch
	
	
	 // ctx := appengine.NewContext(r)
     //   client := urlfetch.Client(ctx)
      //  resp, err := client.Get("https://www.google.com/")
	 
	if err != nil {
		log.Println("error", err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	*/
	url := "https://api.sendgrid.com/v3/mail/send"

	payload := strings.NewReader("{\"personalizations\":[{\"to\":[{\"email\":\"john.doe@example.com\",\"name\":\"John Doe\"}],\"subject\":\"Hello, World!\"}],\"from\":{\"email\":\"sam.smith@example.com\",\"name\":\"Sam Smith\"},\"reply_to\":{\"email\":\"sam.smith@example.com\",\"name\":\"Sam Smith\"}}")

	req, _ := http.NewRequest("POST", url, payload)
	

	
	auth := "Bearer " + os.Getenv("SENDGRID_API_KEY")
	
	log.Println("auth ", auth)
	req.Header.Add("authorization", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
