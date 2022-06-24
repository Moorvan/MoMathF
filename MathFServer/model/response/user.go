package response

import "github.com/Moorvan/MoMathF/MathFServer/model"

type Login struct {
	Token     string         `json:"token"`
	UUID      string         `json:"uuid"`
	UserName  string         `json:"username"`
	Email     string         `json:"email"`
	Remaining int            `json:"remaining"`
	VipLevel  model.VipLevel `json:"vip_level"`
}
