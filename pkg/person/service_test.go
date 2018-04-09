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
		ExpiresIn:    3600,
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

func closeDB() {
	if sqlDB != nil {
		sqlDB.Close()
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

func cleanUpDatabase(db *sqlx.DB) error {
	_, err := db.Queryx("DELETE FROM token WHERE external_access_token='fakeAccessToken'")
	if err != nil {
		return err
	}
	_, err = db.Queryx("DELETE FROM person WHERE email='fakeEmail'")
	if err != nil {
		return err
	}
	return nil
}

// Test
func TestPersonFlow(t *testing.T) {
	initConf()
	initDB()
	defer closeDB()

	initMockOauth()
	tokens := token.New(conf, sqlDB)

	cleanUpDatabase(sqlDB)

	p := person.New(conf, sqlDB, oauthMock, tokens)

	url, _ := p.Login()
	assert.Equal(t, "fakeAuthURL", url)

	accessToken, err := p.Verify("abc")
	assert.Nil(t, err)
	assert.Equal(t, 36, len(accessToken))

	valid, err := tokens.Validate(accessToken)
	assert.Nil(t, err)
	assert.Equal(t, true, valid)

	err = p.Logout(accessToken)
	assert.Nil(t, err)

	cleanUpDatabase(sqlDB)

}
