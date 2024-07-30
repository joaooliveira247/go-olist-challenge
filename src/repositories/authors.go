package repositories

import "gorm.io/gorm"

type Author struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *Author {
	return &Author{db}
}
