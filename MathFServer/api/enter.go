package api

import (
	"MoMathF/MathFServer/service"
	"github.com/go-playground/validator/v10"
)

type APIGroup struct {
	MathAPI
	UserAPI
	AdminAPI
}

var APIGroupApp = new(APIGroup)

var validate = validator.New()

var (
	userService  = service.ServiceGroupApp.UserService
	mathService  = service.ServiceGroupApp.MathService
	adminService = service.ServiceGroupApp.AdminService
)
