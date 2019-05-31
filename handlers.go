package main

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"time"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type optin_email struct {
	Email string `json:"email"`
	//timestamp time.Time `json:"timestamp"`
	Timestamp int64 `json:"timestamp"`
}
// timestamp = time.Now()
// timestamp.Format(time.RFC3339)

// echoHandler reads a JSON object from the body, and writes it back out.
// https://gist.github.com/campoy/7b44f6ec2d9e82d956d34b4989b33192
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var msg interface{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			log.Fatal(w, http.StatusBadRequest, "Body was not valid JSON: %v", err)
			w.WriteHeader(404)
			return
		}
		log.Fatal(w, http.StatusInternalServerError, "Could not get body: %v", err)
		w.WriteHeader(404)
		return
	}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(w, http.StatusInternalServerError, "Could not marshal JSON: %v", err)
		w.WriteHeader(404)
		return
	}
	log.Println("json response: ", string(b))
	w.Write(b)
}

// receive email string
func EmailHandler(w http.ResponseWriter, req *http.Request) {

	var msg optin_email

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	// Unmarshal into struct
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	msg.Timestamp = time.Now().Unix() 
	
	// marshal for verification on client side
	j, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(w, http.StatusInternalServerError, "Could not marshal JSON: %v", err)
		w.WriteHeader(404)
		return
	}
	log.Println("json input: ", string(b))
	log.Println("msg marshalled: ", msg)
	log.Println("json output: ", j)

	w.WriteHeader(200)
	w.Write(j)
}