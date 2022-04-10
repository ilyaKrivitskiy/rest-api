package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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
