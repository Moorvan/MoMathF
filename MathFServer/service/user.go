package service

import (
	"errors"
	"github.com/Moorvan/MoMathF/MathFServer/model"
	"github.com/Moorvan/MoMathF/global"
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

func (service *UserService) Login(user model.User) (model.User, error) {
	db := global.GB_DB
	email := user.Email
	var u model.User
	if err := db.Model(&u).Where("email = ?", email).First(&u).Error; err != nil {
		return u, errors.New("email not exists")
	}
	if u.Password != user.Password {
		return u, errors.New("password error")
	}
	return u, nil
}

func (service *UserService) GetUserByUUID(uuid string) (model.User, error) {
	db := global.GB_DB
	var u model.User
	if err := db.Model(&u).Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return u, errors.New("user not exists")
	}
	return u, nil
}
