package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadyafa/bookstore-apps/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookAppsRoutes(r)

	log.Println("Listening on port localhost:9010")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
