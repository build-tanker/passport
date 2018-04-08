package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/oauth2"
	"github.com/build-tanker/passport/pkg/person"
	"github.com/build-tanker/passport/pkg/token"
)

const (
	oauthRedirectURLPath = "/v1/users/verify"
)

// Handler exposes all handlers
type Handler struct {
	health *healthHandler
	people *personHandler
	tokens *tokenHandler
}

// HTTPHandler is the type which can handle a URL
type httpHandler func(w http.ResponseWriter, r *http.Request)

// New creates a new handler
func New(conf *config.Config, db *sqlx.DB) *Handler {

	// Create oAuth for person
	clientID := conf.OAuthClientID()
	clientSecret := conf.OAuthClientSecret()
	redirctURL := fmt.Sprintf("%s%s", conf.Host(), oauthRedirectURLPath)
	oauth, err := oauth2.NewOAuth2(clientID, clientSecret, redirctURL)
	if err != nil {
		log.Fatalln("Could not initialise OAuth2 Client")
	}

	// Create token
	tokenService := token.New(conf, db)

	// Create person
	personService := person.New(conf, db, oauth, tokenService)

	// Finally, create handlers
	health := newHealthHandler()
	people := newPersonHandler(personService)
	tokens := newTokenHandler(tokenService)

	return &Handler{health, people, tokens}
}

// Route pipes requests to the correct handlers
func (h *Handler) Route() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping", h.health.ping()).Methods(http.MethodGet)

	router.HandleFunc("/v1/users/login", h.people.login()).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/verify", h.people.verify()).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/logout", h.people.logout()).Methods(http.MethodGet)

	return router
}

func fakeHandler() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}

func parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
