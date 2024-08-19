package config

import (
	"my_app/app/actions"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func IntializeRoutes() chi.Router {
	r := chi.NewRouter()

	// adding to log request details
	r.Use(middleware.Logger)

	// adding to recover panics gracefully and log stack trace
	r.Use(middleware.Recoverer)

	// adding server status check
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", actions.HomeIndex)

	return r
}
