package store

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func getTestDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 sslmode=disable user=postgres")
	if err != nil {
		panic(err)
	}

	return db
}

func TestInsertPlayer(t *testing.T) {
	s := Store{db: getTestDb()}
	p, err := s.InsertPlayer("Gisli")
	if err != nil {
		t.Logf("s.InsertPlayer returned error: %v", err)
		t.FailNow()
	}

	assert.Equal(t, "Gisli", p.Name)

}
