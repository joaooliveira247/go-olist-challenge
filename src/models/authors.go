package models

import "github.com/google/uuid"

type AuthorsIn struct {
	Name string `json:"name,omitempty"`
}

type AuthorsOut struct {
	ID uuid.UUID `json:"id,omitempty"`
	AuthorsIn
}

