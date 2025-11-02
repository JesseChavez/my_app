package config

import (
	"my_app"

	"github.com/JesseChavez/enki"
	"github.com/JesseChavez/spt"
)

func NewApplication() enki.Enki {
	// Web App context path, app runs in subfolder, default is root "/"
	// enki.ContextPath = "/my_app"

	// App resocurces config files, templates, assets, etc.
	enki.Resources = resources.LoadFS()

	// App time zone, default is UTC
	enki.TimeZone = "Australia/Melbourne"

	// App session key, session cookie name
	enki.SessionKey = "_my_app_enc_session"

	// App session max age in minutes
	enki.SessionMaxAge = 120

	// App secret key base, recomended 32 or 64 bytes
	enki.SecretKeyBase = "my app secret"

	// App log level "info" or "debug"
	enki.LogLevel = "debug"

	// App salt for cookie encryption
	// enki.AuthenticatedEncryptedCookieSalt = "encrypted cookie salt"

	// Application mode, when API is true response is always JSON
	// enki.API = false

	// Default rendering mode, default is SSR
	// enki.CSR = true

	// Web App port, default is "3000"
	enki.WebPort = spt.FetchEnv("PORT", "3000")

	app := enki.New("my_app")

	return app
}
