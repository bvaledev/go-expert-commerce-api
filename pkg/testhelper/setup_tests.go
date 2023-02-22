package testhelper

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDBTest(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(dst...)
	return db
}
