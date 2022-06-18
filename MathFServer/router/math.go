package router

import (
	"MoMathF/MathFServer/api"
	"github.com/gofiber/fiber/v2"
)

type MathRouter struct{}

func (r *MathRouter) InitMathRouter(router fiber.Router) {
	mathRouter := router.Group("/math")
	{
		mathRouter.Get("/get", api.ServiceApp.GetLatexFromPic)
	}
}
