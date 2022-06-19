package model

type User struct {
	UUID      string `json:"uuid" gorm:"primary_key"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Remaining int    `json:"remaining"`
}
