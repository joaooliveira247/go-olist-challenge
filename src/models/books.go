package models

type Book struct {
	UUIDHolder
	IDPK uint `gorm:"primaryKey;column:id_pk" json:"-"`
	Title string `gorm:"type:varchar(255);not null;column:title"`
	Edition uint8 `gorm:"type:smallint;column:edition"`
	PublicationYear uint16 `gorm:"type:smallint;column:publication_year"`
	Authors []string `gorm:"-" json:"authors"`
}