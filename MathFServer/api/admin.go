package api

import (
	"MoMathF/MathFServer/model"
	resp "MoMathF/MathFServer/model/common/response"
	"MoMathF/MathFServer/model/request"
	"MoMathF/MathFServer/model/response"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AdminAPI struct{}

func (api *AdminAPI) UpdateLevel(ctx *fiber.Ctx) error {
	var r request.UpdateLevel
	if err := ctx.BodyParser(&r); err != nil {
		return resp.FailWithMsg("parse request error: "+err.Error(), ctx)
	}
	if err := validate.Struct(r); err != nil {
		return resp.FailWithMsg("validate request error: "+err.Error(), ctx)
	}
	if err := api.checkRoot(r.UUID, ctx); err != nil {
		return resp.FailWithMsg(err.Error(), ctx)
	}
	u, err := adminService.UpdateLevel(r.UpdatedId)
	if err != nil {
		return resp.FailWithMsg(err.Error(), ctx)
	}
	return resp.OkWithData(response.UpdatedUser{
		UUID:     u.UUID,
		VipLevel: u.VipLevel,
	}, ctx)
}

func (api *AdminAPI) checkRoot(rootUuid string, ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uuid := claims["uuid"].(string)
	if uuid != rootUuid {
		return errors.New("user not match")
	}
	u, err := userService.GetUserByUUID(uuid)
	if err != nil {
		return err
	}
	if u.VipLevel != model.Lv2 {
		return errors.New("you are not root")
	}
	return nil
}
