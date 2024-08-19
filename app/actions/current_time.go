package actions

import (
	"net/http"
	"time"
)

func CurrentTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	res := now.String()

	w.Write([]byte(res + "\n"))
}
