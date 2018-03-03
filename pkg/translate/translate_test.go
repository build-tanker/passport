package translate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/build-tanker/passport/pkg/translate"
)

func TestT(t *testing.T) {
	assert.Equal(t, "Please enter the OAuth2 Client ID", translate.T("config:oauth2clientid:fail"))
	assert.Equal(t, "Message not found", translate.T("this:does:not:exist"))
}
