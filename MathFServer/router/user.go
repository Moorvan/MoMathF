package router

import (
	"github.com/Moorvan/MoMathF/MathFServer/api"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(router fiber.Router) {
	userRouter := router.Group("/user")
	{
		userApi := api.APIGroupApp.UserAPI
		userRouter.Post("/register", userApi.Register)
		userRouter.Post("/login", userApi.Login)
	}
}
