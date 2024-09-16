package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(storagePath string) {
	db, err := gorm.Open(sqlite.Open(storagePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func Migrate(dst ...interface{}) {
	if err := DB.AutoMigrate(dst...); err != nil {
		panic("failed to migrate database")
	}
}
