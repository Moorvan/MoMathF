package service

type ServiceGroup struct {
	UserService
	MathService
	AdminService
}

var ServiceGroupApp = new(ServiceGroup)
