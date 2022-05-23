package core

import (
	"MoMathF/MathFServer"
	"MoMathF/global"
	"fmt"
)

func RunServer() {
	server := MathFServer.New()
	addr := fmt.Sprintf(":%d", global.GB_CONFIG.ServerCfg.Port)
	log.Fatalln(server.Listen(addr))
}
