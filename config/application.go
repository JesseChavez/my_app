package config

import (
	"my_app"

	"github.com/JesseChavez/enki"
	"github.com/JesseChavez/spt"
)

func NewApplication() enki.Enki {
	/* Web App context path, app runs in subfolder, default is root "/" */
	// enki.ContextPath = "/admin"

	// project or app root directory
	enki.BaseDir = resources.BaseDir()

	// App resocurces config files, templates, assets, etc.
	enki.Resources = resources.LoadFS()

	// App time zone, default is UTC
	enki.TimeZone = "Australia/Melbourne"

	// Web App port, default is "3000"
	enki.WebPort = spt.FetchEnv("PORT", "8000")

	app := enki.New("my_app")

	return app
}
