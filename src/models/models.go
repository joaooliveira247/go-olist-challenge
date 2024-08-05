package models

import "github.com/google/uuid"

type UUIDHolder struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"unique;not null;type:uuid;column:id"`
}

type BookAuthors struct {
	BookID   uint    `gorm:"primaryKey;column:book_id"`
	Book     Book    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:BookID;"`
	AuthorID uint    `gorm:"primaryKey;column:author_id"`
	Author   Authors `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AuthorID;"`
}

func (holder *UUIDHolder) GenUUID() {
	holder.ID = uuid.New()
}
