package people

import (
	"errors"
	"fmt"

	"github.com/build-tanker/oauth2"
	"github.com/build-tanker/passport/pkg/appcontext"
	"github.com/build-tanker/passport/pkg/tokens"
	"github.com/build-tanker/passport/pkg/translate"
	"github.com/jmoiron/sqlx"
)

const (
	redirectURLPath = "/v1/users"
)

// Service - inteface for people service
type Service interface {
	Login() (string, error)
	Add(code string) error
}

type service struct {
	ctx       *appcontext.AppContext
	datastore Datastore
	oauth     oauth2.OAuth2
	tokens    tokens.Service
}

// NewService - create a new service for people
func NewService(ctx *appcontext.AppContext, db *sqlx.DB) Service {
	datastore := NewDatastore(ctx, db)

	clientID := ctx.GetConfig().OAuthClientID()
	clientSecret := ctx.GetConfig().OAuthClientSecret()
	redirctURL := fmt.Sprintf("%s%s", ctx.GetConfig().Host(), redirectURLPath)
	tokens := tokens.NewService(ctx, db)

	oauth, err := oauth2.NewOAuth2(clientID, clientSecret, redirctURL)
	if err != nil {
		ctx.GetLogger().Fatalln(translate.T("people:oauth:failed"))
	}

	return &service{ctx, datastore, oauth, tokens}
}

func (s *service) Login() (string, error) {
	url, err := s.oauth.GetAuthURL("", "", "", "", "", "")
	fmt.Println(url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *service) Add(code string) error {
	verified, accessToken, tokenType, expiresIn, refreshToken, idToken, userID, err := s.oauth.GetAndVerifyToken(code)
	if err != nil {
		return err
	}

	if !verified {
		return errors.New("The token from google could not be verified")
	}

	// Get Profile details
	email, name, image, id, gender, err := s.oauth.GetProfileDetails(accessToken)
	if err != nil {
		return err
	}

	// If token is verified, save person and token details
	// Saving person
	personID, err := s.datastore.Add("google", name, email, image, gender, id)
	if err != nil {
		return err
	}
	// Saving token
	err := s.tokens.Add(personID, "google", accessToken, refreshToken, expiresIn, tokenType)
	if err != nil {
		return err
	}

	return nil
}
