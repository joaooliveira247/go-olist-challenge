package utils

import (
	"reflect"

	"github.com/google/uuid"
)

type BookQuery struct {
	Title           string    `form:"title"`
	PublicationYear uint16    `form:"publication_year"`
	Edition         uint8     `form:"edition"`
	Author          uuid.UUID `form:"author"`
	validParams     map[string]any
}

func (params *BookQuery) validate() {
	nullRep := map[string]any{
		"title":            "",
		"publication_year": uint16(0),
		"edition":          uint8(0),
		"author":           uuid.Nil,
	}
	validMap := make(map[string]any)
	val := reflect.Indirect(reflect.ValueOf(params))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)

		if field.PkgPath != "" {
			continue
		}

		key := field.Tag.Get("form")
		value := val.Field(i).Interface()

		if key != "" && value != nullRep[key] {
			validMap[key] = value
		}
	}

	params.validParams = validMap
}

func (params *BookQuery) IsEmpty() bool {
	params.validate()
	return len(params.validParams) < 1
}

func (params *BookQuery) AsQuery() map[string]any {
	params.validate()
	return params.validParams
}
