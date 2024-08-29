package actions

import (
	"net/http"
)

func (axn *Action) Dashboard(w http.ResponseWriter, r *http.Request) {
	err := axn.requireAuth(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	axn.tmpl.Render(w, r, "dashboard/index.tmpl", nil)
}
