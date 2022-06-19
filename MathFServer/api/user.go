package api

import (
	"MoMathF/MathFServer/model/common/response"
	"MoMathF/MathFServer/model/request"
	"github.com/gofiber/fiber/v2"
)

type UserAPI struct{}

func (u *UserAPI) Login(ctx *fiber.Ctx) error {
	return nil
}

func (u *UserAPI) Register(ctx *fiber.Ctx) error {
	var r request.Register
	if err := ctx.BodyParser(&r); err != nil {
		response.FailWithMsg("parse request error: "+err.Error(), ctx)
		return nil
	}
	if err := validate.Struct(r); err != nil {
		response.FailWithMsg("validate request error: "+err.Error(), ctx)
		return nil
	}

	response.Ok(ctx)
	return nil
}
