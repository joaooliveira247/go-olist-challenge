package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"gorm.io/gorm"
)

type Book struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *Book {
	return &Book{db}
}

func (repository *Book) verifyAuthors(ids []uuid.UUID) ([]models.Authors, error) {
	var validAuthors []models.Authors

	if err := repository.db.Where("id IN ?", ids).Find(&validAuthors).Error; err != nil {
		return nil, err
	}

	if len(validAuthors) < 1 {
		return nil, errors.New("at least one author can be valid")
	}

	return validAuthors, nil
}
