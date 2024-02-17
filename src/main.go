package main

import (
	"github.com/goweezeflo/go-rest-api-server/pkg/rest"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", rest.UsersHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
