package global

import (
	"github.com/Moorvan/MoMathF/MathFClient"
	"github.com/Moorvan/MoMathF/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GB_CONFIG config.Config
	GB_VP     *viper.Viper
	GB_Client *MathFClient.MoMathFClient
	GB_DB     *gorm.DB
)
