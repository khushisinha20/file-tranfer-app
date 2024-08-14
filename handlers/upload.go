package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r* http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "./templates/index.html", nil)
	case "POST":
		handleFileUpload(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func handleFileUpload(w http.ResponseWriter, r* http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Error parsing form data: " + err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file: " + err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		http.Error(w, "Error creating directory: " + err.Error(), http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join("./uploads", header.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file: " + err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file: " + err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/download?file=" + header.Filename, http.StatusSeeOther)
}