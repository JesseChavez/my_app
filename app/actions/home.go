package actions

import (
	"net/http"
)

func (axn *Action) HomeIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
