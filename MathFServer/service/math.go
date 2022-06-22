package service

import (
	"MoMathF/MathFServer/model"
	"MoMathF/global"
)

type MathService struct{}

func (service *MathService) GetLatexFromPic(path string) (string, error) {
	return global.GB_Client.GetLatexFromPic(path)
}

func (service *MathService) Consume1(uuid string) error {
	db := global.GB_DB
	var remaining int
	if err := db.Model(&model.User{UUID: uuid}).Select("remaining").First(&remaining).Error; err != nil {
		return err
	}
	if err := db.Model(&model.User{UUID: uuid}).Update("remaining", remaining-1).Error; err != nil {
		return err
	}
	return nil
}

func (service MathService) QueryRemaining(uuid string) (int, error) {
	db := global.GB_DB
	var remaining int
	if err := db.Model(&model.User{UUID: uuid}).Select("remaining").First(&remaining).Error; err != nil {
		return 0, err
	}
	return remaining, nil
}
