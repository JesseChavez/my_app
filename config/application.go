package config

import (
	"my_app"

	"github.com/JesseChavez/enki"
	"github.com/JesseChavez/spt"
)

func RunApplication() {
	/* App context path, app runs in subfolder, default is root "/" */
	// enki.ContextPath = "/admin"

	/* App resocurces config files, templates, assets, etc */
	enki.Resources = resources.LoadFS()

	enki := enki.New("my_app")

	IntializeRoutes(enki.Routes)

	port := spt.FetchEnv("PORT", "3000")

	enki.ListenAndServe(port)
}
