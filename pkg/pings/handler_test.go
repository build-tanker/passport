package pings

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/build-tanker/passport/pkg/common/config"
)

var pingHandlerTestConfig *config.Config

func TestPingHandler(t *testing.T) {
	pingHandler := PingHandler{}

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler.Ping())

	handler.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "{\"success\":\"pong\"}\n", response.Body.String())
}

func NewPingHandlerTestConfig() *config.Config {
	if pingHandlerTestConfig == nil {
		pingHandlerTestConfig = config.New([]string{".", "..", "../.."})
	}
	return pingHandlerTestConfig
}
