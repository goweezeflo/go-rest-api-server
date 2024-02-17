package rest

import (
	"github.com/goweezeflo/go-rest-api-server/pkg/data"
	"log"
	"net/http"
)

var user = data.User{
	FirstName: "Go", LastName: "Weezeflo", Age: 18, ID: "1",
}
var (
	db = data.Db{User: user}
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(db.User.FirstName))
	if err != nil {
		log.Fatal(err)
	}
}
