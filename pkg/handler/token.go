package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/build-tanker/passport/pkg/common/responses"
	"github.com/build-tanker/passport/pkg/token"
	"github.com/gorilla/mux"
)

type tokenHandler struct {
	tokens *token.Service
}

func newTokenHandler(conf *config.Config, db *sqlx.DB) *tokenHandler {
	return &tokenHandler{
		tokens: token.New(conf, db),
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

func (t *tokenHandler) parseKeyFromQuery(r *http.Request, key string) string {
	value := ""
	if len(r.URL.Query()[key]) > 0 {
		value = r.URL.Query()[key][0]
	}
	return value
}

func (t *tokenHandler) parseKeyFromVars(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
