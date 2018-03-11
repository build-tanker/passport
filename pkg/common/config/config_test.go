package config_test

import (
	"testing"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := config.New([]string{"./testutil"})
	assert.Equal(t, "dbname=passportTest user=passportTest password='passportTest' host=localhost port=5432 sslmode=disable", conf.ConnectionString())
	assert.Equal(t, "postgres://passportTest:passportTest@localhost:5432/passportTest?sslmode=disable", conf.ConnectionURL())
}
