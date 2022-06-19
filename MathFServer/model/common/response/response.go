package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data any, msg string, c *fiber.Ctx) {
	if err := c.JSON(Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}); err != nil {
		panic(err)
	}
}

func Ok(c *fiber.Ctx) {
	Result(SUCCESS, map[string]any{}, "success", c)
}

func OkWithData(data any, c *fiber.Ctx) {
	Result(SUCCESS, data, "success", c)
}

func FailWithMsg(msg string, c *fiber.Ctx) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithDetailed(data any, msg string, c *fiber.Ctx) {
	Result(ERROR, data, msg, c)
}
