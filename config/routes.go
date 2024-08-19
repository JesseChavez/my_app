package config

import (
	"my_app/app/actions"

	"github.com/JesseChavez/enki"
)

func IntializeRoutes(mux *enki.Mux) {
	mux.Get("/", actions.HomeIndex)
	mux.Get("/today", actions.CurrentTime)
}
