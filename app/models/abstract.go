package models

import "github.com/JesseChavez/enki"

var repo enki.Repository

func InitModel(app *enki.Enki) {
	repo = app.DB
}
