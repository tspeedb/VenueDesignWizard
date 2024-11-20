package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
)

// HomePage serves the index.html page
func HomePage(w http.ResponseWriter, r *http.Request) {
	// Debugging log: print the absolute file path to ensure it's correct
	absolutePath, err := filepath.Abs("templates/index.html")
	if err != nil {
		http.Error(w, "Unable to resolve path", http.StatusInternalServerError)
		return
	}
	fmt.Println("Serving file from path:", absolutePath)
	
	http.ServeFile(w, r, "templates/index.html")
}
