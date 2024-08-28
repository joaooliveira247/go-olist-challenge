package models

import (
	"errors"

	"github.com/joaooliveira247/go-olist-challenge/src/utils"
)

type Authors struct {
	Name string `json:"name,omitempty" gorm:"type:varchar(255);column:name;unique;not null"`
	UUIDHolder
	IDPK uint `json:"-"              gorm:"primaryKey;column:id_pk"`
}

func (author *Authors) validade() error {
	if author.Name == "" {
		return errors.New("field cannot be empty")
	}
	return nil
}

func (author *Authors) Prepare() error {
	if err := author.validade(); err != nil {
		return err
	}
	author.GenUUID()
	return nil
}

func (author *Authors) ParseValidate(csv utils.AuthorCSV) error {
	author.Name = csv.Name
	if err := author.Prepare(); err != nil {
		return err
	}

	return nil
}
