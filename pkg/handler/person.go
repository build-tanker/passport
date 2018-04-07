package handler

import (
	"net/http"

	"github.com/build-tanker/passport/pkg/common/responses"
	"github.com/build-tanker/passport/pkg/person"
)

type personHandler struct {
	people *person.Service
}

func newPersonHandler(service *person.Service) *personHandler {
	return &personHandler{
		people: service,
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

func (p *personHandler) verify() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		code := parseKeyFromQuery(r, "code")
		accessToken, err := p.people.Verify(code)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("auth:verify:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: struct {
				AccessToken string `json:"accessToken"`
			}{
				AccessToken: accessToken,
			},
			Success: "true",
		})
	}
}
