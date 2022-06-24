package core

import (
	"github.com/Moorvan/MoMathF/MathFClient"
	"github.com/Moorvan/MoMathF/global"
)

func NewClient() *MathFClient.MoMathFClient {
	config := &MathFClient.APIConfig{
		ApiId:  global.GB_CONFIG.Api.ApiId,
		ApiKey: global.GB_CONFIG.Api.ApiKey,
	}
	mathFClient, err := MathFClient.NewMoMathFClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	//res, err := mathFClient.GetLatexFromPicUrl("https://mathpix-ocr-examples.s3.amazonaws.com/cases_hw.jpg")
	//if err != nil || res == "" {
	//	log.Fatalln("Connect to Server Error:", err)
	//}
	//log.Println(res)
	return mathFClient
}
