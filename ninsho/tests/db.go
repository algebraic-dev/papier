package tests

import (
	"ninsho/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	testDB = db
	db.AutoMigrate(&models.User{})
}
