package main

import (
	"go-rest-api/internal/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/getschema", routes.GetSchema).Methods("GET")
	http.ListenAndServe(":8000", route)
}
