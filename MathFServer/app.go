package MathFServer

import (
	"MoMathF/MathFServer/router"
	"MoMathF/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func New() *fiber.App {
	cfg := global.GB_CONFIG.ServerCfg
	app := fiber.New()
	app.Use(logger.New())
	publicGroup := app.Group("")
	{
		router.RouterGroupApp.InitUserRouter(publicGroup)
	}
	privateGroup := app.Group("")
	privateGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(cfg.JWT.SigningKey),
	}))
	{
		router.RouterGroupApp.InitMathRouter(privateGroup)
	}
	return app
}
