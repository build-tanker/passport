package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/appcontext"
	"github.com/build-tanker/passport/pkg/pings"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

// Router pipes requests to the correct handlers
func Router(ctx *appcontext.AppContext, db *sqlx.DB) http.Handler {

	pingHandler := pings.PingHandler{}

	router := mux.NewRouter()
	// GET___ .../ping
	router.HandleFunc("/ping", pingHandler.Ping(ctx)).Methods(http.MethodGet)

	// Auth
	// POST .../v1/users source=google&access_token=tkn&name=name&email=email&user_id=123
	// GET .../v1/users/15
	// PUT .../v1/users/15 access_token=tkn&name=name&deleted=true
	// DELETE .../v1/users/15

	return router
}

func fakeHandler(ctx *appcontext.AppContext, db *sqlx.DB) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}
