package ping

import (
	"net/http"

	"github.com/build-tanker/passport/pkg/common/responses"
)

// HTTPHandler type is a function which can be used an a handler
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// PingHandler is a structure to handle ping and related functionality
type PingHandler struct{}

// Ping handles the /ping request
func (p *PingHandler) Ping() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.WriteJSON(w, http.StatusOK, responses.Response{Success: "pong"})
	}
}
