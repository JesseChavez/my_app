package actions

import "github.com/JesseChavez/enki"

type Action struct {
	repo enki.Repository
	tmpl enki.Renderer
}

func (axn *Action) InitAction(app *enki.Enki) {
	axn.repo = app.DB
	axn.tmpl = app.Render
}

