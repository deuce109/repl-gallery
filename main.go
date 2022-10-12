package main

import (
	database "main/database"
	handlers "main/handlers"
	flags "main/utils/flags"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config := &database.DatabaseConfig{}
	err := flags.ParseArgs(config, "ConnectionString")

	_, err = database.Connect(config)
	if err != nil {
		panic(err)
	}

	handlers.FileHandler(r)

	http.ListenAndServe(":80", r)
}
