package models

type Author struct {
	User_id    int    `json:"authorID"`
	First_name string `json:"firstName"`
	Last_name  string `json:"lastName"`
}
