package handlers

import (
	"encoding/json"
	"fmt"
	// repositories "main/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func getFiles(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  var page, limit int
  var err error
  if pageTemp, ok := params["page"]; ok {
    page, err = strconv.Atoi(pageTemp[0])
    if (err != nil) {
      fmt.Printf("Error occoured in page number parsing:\n%t", err)
    }
  } else {
    page = 1
    
  }
  if limitTemp, ok := params["limit"]; ok {
    limit, err = strconv.Atoi(limitTemp[0])
    if (err != nil) {
      fmt.Printf("Error occoured in page number parsing:\n%t", err)
    }
  } else {
    limit = 10
  }
  fmt.Printf("Page: %d, Limit: %d", page, limit)
  
  // cursor := (page - 1) * limit

  data, err := os.ReadFile("./main")

  if err != nil {
    fmt.Printf("Error occoured in file reading:\n%t", err)
  }  

  output, err := json.Marshal(data)
  if (err != nil) {
      fmt.Printf("Error occoured in json parsing:\n%t", err)
  }
  w.Write(output)
}


func FileHandler(r *mux.Router) {
  s := r.PathPrefix("/files").Subrouter()
  s.HandleFunc("",getFiles )
}