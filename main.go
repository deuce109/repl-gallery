package main

import (
	"fmt"
	"main/repositories"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
)

func main() {
	r := mux.NewRouter()

	session, err := mgo.Dial(os.Getenv("mongo_url"))

	if err != nil {
		panic(err)
	}

	db := session.DB("gallery")

	_ = repositories.NewPhotoRepository(db)

	if err != nil {
		fmt.Printf("%s", err)
	}

	http.ListenAndServe(":80", r)
}
