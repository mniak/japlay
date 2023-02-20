package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type sqliteAdapter struct {
	db *sqlx.DB
}

func NewAdapter(filename string) (*sqliteAdapter, error) {
	db, err := sqlx.Connect("sqlite3", fmt.Sprintf("file:%s?cache=shared", filename))
	if err != nil {
		return nil, err
	}
	return &sqliteAdapter{
		db: db,
	}, nil
}

func (ad *sqliteAdapter) Close() error {
	if ad.db != nil {
		return ad.db.Close()
	}
	return nil
}
