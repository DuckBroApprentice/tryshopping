package models

type User struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Users struct {
	UserList     []User
	UserListSize int
}
