package main

import (
	"MoMathF/cmd"
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
	global.GB_Client = core.NewMoMathF()
	cmd.GetLatexFromPicAndCopyLatexToClipboard()
}
