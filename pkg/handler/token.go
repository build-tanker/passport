package handler

import (
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

func (t *tokenHandler) add() httpHandler {
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
