package actions

import (
	"log"
	"net/http"
)

func (axn *Action) AccountIndex(w http.ResponseWriter, r *http.Request) {

	axn.tmpl.Render(w, r, "accounts/index.tmpl", nil)
}

func (axn *Action) AccountShow(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL params: %+v", r)

	axn.tmpl.Render(w, r, "accounts/show.tmpl", nil)
}
