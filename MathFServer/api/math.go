package api

import (
	"MoMathF/MathFServer/model/common/response"
	"MoMathF/MathFServer/service"
	"MoMathF/global"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type MathAPI struct{}

var mathService = service.ServiceGroupApp.MathService

func (api *MathAPI) GetLatexFromPic(ctx *fiber.Ctx) error {
	if err := api.checkUser(ctx); err != nil {
		return response.FailWithMsg(err.Error(), ctx)
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

	return response.OkWithData(map[string]string{"latex": latex}, ctx)
}

func (api *MathAPI) checkUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uuid := claims["uuid"].(string)
	if uuid != ctx.FormValue("uuid") {
		return errors.New("user not match")
	}
	return nil
}
