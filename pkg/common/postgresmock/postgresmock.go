package postgresmock

import (
	"log"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/build-tanker/passport/pkg/translate"
	"github.com/jmoiron/sqlx"
)

// New creates a new mock sqlx db
func New() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf(translate.T("postgresmock:connection:fail"))
	}
	sqlxDB := sqlx.NewDb(db, "postgres")
	sqlxDB.SetMaxOpenConns(10)
	return sqlxDB, mock
}

// Close the mock connection
func Close(db *sqlx.DB) {
	db.Close()
}
