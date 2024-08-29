package actions

import (
	"log"
	"net/http"
)

func (axn *Action) SessionNew(w http.ResponseWriter, r *http.Request) {
	axn.tmpl.Render(w, r, "sessions/new.tmpl", nil)
}

func (axn *Action) SessionCreate(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL params: %+v", r)

	axn.tmpl.Render(w, r, "sessions/new.tmpl", nil)
}
