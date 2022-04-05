package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilyaKrivitskiy/rest-api/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.CheckerFunc)

	log.Printf("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
