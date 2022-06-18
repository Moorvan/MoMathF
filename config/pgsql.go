package config

type Pgsql struct {
	Path     string `json:"path" yaml:"path"`
	Port     string `json:"port" yaml:"port"`
	Dbname   string `json:"db-name" yaml:"db-name" mapstructure:"db-name"`
	UserName string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + " user=" + p.UserName + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port
}
