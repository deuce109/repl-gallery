package main

import (
	"github.com/gorilla/mux"
	handlers "main/handlers"
  database "main/database"
	"net/http"
  flags "main/utils/flags"
)

func main() {
	r := mux.NewRouter()
  config := &database.DatabaseConfig{}
  flags.ParseArgs(config, "ConnectionString")
  if config.ConnectionString == "" {
    config.ConnectionString = ":memory:"
  }

  _, err := database.Connect(config)
  if err != nil {
    panic(err)
  }

  
	handlers.FileHandler(r)
  
	http.ListenAndServe(":80", r)
}