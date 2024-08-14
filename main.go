package main

import (
	"file-transfer-app/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/files/", handlers.FileHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)

	fmt.Println("Serving at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}