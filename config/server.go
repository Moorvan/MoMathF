package config

type Server struct {
	Port      int    `mapstructure:"port" json:"port,string"`
	CachePath string `mapstructure:"cache-path" json:"cache-path"`
	Pgsql     Pgsql  `mapstructure:"pgsql" json:"pgsql" mapstructure:"pgsql"`
	JWT       JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
