package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
}
