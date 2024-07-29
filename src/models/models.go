package models

import "github.com/google/uuid"

type UUIDHolder struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"unique;not null;type:uuid;column:id"`
}

func (holder *UUIDHolder) GenUUID() {
	holder.ID = uuid.New()
}
