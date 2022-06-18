package core

import (
	"MoMathF/global"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	p := global.GB_CONFIG.ServerCfg.Pgsql
	if p.Dbname == "" {
		panic(errors.New("dbname is empty"))
	}
	pgsqlConfig := postgres.Config{
		DSN: p.Dsn(), // DSN data source name
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig))
	if err != nil {
		panic(err)
	}
	return db
}
