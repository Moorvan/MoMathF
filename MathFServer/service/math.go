package service

import "MoMathF/global"

type MathService struct{}

func (service *MathService) GetLatexFromPic(path string) (string, error) {
	return global.GB_Client.GetLatexFromPic(path)
}
