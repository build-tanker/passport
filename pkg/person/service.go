package person

import (
	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/oauth2"
	"github.com/build-tanker/passport/pkg/token"
	"github.com/jmoiron/sqlx"
)

// Service for people
type Service struct {
	conf   *config.Config
	store  store
	oauth  oauth2.OAuth2
	tokens *token.Service
}

// New - create a new service for people
func New(conf *config.Config, db *sqlx.DB, oauth oauth2.OAuth2, tokens *token.Service) *Service {
	store := newStore(conf, db)
	return &Service{conf, store, oauth, tokens}
}

// Login a person
func (s *Service) Login() (string, error) {
	url, err := s.oauth.GetAuthURL("", "", "", "", "", "")
	if err != nil {
		return "", err
	}
	return url, nil
}

// Logout an accessToken
func (s *Service) Logout(accessToken string) error {
	return s.tokens.Remove(accessToken)
}

// Verify or add a person
func (s *Service) Verify(code string) (string, error) {
	verifyDetails, err := s.oauth.GetAndVerifyToken(code)
	if err != nil {
		return "", err
	}

	// Get Profile details
	profileDetails, err := s.oauth.GetProfileDetails(verifyDetails.AccessToken)
	if err != nil {
		return "", err
	}

	person, err := s.store.viewBySourceID(profileDetails.ID)
	if err != nil {
		return "", err
	}

	personID := person.ID
	// If token is verified, save person and token details

	if personID == "" {
		// Saving person if not found
		personID, err = s.store.add("google", profileDetails.Name, profileDetails.Email, profileDetails.Image, profileDetails.Gender, profileDetails.ID)
		if err != nil {
			return "", err
		}
	}

	// Add a new token
	accessToken, err := s.tokens.Add(personID, "google", verifyDetails.AccessToken, verifyDetails.RefreshToken, verifyDetails.ExpiresIn, verifyDetails.TokenType)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
