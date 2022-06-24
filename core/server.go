package core

import (
	"fmt"
	"github.com/Moorvan/MoMathF/MathFServer"
	"github.com/Moorvan/MoMathF/global"
)

func RunServer() {
	server := MathFServer.New()
	addr := fmt.Sprintf(":%d", global.GB_CONFIG.ServerCfg.Port)
	log.Fatalln(server.Listen(addr))
}
