package actions

import (
	"log"
	"net/http"
)

func (axn *Action) AccountIndex(w http.ResponseWriter, r *http.Request) {

	axn.view.RenderHTML(w, http.StatusOK, "accounts/index.tmpl", nil)
}

func (axn *Action) AccountShow(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL params: %+v", r)

	axn.view.RenderHTML(w, http.StatusOK, "accounts/show.tmpl", nil)
}
