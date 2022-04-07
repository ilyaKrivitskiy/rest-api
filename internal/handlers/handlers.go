package handlers

import "net/http"

func CheckerFunc(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(w, "Wrong url adress!", 404)
		return
	}

	w.Write([]byte("Homepage."))
}

func CreateBook(w http.ResponseWriter, req *http.Request) {

}

func GetAllBooks(w http.ResponseWriter, req *http.Request) {

}

func GetBook(w http.ResponseWriter, req *http.Request) {

}

func UpdateBook(w http.ResponseWriter, req *http.Request) {

}

func DeleteBook(w http.ResponseWriter, req *http.Request) {

}
