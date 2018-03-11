package handler

import (
	"net/http"

	"github.com/build-tanker/tanker/pkg/responses"
)

type ping struct{}

func (p *ping) getPing() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.WriteJSON(w, http.StatusOK, responses.Response{Success: "pong"})
	}
}
