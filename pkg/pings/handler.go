package pings

import (
	"net/http"

	"github.com/build-tanker/passport/pkg/appcontext"
	"github.com/build-tanker/passport/pkg/responses"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request)

type PingHandler struct{}

func (p *PingHandler) Ping(ctx *appcontext.AppContext) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.WriteJSON(w, http.StatusOK, responses.Response{Success: "pong"})
	}
}
