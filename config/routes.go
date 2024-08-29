package config

import (
	"my_app/app/actions"

	"github.com/JesseChavez/enki"
)

func InitRoutes(app *enki.Enki, axn *actions.Action) *enki.Mux {
	mux := app.InitRouting()

	mux.Get("/", axn.Dashboard)
	mux.Get("/today", axn.CurrentTime)

	// accounts
	mux.Get("/accounts", axn.AccountIndex)
	mux.Get("/accounts/{id}", axn.AccountShow)

	// Subrouters:
	mux.Route("/session", func(m enki.Router) {
		m.Post("/", axn.SessionNew)
		m.Get("/new", axn.SessionCreate)
	})

	return mux
}
