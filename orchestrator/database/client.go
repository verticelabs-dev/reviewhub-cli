package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetOrmInstance() *gorm.DB {
	if db != nil {
		return db
	}

	logger := GetLogger()
	databasePath := GetStoragePath("db/reviewhub.db")
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect with sqlite database")
	}

	logger.Info().Msg("Connected with sqlite database")

	return db
}
