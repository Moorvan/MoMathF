package model

type User struct {
	UUID     string `json:"uuid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
