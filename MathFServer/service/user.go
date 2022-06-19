package service

import (
	"MoMathF/MathFServer/model"
	"MoMathF/global"
	"errors"
)

type UserService struct{}

func (service *UserService) Register(user model.User) error {
	db := global.GB_DB
	var c int64
	if db.Model(&user).Where("email = ?", user.Email).Count(&c); c > 0 {
		return errors.New("email already exists")
	}
	if err := db.Create(&user).Error; err != nil {
		return errors.New("create user error: " + err.Error())
	}
	return nil
}

func (service *UserService) Login() {

}
