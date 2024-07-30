package repositories

import (
	"github.com/google/uuid"
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"gorm.io/gorm"
)

type Author struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *Author {
	return &Author{db}
}

func (repository *Author) InsertAuthor(author models.Authors) (uuid.UUID, error) {
	tx := repository.db.Begin()
	if err := tx.Create(&author).Error; err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}
	tx.Commit()
	return author.ID, nil
}

func (repository *Author) GetAuthors() ([]models.Authors, error) {
	var authors []models.Authors

	if err := repository.db.Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (repository *Author) DeleteAuthor(id uuid.UUID) error {
	if err := repository.db.Delete(&models.Authors{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}