package cmd

import (
	"MoMathF/global"
	mlog "MoMathF/mo-log"
	"flag"
)

var (
	log    = mlog.Log
	client = global.GB_Client
)

func GetLatexFromPic() {
	path := flag.String("path", "", "path to the pic")
	flag.Parse()
	if *path == "" {
		log.Fatalln("path can't be empty")
	}
	log.Println("path:", *path)
}
