package people

import (
	"fmt"

	"github.com/build-tanker/oauth2"
	"github.com/build-tanker/passport/pkg/appcontext"
	"github.com/build-tanker/passport/pkg/translate"
	"github.com/jmoiron/sqlx"
)

const (
	redirectURLPath = "/v1/users"
)

// Service - inteface for people service
type Service interface {
	Login() (string, error)
	Add(source, name, email, picture string) error
}

type service struct {
	ctx       *appcontext.AppContext
	datastore Datastore
	oauth     oauth2.OAuth2
}

// NewService - create a new service for people
func NewService(ctx *appcontext.AppContext, db *sqlx.DB) Service {
	datastore := NewDatastore(ctx, db)

	clientID := ctx.GetConfig().OAuthClientID()
	clientSecret := ctx.GetConfig().OAuthClientSecret()
	redirctURL := fmt.Sprintf("%s%s", ctx.GetConfig().Host(), redirectURLPath)

	oauth, err := oauth2.NewOAuth2(clientID, clientSecret, redirctURL)
	if err != nil {
		ctx.GetLogger().Fatalln(translate.T("people:oauth:failed"))
	}

	return &service{ctx, datastore, oauth}
}

func (s *service) Login() (string, error) {
	url, err := s.oauth.GetAuthURL("", "", "", "", "", "")
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *service) Add(source, name, email, picture string) error {
	_, err := s.datastore.Add(source, name, email, picture)
	return err
}
