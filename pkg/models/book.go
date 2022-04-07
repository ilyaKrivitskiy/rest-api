package models

type Book struct {
	Book_id      int    `json:"bookID"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Genre        string `json:"genre"`
	User_id      int    `json:"userID"`
	Author_id    int    `json:"authorID"`
	Release_date int    `json:"releaseDate"`
}
