package models

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if err := os.Mkdir("data", 0755); err != nil {
			panic(err)
		}
	}

	database, err := gorm.Open(sqlite.Open("data/gomenu.db"), &gorm.Config{})

	migrate(database)

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
	fmt.Println("Database connection successfully opened")
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&Menu{},
		&Tag{},
	)
}
