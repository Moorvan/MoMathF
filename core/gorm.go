package core

import (
	"errors"
	"github.com/Moorvan/MoMathF/MathFServer/model"
	"github.com/Moorvan/MoMathF/global"
	mlog "github.com/Moorvan/MoMathF/mo-log"
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

func RegisterTables() {
	if err := global.GB_DB.AutoMigrate(
		model.User{},
	); err != nil {
		panic(err)
	}
	mlog.Log.Println("Register tables success")
}
