package person

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/responses"
	"github.com/build-tanker/passport/pkg/person"
	"github.com/gorilla/mux"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

// Handler for people
type Handler interface {
	Login() httpHandler
	Add() httpHandler
}

type handler struct {
	conf   *config.Config
	person *person.Service
}

// NewHandler - create a new handler for people
func NewHandler(conf *config.Config, db *sqlx.DB) Handler {
	personService := person.New(conf, db)
	return &handler{
		conf:   conf,
		person: personService,
	}
}

func (h *handler) Login() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := h.person.Login()
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("people:login:error", err.Error()))
			return
		}

		http.Redirect(w, r, url, 301)
	}
}

func (h *handler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		code := h.parseKeyFromQuery(r, "code")
		err := h.person.Add(code)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("auth:signup:error", err.Error()))
			return
		}

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
