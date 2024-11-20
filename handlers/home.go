package main

import (
	"net/http"
)

// HomePage serves the index.html page
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
