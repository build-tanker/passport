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
	// GET__ .../ping
	router.HandleFunc("/ping", pingHandler.Ping(ctx)).Methods(http.MethodGet)

	// GET__ .../login
	// POST_ .../v1/users source=google&access_token=tkn&name=name&email=email&user_id=123

	return router
}

func fakeHandler(ctx *appcontext.AppContext, db *sqlx.DB) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}
