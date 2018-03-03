package translate

type translate struct {
	translations map[string]string
	populated    bool
}

var t *translate

// T finds the message with the correct code and sends the tranlation back
func T(code string) string {
	if t == nil {
		populate()
	}

	value, ok := t.translations[code]
	if !ok {
		value = "Message not found"
	}
	return value
}

func populate() {
	t = &translate{
		populated: false,
	}

	t.translations = make(map[string]string)
	t.translations["config:oauth2clientid:fail"] = "Please enter the OAuth2 Client ID"
	t.populated = true
}
