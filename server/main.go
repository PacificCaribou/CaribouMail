package main

import (
	"fmt"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
  fmt.Println("File Upload Endpoint Hit")
}

func setupRoutes() {
  http.HandleFunc("/upload", uploadFile)
  http.ListenAndServe(":8080", nil)
}

func main() {
  fmt.Println("Hello World")
}