package main

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type post struct {
	ID      string `gorm:"primaryKey"`
	Title   string
	Content string
}

func createDb() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&post{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
