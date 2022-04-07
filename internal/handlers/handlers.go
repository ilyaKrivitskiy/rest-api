package handlers

import (
	"net/http"
	//"github.com/ilyaKrivitskiy/rest-api/pkg"
)

func CheckerFunc(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(w, "Wrong url adress!", 404)
		return
	}

	w.Write([]byte("Homepage."))
}

func CreateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("CreateBook function..."))
}

func GetAllBooks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("GetAllBooks function..."))
}

func GetBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("GetBook function..."))
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("UpdateBook function..."))
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("DeleteBook function..."))
}
