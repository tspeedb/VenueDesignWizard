package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

// DownloadFileHandler serves the requested XML file
func DownloadFileHandler(w http.ResponseWriter, r *http.Request, storagePath string) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	filePath := filepath.Join(storagePath, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
