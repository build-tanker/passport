package server

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/urfave/negroni"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/handler"
)

// Server holds the web server
type Server struct {
	conf   *config.Config
	db     *sqlx.DB
	server *http.Server
}

// New initialises a new server
func New(conf *config.Config, db *sqlx.DB) *Server {
	return &Server{
		conf: conf,
		db:   db,
	}
}

// Start a new server
func (s *Server) Start() error {
	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	router := handler.Router(s.conf, s.db)
	serverURL := fmt.Sprintf(":%s", s.conf.Port())

	server.UseHandler(router)
	server.Run(serverURL)
	return nil
}

// Stop the server
func (s *Server) Stop() error {
	// Not sure how to stop a server
	s.server.Shutdown(nil)
	return nil
}
