package model

type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	Pwd       string
}
