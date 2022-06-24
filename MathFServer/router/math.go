package router

import (
	"github.com/Moorvan/MoMathF/MathFServer/api"
	"github.com/gofiber/fiber/v2"
)

type MathRouter struct{}

func (r *MathRouter) InitMathRouter(router fiber.Router) {
	mathRouter := router.Group("/math")
	{
		mathApi := api.APIGroupApp.MathAPI
		mathRouter.Get("/get", mathApi.GetLatexFromPic)
	}
}
