package token

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/responses"
	"github.com/build-tanker/passport/pkg/token"
	"github.com/gorilla/mux"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

// Handler for people
type Handler interface {
}

type handler struct {
	conf   *config.Config
	tokens *token.Service
}

// NewHandler - create a new handler for people
func NewHandler(conf *config.Config, db *sqlx.DB) Handler {
	tokenService := token.New(conf, db)
	return &handler{
		conf:   conf,
		tokens: tokenService,
	}
}

func (h *handler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		// err := h.service.Add()
		// if err != nil {
		// 	responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("auth:signup:error", err.Error()))
		// 	return
		// }

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Success: "true",
		})
	}
}

func (h *handler) parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func (h *handler) parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
