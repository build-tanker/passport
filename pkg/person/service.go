package person

import (
	"errors"
	"fmt"
	"log"

	"github.com/build-tanker/oauth2"
	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/token"
	"github.com/build-tanker/passport/pkg/translate"
	"github.com/jmoiron/sqlx"
)

const (
	redirectURLPath = "/v1/users"
)

// Service for people
type Service struct {
	conf   *config.Config
	store  store
	oauth  oauth2.OAuth2
	tokens *token.Service
}

// New - create a new service for people
func New(conf *config.Config, db *sqlx.DB) *Service {
	store := newStore(conf, db)

	clientID := conf.OAuthClientID()
	clientSecret := conf.OAuthClientSecret()
	redirctURL := fmt.Sprintf("%s%s", conf.Host(), redirectURLPath)
	tokens := token.New(conf, db)

	oauth, err := oauth2.NewOAuth2(clientID, clientSecret, redirctURL)
	if err != nil {
		log.Fatalln(translate.T("people:oauth:failed"))
	}

	return &Service{conf, store, oauth, tokens}
}

// Login a person
func (s *Service) Login() (string, error) {
	url, err := s.oauth.GetAuthURL("", "", "", "", "", "")
	fmt.Println(url)
	if err != nil {
		return "", err
	}
	return url, nil
}

// Add or signup a person
func (s *Service) Add(code string) error {
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

	person, err := s.store.viewBySourceID(id)
	if err != nil {
		// Saving person if not found
		personID, err = s.store.add("google", name, email, image, gender, id)
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