package global

import (
	"MoMathF/MathFClient"
	"MoMathF/config"
	"github.com/spf13/viper"
)

var (
	GB_CONFIG config.Config
	GB_VP     *viper.Viper
	GB_Client *MathFClient.MoMathFClient
)
