package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	UseSSL   bool
}

func (o Options) toConnectionString() string {
	sslMode := "disable"
	if o.UseSSL {
		sslMode = "enable"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=%s", o.Host, o.Port, o.User, o.Password, o.Database, sslMode)
}

func MustConnect(o Options) *sqlx.DB {
	db, err := sqlx.Connect("postgres", o.toConnectionString())
	if err != nil {
		panic(err)
	}

	return db
}
