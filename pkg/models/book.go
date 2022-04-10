package models

import "database/sql"

type Book struct {
	Book_id      int           `json:"bookID"`
	Name         string        `json:"name"`
	Price        int           `json:"price"`
	Genre        string        `json:"genre"`
	User_id      sql.NullInt32 `json:"userID"`
	Author_id    int           `json:"authorID"`
	Release_date sql.NullInt32 `json:"releaseDate"`
}
