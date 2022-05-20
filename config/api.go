package config

type API struct {
	ApiId  string `mapstructure:"api-id" yaml:"api-id"`
	ApiKey string `mapstructure:"api-key" yaml:"api-key"`
}
