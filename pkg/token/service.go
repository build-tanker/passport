package token

import (
	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// New - create a new service for tokens
func New(conf *config.Config, db *sqlx.DB) *Service {
	datastore := newStore(conf, db)
	return &Service{conf, datastore}
}

// Service to handle tokens
type Service struct {
	conf  *config.Config
	store store
}

// Add a token
func (s *Service) Add(person, source, externalAccessToken, externalRefreshToken string, externalExpiresIn int64, externalTokenType string) (string, error) {
	accessToken, err := s.store.add(person, "google", externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// Remove a token
func (s *Service) Remove(accessToken string) error {
	return s.store.remove(accessToken)
}

// Validate a token
func (s *Service) Validate(accessToken string) (bool, string, error) {
	token, err := s.store.validate(accessToken)
	if err != nil {
		return false, "", err
	}

	return (token.ID != ""), token.Person, nil
}
