package actions

import (
	"log"
	"net/http"

	"github.com/JesseChavez/enki"
)

func (axn *Action) AccountIndex(w http.ResponseWriter, r *http.Request) {
	view := enki.ActionView{
		Template: "accounts/index.tmpl",
		Name: "AccountShow",
		Debug: true,
		Data: nil,
	}

	axn.view.Render(w, http.StatusOK, &view)
}

func (axn *Action) AccountShow(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL params: %+v", r)

	view := enki.ActionView{
		Template: "accounts/show.tmpl",
		Name: "AccountShow",
		Debug: true,
		Data: nil,
	}

	axn.view.Render(w, http.StatusOK, &view)
}
