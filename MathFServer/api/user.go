package api

import (
	"MoMathF/MathFServer/model"
	resp "MoMathF/MathFServer/model/common/response"
	"MoMathF/MathFServer/model/request"
	"MoMathF/MathFServer/model/response"
	"MoMathF/MathFServer/utils"
	"MoMathF/global"
	"github.com/gofiber/fiber/v2"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserAPI struct{}

func (u *UserAPI) Login(ctx *fiber.Ctx) error {
	var r request.Login
	if err := ctx.BodyParser(&r); err != nil {
		return resp.FailWithMsg("parse request error: "+err.Error(), ctx)
	}
	if err := validate.Struct(r); err != nil {
		return resp.FailWithMsg("validate request error: "+err.Error(), ctx)
	}
	user := model.User{
		Email:    r.Email,
		Password: utils.MD5V(r.Password),
	}
	if rUser, err := userService.Login(user); err != nil {
		return resp.FailWithMsg(err.Error(), ctx)
	} else {
		token, err := u.getToken(rUser)
		if err != nil {
			return resp.FailWithMsg(err.Error(), ctx)
		}
		return resp.OkWithData(response.Login{
			Token:     token,
			UUID:      rUser.UUID,
			UserName:  rUser.UserName,
			Email:     rUser.Email,
			Remaining: rUser.Remaining,
			VipLevel:  rUser.VipLevel,
		}, ctx)
	}
}

func (u *UserAPI) getToken(user model.User) (string, error) {
	cfg := global.GB_CONFIG.ServerCfg.JWT
	var claims = jwt.MapClaims{
		"uuid": user.UUID,
		"name": user.UserName,
		"exp":  time.Now().Add(time.Second * time.Duration(cfg.ExpiresTime)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(cfg.SigningKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (u *UserAPI) Register(ctx *fiber.Ctx) error {
	var r request.Register
	if err := ctx.BodyParser(&r); err != nil {
		return resp.FailWithMsg("parse request error: "+err.Error(), ctx)
	}
	if err := validate.Struct(r); err != nil {
		return resp.FailWithMsg("validate request error: "+err.Error(), ctx)
	}
	user := model.User{
		UUID:      fiberUtils.UUIDv4(),
		UserName:  r.UserName,
		Email:     r.Email,
		Password:  utils.MD5V(r.Password),
		Remaining: 10,
		VipLevel:  model.Lv0,
	}

	if err := userService.Register(user); err != nil {
		return resp.FailWithMsg(err.Error(), ctx)
	}
	return resp.Ok(ctx)
}
