package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func createDirIfDNE(filename string) err {
  // ensure file exists
  fileInfo, err := os.Stat(filename)
  if os.IsNotExist(err) {
    // make dir if file DNE
    err := os.Mkdir(filename, 0755)
    if (err != nil) {
      fmt.Println("Error: while making temp dir: ", err)
      return
    }
  } else if (err != nil) {
    fmt.Println("Error while opening temp dir: ", err)
    return
  } else if (! fileInfo.IsDir()) {
    fmt.Printf("File %s is not a directory\n", fileInfo.Name())
    return
  }
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  fmt.Println("Uploading file...")

  r.ParseMultipartForm(10 << 20)

  file, _, err := r.FormFile("myFile")
  if err != nil {
    fmt.Printf("While reading user file - %s", err)
    return
  }
  defer file.Close()

  // ensure uploads dir exists
  fileInfo, err := os.Stat("uploads")
  if os.IsNotExist(err) {
    // make dir if file DNE
    err := os.Mkdir("uploads", 0755)
    if (err != nil) {
      fmt.Println("Error: while making temp dir: ", err)
      return
    }
  } else if (err != nil) {
    fmt.Println("Error while opening temp dir: ", err)
    return
  } else if (! fileInfo.IsDir()) {
    fmt.Printf("File %s is not a directory\n", fileInfo.Name())
    return
  }

  // read user file and write to disk
  tmpFile, err := ioutil.TempFile("uploads", "upload-*")
  if err != nil {
    fmt.Println("Error while opening temp file: ", err)
    return
  }
  defer tmpFile.Close()

  fileBytes, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println("Error while reading user-uploaded file: ", err)
    fmt.Println(err)
  }

  tmpFile.Write(fileBytes)
  fmt.Println(w, "File uploaded.")
}

func setupRoutes() {
  http.HandleFunc("/upload", uploadHandler)
  http.HandleFunc("/download", downloadHandler)
  http.ListenAndServe(":8080", nil)
}

func main() {
  setupRoutes()
}
