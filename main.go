package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Ensure the storage directory exists
	storagePath := "./storage"
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		err := os.Mkdir(storagePath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create storage directory: %v", err)
		}
	}

	// Initialize Gorilla Mux router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		UploadFileHandler(w, r, storagePath)
	}).Methods("POST")
	r.HandleFunc("/download/{fileName}", func(w http.ResponseWriter, r *http.Request) {
		DownloadFileHandler(w, r, storagePath)
	}).Methods("GET")

	// Start the server
	port := "8080"
	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
