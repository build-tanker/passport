package person_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/build-tanker/passport/pkg/oauth2"
	"github.com/build-tanker/passport/pkg/person"
	"github.com/build-tanker/passport/pkg/token"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/postgres"
	"github.com/jmoiron/sqlx"
)

var sqlDB *sqlx.DB
var conf *config.Config
var oauthMock oauth2.OAuth2

// Mock OAuth
type mockOauth struct{}

func (m mockOauth) GetAuthURL(scope, accessType, state, includeGrantedScopes, loginHint, prompt string) (string, error) {
	return "fakeAuthURL", nil
}
func (m mockOauth) GetToken(code string) ([]byte, error) {
	return nil, nil
}
func (m mockOauth) VerifyToken(accessToken string) (string, error) {
	return "", nil
}
func (m mockOauth) RefreshToken(refreshToken string) (string, error) {
	return "", nil
}
func (m mockOauth) RevokeToken(accessToken string) error {
	return nil
}
func (m mockOauth) GetProfileDetails(accessToken string) (oauth2.ProfileDetails, error) {
	return oauth2.ProfileDetails{
		Email:  "fakeEmail",
		Name:   "fakeName",
		Image:  "fakeImage",
		ID:     "fakeID",
		Gender: "fakeGender",
	}, nil
}
func (m mockOauth) GetAndVerifyToken(code string) (oauth2.VerifyDetails, error) {
	return oauth2.VerifyDetails{
		AccessToken:  "fakeAccessToken",
		TokenType:    "fakeTokenType",
		ExpiresIn:    "fakeExpiresIn",
		RefreshToken: "fakeRefreshToken",
		IDToken:      "fakeIDToken",
		UserID:       "fakeUserID",
	}, nil
}

// Initialise
func initDB() {
	if sqlDB == nil {
		sqlDB = postgres.New(conf.ConnectionURL(), conf.MaxPoolSize())
	}
}

func initConf() {
	if conf == nil {
		conf = config.New([]string{".", "..", "../.."})
	}
}

func initMockOauth() {
	oauthMock = &mockOauth{}
}

// Test
func TestPersonFlow(t *testing.T) {
	initConf()
	initDB()
	initMockOauth()
	tokens := token.New(conf, sqlDB)

	p := person.New(conf, sqlDB, oauthMock, tokens)

	url, _ := p.Login()
	assert.Equal(t, "fakeAuthURL", url)

	err := p.Verify("abc")
	assert.Nil(t, err)
}
