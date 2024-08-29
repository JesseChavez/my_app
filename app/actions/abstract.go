package actions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JesseChavez/enki"
)

type Action struct {
	repo enki.Repository
	tmpl enki.Renderer
	ssto enki.SessionStore
	// cjar enki.CookieJar
}

func (axn *Action) InitAction(app *enki.Enki) {
	axn.repo = app.DB
	axn.tmpl = app.Render
	axn.ssto = app.SessStore
}

// Helper to check auth.
func (axn *Action) requireAuth(w http.ResponseWriter, r *http.Request) error {	
	session := axn.ssto.GetSession(r)

	fmt.Println("session:", session)

	userID := session.Get("user_id")

	fmt.Println("user_id:", userID)

	// Set some session values.
	session.Set("user_id", time.Now().String())
	// Save it before we write to the response/return from the handler.

	// err := session.Save(r, w)
	err := session.Save(w)

	if err != nil {
		return err
	}

	return nil
}
