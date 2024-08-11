package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"github.com/joaooliveira247/go-olist-challenge/src/utils"
	"gorm.io/gorm"
)

type Book struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *Book {
	return &Book{db}
}

func (repository *Book) verifyAuthors(
	ids []uuid.UUID,
) ([]models.Authors, error) {
	var validAuthors []models.Authors

	if err := repository.db.Where("id IN ?", ids).Find(&validAuthors).Error; err != nil {
		return nil, err
	}

	if len(validAuthors) < 1 {
		return nil, errors.New("at least one author can be valid")
	}

	return validAuthors, nil
}

func (repository *Book) InsertBook(book models.Book) (uuid.UUID, error) {
	tx := repository.db.Begin()
	result := tx.Where(
		&models.Book{
			Title:           book.Title,
			Edition:         book.Edition,
			PublicationYear: book.PublicationYear,
		},
	).FirstOrCreate(&book)
	if err := result.Error; err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}

	if result.RowsAffected < 1 {
		return uuid.UUID{}, utils.BookAlreadyExistsError
	}

	authors, err := repository.verifyAuthors(book.Authors)
	if err != nil {
		return uuid.UUID{}, err
	}

	var bookAuthors []models.BookAuthors

	for _, author := range authors {
		bookAuthors = append(
			bookAuthors,
			models.BookAuthors{BookID: book.IDPK, AuthorID: author.IDPK},
		)
	}

	if err := tx.Create(&bookAuthors).Error; err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}
	tx.Commit()

	return book.ID, nil
}

func (repository *Book) GetBooks() ([]models.Book, error) {
	var books []models.Book

	if err := repository.db.Raw(`
	SELECT 
		array_agg(a.id::text) AS authors, b.*
	FROM 
		book_authors ba
	INNER JOIN 
		authors a 
	ON 
		ba.author_id = a.id_pk 
	INNER JOIN 
		books b
	ON
		ba.book_id = b.id_pk
	group by
		b.id_pk;
	`).Scan(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
