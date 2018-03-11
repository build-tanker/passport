package postgres

import (
	"log"
	"time"

	"github.com/build-tanker/passport/pkg/translate"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

const connMaxLifetime = 30 * time.Minute

// New initialize a new postgres connection
func New(url string, maxOpenConns int) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatalln(translate.T("postgres:connection:failed"), err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(translate.T("postgres:ping:failed"), err.Error(), translate.T("postgres:ping:failed:2"), url)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	log.Println(translate.T("postgres:connection:success"))

	return db
}
