package utils

import "github.com/google/uuid"


type BookQuery struct {
	Title string `form:"title"`
	PublicationYear uint16 `form:"publication_year"`
	Edition uint8 `form:"edition"`
	Author uuid.UUID `form:"author"`
}