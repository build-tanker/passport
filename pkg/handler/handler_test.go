package handler

import "github.com/build-tanker/passport/pkg/common/config"

var testConfig *config.Config

func NewTestConfig() *config.Config {
	if testConfig == nil {
		testConfig = config.New([]string{".", "..", "../.."})
	}
	return testConfig
}
