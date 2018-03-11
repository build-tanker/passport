package appcontext

import (
	"github.com/build-tanker/passport/pkg/common/config"
)

// AppContext - global context for config and logging
type AppContext struct {
	config *config.Config
}

// NewAppContext - function to create a global context for conf and logging
func NewAppContext(config *config.Config) *AppContext {
	return &AppContext{
		config: config,
	}
}

// GetConfig - fetch the config from the global AppContext
func (a *AppContext) GetConfig() *config.Config {
	return a.config
}
