package api

import (
	"MoMathF/MathFServer/model"
	"MoMathF/MathFServer/model/common/response"
	"MoMathF/MathFServer/model/request"
	"MoMathF/MathFServer/service"
	"MoMathF/MathFServer/utils"
	"github.com/gofiber/fiber/v2"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
)

type UserAPI struct{}

var userService = service.ServiceGroupApp.UserService

func (u *UserAPI) Login(ctx *fiber.Ctx) error {
	return nil
}

func (u *UserAPI) Register(ctx *fiber.Ctx) error {
	var r request.Register
	if err := ctx.BodyParser(&r); err != nil {
		return response.FailWithMsg("parse request error: "+err.Error(), ctx)
	}
	if err := validate.Struct(r); err != nil {
		return response.FailWithMsg("validate request error: "+err.Error(), ctx)
	}
	user := model.User{
		UUID:      fiberUtils.UUIDv4(),
		UserName:  r.UserName,
		Email:     r.Email,
		Password:  utils.MD5V(r.Password),
		Remaining: 100,
	}

	if err := userService.Register(user); err != nil {
		return response.FailWithMsg(err.Error(), ctx)
	}
	return response.Ok(ctx)
}
