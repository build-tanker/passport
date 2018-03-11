package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/config"
)

// Handler exposes all handlers
type Handler struct {
	pings  *pingHandler
	people *personHandler
	tokens *tokenHandler
}

// HTTPHandler is the type which can handle a URL
type httpHandler func(w http.ResponseWriter, r *http.Request)

// New creates a new handler
func New(conf *config.Config, db *sqlx.DB) *Handler {
	pings := newPingHandler()
	people := newPersonHandler(conf, db)
	tokens := newTokenHandler(conf, db)
	return &Handler{pings, people, tokens}
}

// Route pipes requests to the correct handlers
func (h *Handler) Route() http.Handler {
	router := mux.NewRouter()
	// GET__ .../ping
	router.HandleFunc("/ping", h.pings.ping()).Methods(http.MethodGet)

	// GET__ .../login
	router.HandleFunc("/login", h.people.login()).Methods(http.MethodGet)
	// GET_ .../v1/users source=google&access_token=tkn&name=name&email=email&user_id=123
	router.HandleFunc("/v1/users", h.people.signup()).Methods(http.MethodGet)

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
