package response

import "MoMathF/MathFServer/model"

type UpdatedUser struct {
	UUID     string         `json:"uuid"`
	VipLevel model.VipLevel `json:"vip_level"`
}
