package main

// handler functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// explicit typing
type emailTemplate struct {
	personalizations string
	from             []string // map [string]interface{}
	reply_to         []string
}

// dynamic typing using empty interface
const jsonExample = `{ 'personalizations': 
	[ { 'to': [ { email: 'john.doe@example.com', 'name': 'John Doe' } ],
		'subject': 'Hello, World!' } ],
   'from': { 'email': 'sam.smith@example.com', 'name': 'Sam Smith' },
   'reply_to': { 'email': 'sam.smith@example.com', 'name': 'Sam Smith' } }`

type Animal struct {
	Name  string
	Order string
}

// type Animals struct {
// 	[Animal]interface{}
// }

// https://blog.golang.org/json-and-go
//https://golang.org/src/encoding/json/example_test.go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	var jsonBlob = []byte(`{[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]}`)
	var animals []Animal

	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	log.Fatal("json example: ", animals)
	// Output:
	// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
}

func jsonHandler1(w http.ResponseWriter, r *http.Request) {

	const jsonStream = `
	{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	log.Fatal("json example", dec)

}

// echoHandler reads a JSON object from the body, and writes it back out.
// https://gist.github.com/campoy/7b44f6ec2d9e82d956d34b4989b33192
func echoHandler(w http.ResponseWriter, r *http.Request) {
	var msg interface{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			// log.Fatal(w, http.StatusBadRequest, "Body was not valid JSON: %v", err)
			log.Fatal("Body was not valid JSON: %v", err)
			return
		}
		// log.Fatal(w, http.StatusInternalServerError, "Could not get body: %v", err)
		log.Fatal("Could not get body: %v", err)
		return
	}

	b, err := json.Marshal(msg)
	if err != nil {
		// log.Fatal(w, http.StatusInternalServerError, "Could not marshal JSON: %v", err)
		log.Fatal("Could not marshal JSON: %v", err)
		return
	}
	w.Write(b)
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.sendgrid.com/v3/mail/send"

	// must be in JSON
	payload := strings.NewReader("{\"personalizations\":[{\"to\":[{\"email\":\"cyrus@tasselvr.com\",\"name\":\"John Doe\"}],\"subject\":\"Hello, World!\"}],\"from\":{\"email\":\"sam.smith@example.com\",\"name\":\"Sam Smith\"},\"reply_to\":{\"email\":\"sam.smith@example.com\",\"name\":\"Sam Smith\"}}")

	req, _ := http.NewRequest("POST", url, payload)
	auth := "Bearer " + os.Getenv("SENDGRID_API_KEY")
	log.Println("auth ", auth)
	req.Header.Add("authorization", auth)
	req.Header.Add("content-type", "application/json")

	// res, _ := http.DefaultClient.Do(req)
	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// log.Println("request: ", r)
	ctx := appengine.NewContext(r)
	// ctx := r.Context()
	client := urlfetch.Client(ctx)

	resp, err := client.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// TODO the post request to sendgrid is returning a 405 response due to malformed json payload

	fmt.Fprint(w, "SendGrid response: ", resp.Status)

	fmt.Println(resp)
	fmt.Println(string(body))
	defer resp.Body.Close()
}
