package request

type Login struct {
	UserName string `json:"username" validate:"required,min=2,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
