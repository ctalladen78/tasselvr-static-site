package main

// handler functions

import (
	"fmt"
	"net/http"
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
