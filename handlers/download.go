package handlers

import "net/http"

func DownloadHandler(w http.ResponseWriter, r* http.Request) {
	fileName := r.URL.Query().Get("file")
	filePath := "/files/" + fileName

	data := struct {
		FilePath string
	}{
		FilePath: filePath,
	}

	renderTemplate(w, "./templates/download.html", data)
}