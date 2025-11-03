package actions

import (
	"net/http"

	"github.com/JesseChavez/enki"
)

func (axn *Action) Dashboard(w http.ResponseWriter, r *http.Request) {
	view := enki.ActionView{
		Template: "dashboard/index.tmpl",
		Name: "AccountShow",
		Debug: true,
		Data: nil,
	}

	axn.view.Render(w, http.StatusOK, &view)
}
