package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/gorilla/mux"
	models "main/models"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	var request models.Request
	err := faker.FakeData(&request)

	if err != nil {
		fmt.Printf("%s", err)
	}

	data, err := request.ToJson()

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(data))
	http.ListenAndServe(":80", r)
}
