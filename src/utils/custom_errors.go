package utils

import (
	"fmt"
)

type BaseError struct {
	Resource string
	Message  string
}

type AlreadyExistsError struct {
	BaseError
}

type NotFoundError struct {
	BaseError
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("%s %s", e.Resource, e.Message)
}

var (
	AuthorAlreadyExistsError = &AlreadyExistsError{
		BaseError{"author", "already exists"},
	}
	BookAlreadyExistsError = &AlreadyExistsError{
		BaseError{"book", "already exists"},
	}
	BookNotFoundError = &NotFoundError{BaseError{"book", "not found"}}
)
