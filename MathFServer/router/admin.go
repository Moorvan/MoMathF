package router

import (
	"github.com/Moorvan/MoMathF/MathFServer/api"
	"github.com/gofiber/fiber/v2"
)

type AdminRouter struct{}

func (r *AdminRouter) InitAdminRouter(router fiber.Router) {
	adminRouter := router.Group("/admin")
	{
		adminApi := api.APIGroupApp.AdminAPI
		adminRouter.Post("/updateLevel", adminApi.UpdateLevel)
	}
}
