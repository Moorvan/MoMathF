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
}
