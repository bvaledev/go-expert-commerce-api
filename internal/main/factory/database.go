package factory

import (
	"github.com/bvaledev/go-expert-commerce-api/internal/domain/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DatabaseConection *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		panic(err)
	}

	DatabaseConection = db
}
