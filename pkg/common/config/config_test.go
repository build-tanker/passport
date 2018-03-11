package config_test

import (
	"testing"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := config.NewConfig([]string{"./testutil"})
	assert.Equal(t, "dbname=passport user=passport password='passport' host=localhost port=5432 sslmode=disable", conf.Database().ConnectionString())
	assert.Equal(t, "postgres://passport:passport@localhost:5432/passport?sslmode=disable", conf.Database().ConnectionURL())
}