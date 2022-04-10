package postgresql

import (
	"database/sql"
	"fmt"
	"log"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "1923247", "library")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
