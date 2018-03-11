package tokens

import (
	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service - inteface for token service
type Service interface {
	Add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) error
}

type service struct {
	conf      *config.Config
	datastore Datastore
}

// NewService - create a new service for tokens
func NewService(conf *config.Config, db *sqlx.DB) Service {
	datastore := NewDatastore(conf, db)
	return &service{conf, datastore}
}

func (s *service) Add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) error {
	return nil
}
