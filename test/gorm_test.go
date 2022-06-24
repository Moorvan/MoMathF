package test

import (
	"github.com/Moorvan/MoMathF/MathFServer/model"
	"github.com/Moorvan/MoMathF/core"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

func init() {
	db = core.Gorm()
}

func TestDBConnectAndCreateTable(t *testing.T) {
	if db == nil {
		t.Error("db is nil")
	}
	err := db.AutoMigrate(model.User{})
	if err != nil {
		t.Error(err)
	}
}
