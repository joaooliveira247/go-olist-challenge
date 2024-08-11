package utils

import (
	"fmt"
)

type AlreadyExistsError struct {
	Resource string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%s already exists", e.Resource)
}

var (
	AuthorAlreadyExistsError = &AlreadyExistsError{"author"}
	BookAlreadyExistsError   = &AlreadyExistsError{"book"}
)
