package router

//type Router struct{}
//
//var RouterApp = new(Router)
//
//func (r *Router) InitGetLatexFromPicRouter(router fiber.Router) {
//	router.Get("/", api.ServiceApp.GetLatexFromPic)
//}

type RouterGroup struct {
	MathRouter
	UserRouter
}

var RouterGroupApp = new(RouterGroup)
