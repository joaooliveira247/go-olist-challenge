package db

import (
	"log"

	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Authors{}); err != nil {
		log.Fatal(err)
	}
}
