package config

import (
	"fmt"
	"log"

	"github.com/build-tanker/passport/pkg/translate"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for passport
type Config struct {
	server struct {
		port string
		host string
	}
	oauth2 struct {
		clientID     string
		clientSecret string
	}
	database struct {
		name        string
		host        string
		user        string
		password    string
		port        int
		maxPoolSize int
	}
}

// New creates a new configuration
func New(paths []string) *Config {
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
		log.Println(translate.T("config:change:reload"), e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

// Port - get the port from config
func (c *Config) Port() string {
	return c.server.port
}

// Host - get the host from config
func (c *Config) Host() string {
	return c.server.host
}

// OAuthClientID - get the oauth client id
func (c *Config) OAuthClientID() string {
	return c.oauth2.clientID
}

// OAuthClientSecret - get the oauth client secret
func (c *Config) OAuthClientSecret() string {
	return c.oauth2.clientSecret
}

// ConnectionString - get the connectionstring to connect to postgres
func (c *Config) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		c.database.name,
		c.database.user,
		c.database.password,
		c.database.host,
		c.database.port,
	)
}

// ConnectionURL - get the connection URL to connect to postgres
func (c *Config) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.database.user,
		c.database.password,
		c.database.host,
		c.database.port,
		c.database.name,
	)
}

// MaxPoolSize - get the max pool size for db connections
func (c *Config) MaxPoolSize() int {
	return c.database.maxPoolSize
}

func (c *Config) readLatestConfig() {
	c.server.port = viper.GetString("server.port")
	c.server.host = viper.GetString("server.host")

	c.oauth2.clientID = viper.GetString("oauth2.id")
	if c.oauth2.clientID == "" {
		log.Fatalln(translate.T("config:oauth2clientid:fail"))
	}

	c.oauth2.clientSecret = viper.GetString("oauth2.secret")
	if c.oauth2.clientSecret == "" {
		log.Fatalln(translate.T("config:oauth2clientsecret:fail"))
	}

	c.database.name = viper.GetString("database.name")
	c.database.host = viper.GetString("database.host")
	c.database.user = viper.GetString("database.user")
	c.database.password = viper.GetString("database.password")
	c.database.port = viper.GetInt("database.port")
	c.database.maxPoolSize = viper.GetInt("database.maxPoolSize")
}
