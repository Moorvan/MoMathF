package service

type ServiceGroup struct {
	UserService
	MathService
}

var ServiceGroupApp = new(ServiceGroup)
