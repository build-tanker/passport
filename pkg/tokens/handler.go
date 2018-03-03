package tokens

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/appcontext"
	"github.com/build-tanker/passport/pkg/responses"
	"github.com/gorilla/mux"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

// Handler for people
type Handler interface {
}

type handler struct {
	ctx     *appcontext.AppContext
	service Service
}

// NewHandler - create a new handler for people
func NewHandler(ctx *appcontext.AppContext, db *sqlx.DB) Handler {
	s := NewService(ctx, db)
	return &handler{
		ctx:     ctx,
		service: s,
	}
}

func (h *handler) Add() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.service.Add()
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
