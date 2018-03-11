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

	t.translations["passport:app:usage"] = "this service saves files and makes them available for distribution"
	t.translations["passport:app:start"] = "Starting passport"
	t.translations["passport:cli:start"] = "start the service"
	t.translations["passport:cli:migrate"] = "run database migrations"
	t.translations["passport:cli:rollback"] = "rollback the latest database migration"

	t.translations["config:change:reload"] = "Config file %s was edited, reloading config\n"
	t.translations["config:oauth2clientid:fail"] = "Please enter the OAuth2 Client ID"
	t.translations["config:oauth2clientsecret:fail"] = "Please enter the OAuth2 Client Secret"

	t.translations["postgres:migration:up:fail"] = "Sadly, found no new migrations to run"
	t.translations["postgres:migration:up:success"] = "Migration has been successfully done"
	t.translations["postgres:migration:down:fail"] = "We have already removed all migrations"
	t.translations["postgres:migration:down:success"] = "Rollback Successful"

	t.translations["postgres:connection:failed"] = "Could not connect to database:"
	t.translations["postgres:ping:failed"] = "Ping to the database failed:"
	t.translations["postgres:ping:failed:2"] = "on connString"
	t.translations["postgres:connection:success"] = "Connected to database"

	t.translations["postgresmock:connection:fail"] = "Should not get an error in creating a mock database"

	t.translations["people:oauth:failed"] = "Could not initialise OAuth2 Client"

	t.populated = true
}
