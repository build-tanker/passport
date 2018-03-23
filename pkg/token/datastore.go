package token

import (
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	"github.com/build-tanker/passport/pkg/common/config"
)

// Token saves token details for a user
type Token struct {
	ID                   string    `db:"id" json:"id,omitempty"`
	Person               string    `db:"person" json:"person,omitempty"`
	Source               string    `db:"source" json:"source,omitempty"`
	AccessToken          string    `db:"access_token" json:"access_token,omitempty"`
	ExternalAccessToken  string    `db:"external_access_token" json:"external_access_token,omitempty"`
	ExternalRefreshToken string    `db:"external_refresh_token" json:"external_refresh_token,omitempty"`
	ExternalExpiresIn    int       `db:"external_expires_in" json:"external_expires_in,omitempty"`
	ExternalTokenType    string    `db:"external_token_type" json:"external_token_type,omitempty"`
	CreatedAt            time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// Datastore for people
type store interface {
	add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) (string, error)
}

type persistentStore struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewDatastore - create a new datastore for people
func newStore(conf *config.Config, db *sqlx.DB) store {
	return &persistentStore{
		conf: conf,
		db:   db,
	}
}

func (s *persistentStore) add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) (string, error) {
	id := s.generateUUID()
	// #TODO
	return id, nil
}

func (s *persistentStore) generateUUID() string {
	return uuid.NewV4().String()
}
