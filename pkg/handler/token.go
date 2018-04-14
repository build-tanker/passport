package handler

import (
	"fmt"
	"net/http"

	"github.com/build-tanker/passport/pkg/common/responses"
	"github.com/build-tanker/passport/pkg/token"
)

type tokenHandler struct {
	tokens *token.Service
}

func newTokenHandler(service *token.Service) *tokenHandler {
	return &tokenHandler{
		tokens: service,
	}
}

func (t *tokenHandler) validate() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := parseKeyFromQuery(r, "accessToken")
		if accessToken == "" {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("token:validate:error", "AccessToken missing"))
			return
		}
		valid, person, err := t.tokens.Validate(accessToken)
		if err != nil {
			responses.WriteJSON(w, http.StatusBadRequest, responses.NewErrorResponse("token:validate:error", err.Error()))
			return
		}

		responses.WriteJSON(w, http.StatusOK, &responses.Response{
			Data: struct {
				Person string `json:"person"`
			}{
				Person: person,
			},
			Success: fmt.Sprintf("%v", valid),
		})
	}
}
