package config

import (
	"my_app/app/actions"

	"github.com/JesseChavez/enki"
)

func InitRoutes(app *enki.Enki, axn *actions.Action) *enki.Mux {
	mux := app.InitRouting()

	mux.Get("/", axn.HomeIndex)
	mux.Get("/today", axn.CurrentTime)
	mux.Get("/today", axn.CurrentTime)

	// accounts
	mux.Get("/accounts", axn.AccountIndex)
	mux.Get("/accounts/{id}", axn.AccountShow)

	return mux
}
