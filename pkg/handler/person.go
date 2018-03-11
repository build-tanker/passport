package handler

import (
	"net/http"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/responses"
	"github.com/build-tanker/passport/pkg/person"
	"github.com/jmoiron/sqlx"
)

type personHandler struct {
	people *person.Service
}

func newPersonHandler(conf *config.Config, db *sqlx.DB) *personHandler {
	return &personHandler{
		people: person.New(conf, db),
	}
}

func (p *personHandler) login() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := p.people.Login()
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("people:login:error", err.Error()))
			return
		}

		http.Redirect(w, r, url, 301)
	}
}

func (p *personHandler) signup() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		code := parseKeyFromQuery(r, "code")
		err := p.people.Add(code)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("auth:signup:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Success: "true",
		})
	}
}
