package postgresmock

import (
	"log"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/build-tanker/passport/pkg/translate"
	"github.com/jmoiron/sqlx"
)

// NewMockSqlxDB - create a new mock sqlx db
func NewMockSqlxDB() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf(translate.T("postgresmock:connection:fail"))
	}
	sqlxDB := sqlx.NewDb(db, "postgres")
	sqlxDB.SetMaxOpenConns(10)
	return sqlxDB, mock
}

// CloseMockSqlxDB - close the mock connection
func CloseMockSqlxDB(db *sqlx.DB) {
	db.Close()
}
