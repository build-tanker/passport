package token_test

import (
	"testing"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/postgres"
	"github.com/build-tanker/passport/pkg/token"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var sqlDB *sqlx.DB
var conf *config.Config

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

func cleanUpDatabase(db *sqlx.DB) error {
	_, err := db.Queryx("DELETE FROM token WHERE external_access_token='fakeExternalAccessToken'")
	if err != nil {
		return err
	}
	_, err = db.Queryx("DELETE FROM person WHERE email='fakeEmail'")
	if err != nil {
		return err
	}
	return nil
}

func prepareDatabase(db *sqlx.DB) error {
	_, err := db.Queryx("INSERT INTO person (id, source, name, email, picture_url, gender, source_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", "AAE2C7DA-859E-4066-B84F-D55E06EFAF70", "fakeSource", "fakeName", "fakeEmail", "fakeImage", "fakeGender", "fakeSourceId")
	if err != nil {
		return err
	}
	return nil
}

func TestTokenFlow(t *testing.T) {
	initConf()
	initDB()
	defer closeDB()

	tok := token.New(conf, sqlDB)

	cleanUpDatabase(sqlDB)

	_, err := tok.Add("AAE2C7DA-859E-4066-B84F-D55E06EFAF69", "fakeSource", "fakeExternalAccessToken", "fakeExternalRefreshToken", int64(3600), "fakeExternalTokenType")
	assert.Contains(t, err.Error(), "violates foreign key constraint")

	valid, err := tok.Validate("AAE2C7DA-859E-4066-B84F-D55E06EFAF69")
	assert.Nil(t, err)

	prepareDatabase(sqlDB)

	accessToken, err := tok.Add("AAE2C7DA-859E-4066-B84F-D55E06EFAF70", "fakeSource", "fakeExternalAccessToken", "fakeExternalRefreshToken", int64(3600), "fakeExternalTokenType")
	assert.Nil(t, err)

	valid, err = tok.Validate(accessToken)
	assert.Nil(t, err)
	assert.Equal(t, true, valid)

	err = tok.Remove(accessToken)
	assert.Nil(t, err)

	cleanUpDatabase(sqlDB)
}
