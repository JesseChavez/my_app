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

	// App authentication secret key, recomended 32 or 64 bytes
	enki.SecretAuthKey = "b0ae69a9f367ad86fa50b4098e1ee9aabe6a6012a91f06248f084f205793a3cd"

	// App encryption secret key, must be either 16, 24, or 32 bytes to select
	// AES-128, AES-192, or AES-256 modes
	enki.SecretEncrKey = "7eb98e9534f49eb07157874c7adeb6cf"

	// Web App port, default is "3000"
	enki.WebPort = spt.FetchEnv("PORT", "3000")

	app := enki.New("my_app")

	return app
}
