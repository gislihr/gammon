package db

import (
	"github.com/jmoiron/sqlx"
)

func MustConnect(uri string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		panic(err)
	}

	return db
}
