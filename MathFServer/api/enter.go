package api

import "github.com/gofiber/fiber/v2"

type Service struct{}

var ServiceApp = new(Service)

func (service *Service) GetLatexFromPic(ctx *fiber.Ctx) error {
	return ctx.SendString("abc")
}
