package router

import (
	"MoMathF/MathFServer/api"
	"github.com/gofiber/fiber/v2"
)

type Router struct{}

var RouterApp = new(Router)

func (r *Router) InitGetLatexFromPicRouter(router fiber.Router) {
	router.Get("/", api.ServiceApp.GetLatexFromPic)

}
