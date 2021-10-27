package store

import "github.com/jmoiron/sqlx"

const migration = `
create schema if not exists gammon;

create table gammon.player (
    id serial primary key,
    name varchar(256) not null,
    elo integer
);
`

func migrate(db *sqlx.DB) error {
	_, err := db.Exec(migration)
	return err
}
