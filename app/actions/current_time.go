package actions

import (
	"fmt"
	"net/http"
	"time"
)

func (axn *Action) CurrentTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	res := now.String()

	delim := "========================================================="

	output := fmt.Sprintf("%s\n%s\n%s\n", delim, res, delim)

	w.Write([]byte(output))
}
