package people

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/passport/pkg/common/appcontext"
	"github.com/build-tanker/passport/pkg/common/responses"
	"github.com/gorilla/mux"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

// Handler for people
type Handler interface {
	Login() httpHandler
	Add() httpHandler
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

func (h *handler) Login() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := h.service.Login()
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
		err := h.service.Add(code)
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
