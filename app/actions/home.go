package actions

import (
	"net/http"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
