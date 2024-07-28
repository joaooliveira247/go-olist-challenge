package db

import (
	"github.com/joaooliveira247/go-olist-challenge/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
