package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"
)

var Connection *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("data/db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Connection = db
}
