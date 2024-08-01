package models

import "errors"

type Book struct {
	UUIDHolder
	IDPK uint `gorm:"primaryKey;column:id_pk" json:"-"`
	Title string `gorm:"type:varchar(255);not null;column:title" json:"title"`
	Edition uint8 `gorm:"type:smallint;column:edition" json:"edition"`
	PublicationYear uint16 `gorm:"type:smallint;column:publication_year" json:"publication_year"`
	Authors []string `gorm:"-" json:"authors"`
}

func (book *Book) validade() error {
	if book.Title == "" {
		return errors.New("field 'title' cannot be empty")
	}
	return nil
}

func (book *Book) Prepare() error {
	if err := book.validade(); err != nil {
		return err
	}
	book.GenUUID()
	return nil
}