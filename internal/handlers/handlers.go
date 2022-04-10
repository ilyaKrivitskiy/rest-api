package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilyaKrivitskiy/rest-api/pkg/models"
	"github.com/ilyaKrivitskiy/rest-api/pkg/postgresql"
	_ "github.com/lib/pq"
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

	if req.Method != http.MethodGet {
		w.Header().Add("Allow", "GET")
		http.Error(w, "This method is not allowed!", http.StatusMethodNotAllowed)
		return
	}

	db := postgresql.SetupDB()
	log.Println("Db is working in GetAllBooks!")
	defer db.Close()

	books := []models.Book{}

	rows, err := db.Query("select * from book order by book_id")
	if err != nil {
		log.Fatalln(err.Error())
	}

	for rows.Next() {
		var book_id int
		var name string
		var price int
		var genre string
		var user_id sql.NullInt32
		var author_id int
		var release_date sql.NullInt32

		err = rows.Scan(&book_id, &name, &price, &genre, &user_id, &author_id, &release_date)
		if err != nil {
			log.Fatalln(err.Error())
		}

		books = append(books, models.Book{Book_id: book_id, Name: name, Price: price, Genre: genre, User_id: user_id,
			Author_id: author_id, Release_date: release_date})
	}

	json.NewEncoder(w).Encode(&books)
}

func GetBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodGet {
		w.Header().Add("Allow", "GET")
		http.Error(w, "This method is not allowed!", http.StatusMethodNotAllowed)
		return
	}

	book_id := mux.Vars(req)["id"]

	if book_id == "" {
		http.Error(w, "Wrong id!", 404)
		return
	}
	res_id, err := strconv.Atoi(book_id)

	db := postgresql.SetupDB()
	log.Println("Db is working in getItem by id!")
	defer db.Close()

	var max_id int
	max_id_row := db.QueryRow("select max(book_id) from book")
	check_id_error := max_id_row.Scan(&max_id)

	if err != nil || res_id < 1 || check_id_error != nil || max_id < res_id {
		http.Error(w, "There is no such id!", 404)
		return
	}

	book := models.Book{}

	row := db.QueryRow("select * from book where book_id = $1", res_id)

	err = row.Scan(&book.Book_id, &book.Name, &book.Price, &book.Genre, &book.User_id, &book.Author_id, &book.Release_date)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "There is no any item with this id!", 404)
		return
	}

	json.NewEncoder(w).Encode(&book)
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("UpdateBook function..."))
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("DeleteBook function..."))
}
