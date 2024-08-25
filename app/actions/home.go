package actions

import (
	"net/http"
)

func (axn *Action) HomeIndex(w http.ResponseWriter, r *http.Request) {
	axn.tmpl.Render(w, r, "home/index.tmpl", nil)
}
