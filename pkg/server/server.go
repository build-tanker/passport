package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"

	"github.com/build-tanker/passport/pkg/appcontext"
	"github.com/build-tanker/passport/pkg/translate"
)

// Server holds the web server
type Server struct {
	ctx    *appcontext.AppContext
	db     *sqlx.DB
	server *http.Server
}

// NewServer initialises a new server
func NewServer(ctx *appcontext.AppContext, db *sqlx.DB) *Server {
	return &Server{
		ctx: ctx,
		db:  db,
	}
}

// Start a new server
func (s *Server) Start() error {
	config := s.ctx.GetConfig()
	log := s.ctx.GetLogger()

	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())

	router := Router(s.ctx, s.db)
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	serverURL := fmt.Sprintf(":%s", config.Port())

	server.Use(recover())
	server.UseHandler(router)
	server.Run(serverURL)

	s.server = &http.Server{
		Addr:         serverURL,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Handler:      server,
	}

	log.Infoln(translate.T("server:negroni:listen"), serverURL)
	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			if err.Error() != "http: Server closed" { // This is an error response not an error code
				fmt.Println(translate.T("server:listen:fail"), err.Error())
			}
		}
	}()

	http.ListenAndServe(serverURL, server)

	return nil
}

// Stop the server
func (s *Server) Stop() error {
	// Not sure how to stop a server
	s.server.Shutdown(nil)
	return nil
}

func recover() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf(translate.T("server:panic:recover"), err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next(w, r)
	})
}
