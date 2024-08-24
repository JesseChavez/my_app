package config

import (
	"my_app/app/actions"

	"github.com/JesseChavez/enki"
)

func InitRoutes(app *enki.Enki, axn *actions.Action) *enki.Mux {
	mux := app.InitRouting()

	mux.Get("/", axn.HomeIndex)
	mux.Get("/today", axn.CurrentTime)

	return mux
}
