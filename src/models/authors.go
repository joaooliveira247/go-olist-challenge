package models

import "github.com/google/uuid"

type AuthorsIn struct {
	Name string `json:"name,omitempty" gorm:"type:varchar(255);column:name"`
}

type AuthorsOut struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;column:id"` 
	AuthorsIn
}
