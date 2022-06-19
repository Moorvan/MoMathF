package MathFServer

import (
	"MoMathF/MathFServer/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	publicGroup := app.Group("")
	{
		router.RouterGroupApp.InitMathRouter(publicGroup)
	}
	return app
}
