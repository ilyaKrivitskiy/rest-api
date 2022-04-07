package models

type User struct {
	User_id    int    `json:"userID"`
	First_name string `json:"firstName"`
	Last_name  string `json:"lastName"`
}
