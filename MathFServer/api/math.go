package api

import (
	"MoMathF/MathFServer/model"
	"MoMathF/MathFServer/model/common/response"
	"MoMathF/global"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type MathAPI struct{}

func (api *MathAPI) GetLatexFromPic(ctx *fiber.Ctx) error {
	uuid, err := api.checkUser(ctx)
	if err != nil {
		return response.FailWithMsg(err.Error(), ctx)
	}
	user, err := mathService.QueryUserInfo(uuid)
	if err != nil {
		return response.FailWithMsg("internal error", ctx)
	}
	if user.VipLevel != model.Lv2 && user.Remaining <= 0 {
		return response.FailWithMsg("remaining is 0", ctx)
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return response.FailWithMsg("can't get multipartForm", ctx)
	}
	files := form.File["img"]
	if len(files) == 0 {
		return response.FailWithMsg("can't get file", ctx)
	}

	file := files[0]
	path := global.GB_CONFIG.ServerCfg.CachePath
	if path[len(path)-1] != '/' {
		path += "/"
	}
	path += file.Filename

	if err := ctx.SaveFile(file, path); err != nil {
		return response.FailWithMsg("can't save file", ctx)
	}

	latex, err := mathService.GetLatexFromPic(path)
	if err != nil {
		return response.FailWithMsg(err.Error(), ctx)
	}
	_ = os.Remove(path)
	if user.VipLevel != model.Lv2 {
		if err := mathService.Consume1(uuid); err != nil {
			return response.FailWithMsg("internal error", ctx)
		}
	}
	return response.OkWithData(map[string]string{"latex": latex}, ctx)
}

func (api *MathAPI) checkUser(ctx *fiber.Ctx) (string, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uuid := claims["uuid"].(string)
	if uuid != ctx.FormValue("uuid") {
		return "", errors.New("user not match")
	}
	return uuid, nil
}
