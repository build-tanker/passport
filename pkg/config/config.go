package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for passport
type Config struct {
	port              string
	host              string
	database          DatabaseConfig
	oauthClientID     string
	oauthClientSecret string
}

// NewConfig - create a new configuration
func NewConfig(paths []string) *Config {
	config := &Config{}

	viper.AutomaticEnv()

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("passport")
	viper.SetConfigType("toml")

	viper.SetDefault("server.port", "4000")
	viper.SetDefault("server.host", "http://localhost")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

// Port - get the port from config
func (c *Config) Port() string {
	return c.port
}

// Host - get the host from config
func (c *Config) Host() string {
	return c.host
}

// OAuthClientID - get the oauth client id
func (c *Config) OAuthClientID() string {
	return c.oauthClientID
}

// OAuthClientSecret - get the oauth client secret
func (c *Config) OAuthClientSecret() string {
	return c.oauthClientSecret
}

// Database - load the database config
func (c *Config) Database() DatabaseConfig {
	return c.database
}

func (c *Config) readLatestConfig() {
	c.port = viper.GetString("server.port")
	c.host = viper.GetString("server.host")

	c.oauthClientID = viper.GetString("oauth2.id")
	if c.oauthClientID == "" {
		log.Fatalln("Please enter the OAuth2 Client ID")
	}

	c.oauthClientSecret = viper.GetString("oauth2.secret")
	if c.oauthClientSecret == "" {
		log.Fatalln("Please enter the OAuth2 Client Secret")
	}

	c.database = NewDatabaseConfig()
}
