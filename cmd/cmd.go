package cmd

import (
	"MoMathF/global"
	mlog "MoMathF/mo-log"
	"flag"
	"github.com/atotto/clipboard"
)

var (
	log = mlog.Log
)

func GetLatexFromPicAndCopyLatexToClipboard() {
	path := flag.String("path", "", "path to the pic")
	flag.Parse()
	if *path == "" {
		log.Fatalln("path can't be empty")
	}
	log.Println("path:", *path)
	res, err := global.GB_Client.GetLatexFromPic(*path)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	if err := clipboard.WriteAll(res); err != nil {
		log.Fatalln(err)
	}
}
