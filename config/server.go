package config

type Server struct {
	Port      int    `mapstructure:"port" json:"port,string"`
	CachePath string `mapstructure:"cache-path" json:"cache-path"`
}
