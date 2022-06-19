package router

type RouterGroup struct {
	MathRouter
	UserRouter
}

var RouterGroupApp = new(RouterGroup)
