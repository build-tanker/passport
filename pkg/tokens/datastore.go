package tokens

import (
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	"github.com/build-tanker/passport/pkg/common/appcontext"
	"github.com/build-tanker/passport/pkg/common/logger"
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
type Datastore interface {
	Add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) (string, error)
}

type datastore struct {
	ctx *appcontext.AppContext
	db  *sqlx.DB
	log logger.Logger
}

// NewDatastore - create a new datastore for people
func NewDatastore(ctx *appcontext.AppContext, db *sqlx.DB) Datastore {
	return &datastore{
		ctx: ctx,
		db:  db,
		log: ctx.GetLogger(),
	}
}

func (s *datastore) Add(person, source, externalAccessToken, externalRefreshToken, externalExpiresIn, externalTokenType string) (string, error) {
	id := s.generateUUID()
	return id, nil
}

func (s *datastore) generateUUID() string {
	return uuid.NewV4().String()
}
