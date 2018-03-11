package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/people"
	"github.com/build-tanker/passport/pkg/pings"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

// Router pipes requests to the correct handlers
func Router(conf *config.Config, db *sqlx.DB) http.Handler {

	pingHandler := pings.PingHandler{}
	peopleHandler := people.NewHandler(conf, db)

	router := mux.NewRouter()
	// GET__ .../ping
	router.HandleFunc("/ping", pingHandler.Ping()).Methods(http.MethodGet)

	// GET__ .../login
	router.HandleFunc("/login", peopleHandler.Login()).Methods(http.MethodGet)
	// POST_ .../v1/users source=google&access_token=tkn&name=name&email=email&user_id=123
	router.HandleFunc("/v1/users", peopleHandler.Add()).Methods(http.MethodGet)

	// http://localhost:3000/v1/users?code=4%2FAACqC5q123CTDUUKHCepsD-1fbYupBtGKDw_8TnN8Dk-DR4b3Y6kMfdEn2miNtMNCbyYyvV7W5MfTp9v5tV5zw8&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile#

	return router
}

func fakeHandler(conf *config.Config, db *sqlx.DB) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"fake\":\"yes\"}"))
	}
}
