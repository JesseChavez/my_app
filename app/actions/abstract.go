package actions

import "github.com/JesseChavez/enki"

type Action struct {
	repo enki.Repository
	view enki.Renderer
}

func (axn *Action) InitAction(app *enki.Enki) {
	axn.repo = app.DB
	axn.view = app.Render
}

