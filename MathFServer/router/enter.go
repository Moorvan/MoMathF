package router

type RouterGroup struct {
	MathRouter
	UserRouter
	AdminRouter
}

var RouterGroupApp = new(RouterGroup)
