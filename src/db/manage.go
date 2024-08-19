package db

import (
	"github.com/joaooliveira247/go-olist-challenge/src/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Authors{}, &models.Book{}, &models.BookAuthors{}); err != nil {
		return err
	}
	return nil
}

func Delete(db *gorm.DB) {
	db.Exec(
		`
do $$ declare
    r record;
begin
    for r in (select tablename from pg_tables where schemaname = 'public') loop
        execute 'drop table if exists ' || quote_ident(r.tablename) || ' cascade';
    end loop;
end $$;
		`,
	)
}
