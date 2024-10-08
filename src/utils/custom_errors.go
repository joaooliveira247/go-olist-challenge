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

type NothingToUpdateError struct {
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
	BookNotFoundError   = &NotFoundError{BaseError{"book", "not found"}}
	AuthorNotFoundError = &NotFoundError{BaseError{"author", "not found"}}
	BookNothingToUpdateError = &NothingToUpdateError{BaseError{"book", "any field has the same value registred"}}
)
