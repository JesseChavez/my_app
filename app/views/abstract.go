package views

import (
	"github.com/JesseChavez/enki"
)

var helper enki.IViewSupport

func InitView(app *enki.Enki) {
	helper = app.ViewSupport
}
