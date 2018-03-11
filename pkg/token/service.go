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
func (s *Service) Add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) error {
	return nil
}
