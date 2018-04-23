
package main

import (
        "fmt"
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
