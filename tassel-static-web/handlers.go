package main

// handler functions
/*
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

// boilerplate
// using json.NewDecoder & strings.NewReader
// https://attilaolah.eu/2013/11/29/json-decoding-in-go/
func jsonStreamHandler(w http.ResponseWriter, r *http.Request) {
	var msg interface{}

	const jsonStream = `{"Message": "Hello", "Array": [1, 2, 3]}`
	dec := json.NewDecoder(strings.NewReader(jsonStream)) // TODO do not use this on large datasets
	// log.Println("json example", dec)

	if err := dec.Decode(&msg); err != nil {
		return
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}
	w.Write(b)

}

// echoHandler reads a JSON object from the body, and writes it back out.
// https://gist.github.com/campoy/7b44f6ec2d9e82d956d34b4989b33192
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var msg interface{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			log.Fatal(w, http.StatusBadRequest, "Body was not valid JSON: %v", err)
			return
		}
		log.Fatal(w, http.StatusInternalServerError, "Could not get body: %v", err)
		return
	}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(w, http.StatusInternalServerError, "Could not marshal JSON: %v", err)
		return
	}
	log.Println("json response: ", string(b))
	w.Write(b)
}

// errorf writes a swagger-compliant error response.
func errorHandler(w http.ResponseWriter, code int, format string, a ...interface{}) {
	var out struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	out.Code = code
	out.Message = fmt.Sprintf(format, a...)

	b, err := json.Marshal(out)
	if err != nil {
		http.Error(w, `{"code": 500, "message": "Could not format JSON for original message."}`, 500)
		return
	}

	http.Error(w, string(b), code)
}

*/
