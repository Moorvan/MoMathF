package main

import (
	"MoMathF/core"
	"MoMathF/global"
	mlog "MoMathF/mo-log"
)

var (
	log = mlog.Log
)

func main() {
	log.Println("Start...")
	global.GB_VP = core.Viper("./config.yaml")
	//log.PrintStruct(global.GB_CONFIG)
	global.GB_Client = core.NewClient()
	global.GB_DB = core.Gorm()
	core.RegisterTables()
	//cmd.GetLatexFromPicAndCopyLatexToClipboard()
	core.RunServer() //启动服务器
}
