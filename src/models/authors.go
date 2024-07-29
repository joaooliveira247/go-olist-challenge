package models

type AuthorsIn struct {
	Name string `json:"name,omitempty" gorm:"type:varchar(255);column:name"`
}

type AuthorsOut struct {
	UUIDHolder
	AuthorsIn
}

type Authors struct {
	AuthorsOut
	IDPK uint `gorm:"type:primaryKey;column:id_pk"`
}
