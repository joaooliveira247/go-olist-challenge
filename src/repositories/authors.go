package repositories

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"github.com/joaooliveira247/go-olist-challenge/src/utils"
	"gorm.io/gorm"
)

type Author struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *Author {
	return &Author{db}
}

func (repository *Author) InsertAuthor(
	author models.Authors,
) (uuid.UUID, error) {
	tx := repository.db.Begin()

	result := tx.Where("name = ?", author.Name).FirstOrCreate(&author)
	if err := result.Error; err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}

	if result.RowsAffected < 1 {
		tx.Rollback()
		return uuid.UUID{}, utils.AuthorAlreadyExistsError
	}

	tx.Commit()
	return author.ID, nil
}

func (repository *Author) InsertAuthors(
	authors []models.Authors,
) ([]uuid.UUID, error) {
	var uuids []uuid.UUID
	for _, author := range authors {
		uuid, err := repository.InsertAuthor(author)
		if err != nil {
			if errors.Is(err, utils.AuthorAlreadyExistsError) {
				continue
			}
			return nil, err
		}
		uuids = append(uuids, uuid)
	}
	return uuids, nil
}

func (repository *Author) GetAuthors() ([]models.Authors, error) {
	var authors []models.Authors

	if err := repository.db.Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (repository *Author) GetAuthorsByName(
	name string,
) ([]models.Authors, error) {
	var authors []models.Authors

	result := repository.db.
		Find(&authors, "name LIKE ?", fmt.Sprintf("%%%s%%", name))

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected < 1 {
		return nil, utils.AuthorNotFoundError
	}
	return authors, nil
}

func (repository *Author) GetAuthorByID(id uuid.UUID) (models.Authors, error) {
	var author models.Authors

	if err := repository.db.Where("id = ?", id).First(&author).Error; err != nil {
		return models.Authors{}, err
	}
	return author, nil
}

func (repository *Author) DeleteAuthor(id uuid.UUID) error {
	if err := repository.db.Delete(&models.Authors{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
