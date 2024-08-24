package actions

import (
	"log"
	"net/http"
)

func (axn *Action) AccountIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Account Index!\n"))
}

func (axn *Action) AccountShow(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL params: %+v", r)
	w.Write([]byte("Account Show!\n"))
}
