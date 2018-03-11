package people

import (
	"errors"
	"fmt"
	"log"

	"github.com/build-tanker/oauth2"
	"github.com/build-tanker/passport/pkg/common/config"
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
	conf      *config.Config
	datastore Datastore
	oauth     oauth2.OAuth2
	tokens    tokens.Service
}

// NewService - create a new service for people
func NewService(conf *config.Config, db *sqlx.DB) Service {
	datastore := NewDatastore(conf, db)

	clientID := conf.OAuthClientID()
	clientSecret := conf.OAuthClientSecret()
	redirctURL := fmt.Sprintf("%s%s", conf.Host(), redirectURLPath)
	tokens := tokens.NewService(conf, db)

	oauth, err := oauth2.NewOAuth2(clientID, clientSecret, redirctURL)
	if err != nil {
		log.Fatalln(translate.T("people:oauth:failed"))
	}

	return &service{conf, datastore, oauth, tokens}
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
	verified, accessToken, tokenType, expiresIn, refreshToken, _, _, err := s.oauth.GetAndVerifyToken(code)
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
	personID := ""

	person, err := s.datastore.ViewBySourceID(id)
	if err != nil {
		// Saving person if not found
		personID, err = s.datastore.Add("google", name, email, image, gender, id)
		if err != nil {
			return err
		}
	}
	// Otherwise use the found person
	personID = person.ID

	// Add a new token
	err = s.tokens.Add(personID, "google", accessToken, refreshToken, expiresIn, tokenType)
	if err != nil {
		return err
	}

	return nil
}
