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
	ExpiresIn            int       `db:"expires_in" json:"expires_in,omitempty"`
	ExternalAccessToken  string    `db:"external_access_token" json:"external_access_token,omitempty"`
	ExternalRefreshToken string    `db:"external_refresh_token" json:"external_refresh_token,omitempty"`
	ExternalExpiresIn    int       `db:"external_expires_in" json:"external_expires_in,omitempty"`
	ExternalTokenType    string    `db:"external_token_type" json:"external_token_type,omitempty"`
	Deleted              bool      `db:"deleted" json:"deleted,omitempty"`
	CreatedAt            time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// Datastore for people
type store interface {
	add(person, source, externalAccessToken, externalRefreshToken string, externalExpiresIn int64, externalTokenType string) (string, error)
	remove(accessToken string) error
	validate(accessToken string) (Token, error)
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

func (s *persistentStore) add(person, source, externalAccessToken, externalRefreshToken string, externalExpiresIn int64, externalTokenType string) (string, error) {
	id := s.generateUUID()
	_, err := s.db.Queryx("INSERT INTO token (id, person, source, access_token, external_access_token, external_refresh_token, external_expires_in, external_token_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", id, person, source, id, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *persistentStore) validate(accessToken string) (Token, error) {
	rows, err := s.db.Queryx("SELECT * FROM token WHERE deleted = FALSE AND access_token=$1 AND now() < created_at + (expires_in::TEXT || ' secs')::INTERVAL AND now() > created_at LIMIT 1", accessToken)

	if err != nil {
		return Token{}, err
	}

	var token Token
	for rows.Next() {
		err = rows.StructScan(&token)
		if err != nil {
			return Token{}, err
		}
	}

	return token, nil
}

func (s *persistentStore) remove(accessToken string) error {
	_, err := s.db.Queryx("UPDATE token SET deleted = TRUE WHERE access_token=$1", accessToken)
	if err != nil {
		return err
	}
	return nil
}

func (s *persistentStore) generateUUID() string {
	return uuid.NewV4().String()
}
