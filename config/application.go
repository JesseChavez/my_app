package config

import (
	"github.com/JesseChavez/enki"
	"github.com/JesseChavez/spt"
)

func RunApplication() {
	/* App context path, app runs in subfolder, default is root "/" */
	// enki.ContextPath = "/admin"

	enki := enki.New("my_app")

	IntializeRoutes(enki.Routes)

	port := spt.FetchEnv("PORT", "3000")

	enki.ListenAndServe(port)
}
