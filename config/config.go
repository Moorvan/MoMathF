package config

type Config struct {
	Api        API `json:"api"`
	ConfigPath string
	PicPath    string
	LogPath    string
	ServerPort int `mapstructure:"server-port" json:"server-port"`
}
