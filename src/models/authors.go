package models

import "errors"

type Authors struct {
	Name string `json:"name,omitempty" gorm:"type:varchar(255);column:name"`
	UUIDHolder
	IDPK uint `json:"-" gorm:"primaryKey;column:id_pk"`
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
