package database

import (
	"sang-kaghaz-gheichi/game"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("games.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&game.GameRecord{})
	return db
}
