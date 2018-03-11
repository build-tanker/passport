package handler

import (
	"net/http"

	"github.com/build-tanker/tanker/pkg/responses"
)

type pingHandler struct{}

func newPingHandler() *pingHandler {
	return &pingHandler{}
}

func (p *pingHandler) ping() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.WriteJSON(w, http.StatusOK, responses.Response{Success: "pong"})
	}
}
