package views

import (
	"html/template"

	"github.com/JesseChavez/enki"
)

var helper enki.IHelper
var Help enki.IHelper

func InitView(app *enki.Enki) {
	helper = app.Helper
	Help = app.Helper
}

type ReactViewSpec struct {
	SpecID     string
	Production bool
	Prerender  bool
	Name       string
	Data       template.JS
}
