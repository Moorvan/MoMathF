package config

type Config struct {
	Api        API `json:"api"`
	ConfigPath string
	PicPath    string
	LogPath    string
	ServerCfg  Server `mapstructure:"server" json:"server"`
}
