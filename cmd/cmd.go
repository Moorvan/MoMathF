package cmd

import (
	"github.com/Moorvan/MoMathF/global"
	mlog "github.com/Moorvan/MoMathF/mo-log"
	"github.com/atotto/clipboard"
)

var (
	log = mlog.Log
)

func GetLatexFromPicAndCopyLatexToClipboard() {
	path := global.GB_CONFIG.PicPath
	if path == "" {
		log.Fatalln("path can't be empty")
	}
	log.Println("path:", path)
	res, err := global.GB_Client.GetLatexFromPic(path)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	if err := clipboard.WriteAll(res); err != nil {
		log.Fatalln(err)
	}
}
