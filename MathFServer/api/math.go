package api

import (
	"MoMathF/MathFServer/model/common/response"
	"MoMathF/MathFServer/service"
	"MoMathF/global"
	"github.com/gofiber/fiber/v2"
	"os"
)

type MathAPI struct{}

var mathService = service.ServiceGroupApp.MathService

func (api *MathAPI) GetLatexFromPic(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		response.FailWithMsg("can't get multipartForm", ctx)
		return nil
	}
	files := form.File["img"]
	if len(files) == 0 {
		response.FailWithMsg("can't get file", ctx)
		return nil
	}

	file := files[0]
	path := global.GB_CONFIG.ServerCfg.CachePath
	if path[len(path)-1] != '/' {
		path += "/"
	}
	path += file.Filename

	if err := ctx.SaveFile(file, path); err != nil {
		return err
	}

	latex, err := mathService.GetLatexFromPic(path)
	if err != nil {
		return err
	}
	_ = os.Remove(path)

	response.OkWithData(map[string]string{"latex": latex}, ctx)

	return nil
}
