package service

import (
	"MoMathF/MathFServer/model"
	"MoMathF/global"
	"errors"
)

type AdminService struct{}

func (service *AdminService) UpdateLevel(uuid string) (model.User, error) {
	db := global.GB_DB
	u := model.User{UUID: uuid}
	if err := db.Model(&u).First(&u).Error; err != nil {
		return u, err
	}
	if u.VipLevel == model.Lv0 {
		if err := db.Model(&u).Update("vip_level", model.Lv1).Error; err != nil {
			return u, err
		}
		u.VipLevel = model.Lv1
		return u, nil
	} else if u.VipLevel == model.Lv1 {
		if err := db.Model(&u).Update("vip_level", model.Lv0).Error; err != nil {
			return u, err
		}
		u.VipLevel = model.Lv0
		return u, nil
	}
	return u, errors.New("root user level can't be updated")
}
