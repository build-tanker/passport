package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/handler/v1/person"
	"github.com/build-tanker/passport/pkg/handler/v1/ping"
)

// HTTPHandler is the type which can hanlde a URL
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// Router pipes requests to the correct handlers
func Router(conf *config.Config, db *sqlx.DB) http.Handler {

	pingHandler := pings.PingHandler{}
	personHandler := person.NewHandler(conf, db)

	router := mux.NewRouter()
	// GET__ .../ping
	router.HandleFunc("/ping", pingHandler.Ping()).Methods(http.MethodGet)

	// GET__ .../login
	router.HandleFunc("/login", personHandler.Login()).Methods(http.MethodGet)
	// GET_ .../v1/users source=google&access_token=tkn&name=name&email=email&user_id=123
	router.HandleFunc("/v1/users", personHandler.Add()).Methods(http.MethodGet)

	return router
}

func fakeHandler(conf *config.Config, db *sqlx.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}
