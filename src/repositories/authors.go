package repositories

import (
	"fmt"

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

func (repository *Author) GetAuthorsByName(name string) ([]models.Authors, error) {
	var authors []models.Authors

	if err := repository.db.Find(&authors, "name LIKE ?", fmt.Sprintf("%%%s%%", name)).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (repository *Author) GetAuthorByName(name string) (models.Authors, error) {
	var author models.Authors

	if err := repository.db.Where("name = ?", name).First(&author).Error; err != nil {
		return models.Authors{}, err
	}

	if author.Name == "" {
		author.Name = name
		if err := author.Prepare(); err != nil {
			return models.Authors{}, err
		}
		if _, err := repository.InsertAuthor(author); err != nil {
			return models.Authors{}, err
		}
		return author, nil
	}

	return author, nil
}

func (repository *Author) DeleteAuthor(id uuid.UUID) error {
	if err := repository.db.Delete(&models.Authors{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
