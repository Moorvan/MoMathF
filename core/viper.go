package core

import (
	"MoMathF/global"
	mlog "MoMathF/mo-log"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func Viper(path string) *viper.Viper {
	flag.StringVar(&global.GB_CONFIG.ConfigPath, "config", path, "config file path")
	flag.StringVar(&global.GB_CONFIG.PicPath, "image", "", "image file path")
	flag.Parse()
	log.Println("config path:", global.GB_CONFIG.ConfigPath)
	log.Println("image path:", global.GB_CONFIG.PicPath)
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(global.GB_CONFIG.ConfigPath)
	if err := v.ReadInConfig(); err != nil {
		mlog.Log.Fatalln("Fail to read config:", err.Error())
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		mlog.Log.Println("Config file changed:", e.Name)
		if err := v.Unmarshal(&global.GB_CONFIG); err != nil {
			mlog.Log.Errorln("Fail to update config:", err.Error())
		}
	})

	if err := v.Unmarshal(&global.GB_CONFIG); err != nil {
		mlog.Log.Fatalln("Fail to get config: " + err.Error())
	}

	return v
}
