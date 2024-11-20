package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFileHandler processes file uploads and converts CSV to XML
func UploadFileHandler(w http.ResponseWriter, r *http.Request, storagePath string) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file
	csvFilePath := filepath.Join(storagePath, handler.Filename)
	out, err := os.Create(csvFilePath)
	if err != nil {
		http.Error(w, "Unable to save uploaded file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = out.ReadFrom(file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Convert CSV to XML
	xmlFilePath := filepath.Join(storagePath, handler.Filename+".xml")
	err = ConvertCSVToXML(csvFilePath, xmlFilePath)
	if err != nil {
		http.Error(w, "Error converting file to XML", http.StatusInternalServerError)
		return
	}

	// Notify the user
	fmt.Fprintf(w, "File converted successfully! You can download it <a href='/download/%s'>here</a>.", handler.Filename+".xml")
}

// ConvertCSVToXML converts a CSV file to an XML file
func ConvertCSVToXML(csvFilePath, xmlFilePath string) error {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	type Row struct {
		Columns []string `xml:"Column"`
	}
	type Table struct {
		Rows []Row `xml:"Row"`
	}

	var table Table
	for _, record := range records {
		row := Row{Columns: record}
		table.Rows = append(table.Rows, row)
	}

	xmlData, err := xml.MarshalIndent(table, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(xmlFilePath, xmlData, 0644)
}
