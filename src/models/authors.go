package models

type Authors struct {
	Name string `json:"name,omitempty" gorm:"type:varchar(255);column:name"`
	UUIDHolder
	IDPK uint `json:"-" gorm:"type:primaryKey;column:id_pk"`
}
