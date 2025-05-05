package actions

import (
	"log"
	"net/http"
)

func (axn *Action) SessionNew(w http.ResponseWriter, r *http.Request) {
	axn.view.RenderHTML(w, http.StatusOK, "sessions/new.tmpl", nil)
}

func (axn *Action) SessionCreate(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL params: %+v", r)

	axn.view.RenderHTML(w, http.StatusOK, "sessions/new.tmpl", nil)
}
