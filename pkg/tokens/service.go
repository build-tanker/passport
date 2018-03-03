package tokens

import (
	"github.com/build-tanker/passport/pkg/appcontext"
	"github.com/jmoiron/sqlx"
)

// Service - inteface for token service
type Service interface {
	Add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) error
}

type service struct {
	ctx       *appcontext.AppContext
	datastore Datastore
}

// NewService - create a new service for tokens
func NewService(ctx *appcontext.AppContext, db *sqlx.DB) Service {
	datastore := NewDatastore(ctx, db)
	return &service{ctx, datastore}
}

func (s *service) Add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) error {
	return nil
}
