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
	router.HandleFunc("/library/books/create", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/library/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/library/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/library/books/update/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/library/books/delete/{id}", handlers.DeleteBook).Methods("DELETE")

	log.Printf("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
