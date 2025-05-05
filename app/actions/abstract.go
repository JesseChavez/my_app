package actions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JesseChavez/enki"
)

type Action struct {
	repo enki.Repository
	view enki.IRenderer
	ssto enki.ISessionStore
	help enki.IHelper
	nlog enki.ILogger
	// cjar enki.ICookieJar
}

func (axn *Action) InitAction(app *enki.Enki) {
	axn.repo = app.DB
	axn.view = app.Renderer
	axn.ssto = app.SessionStore
	axn.help = app.Helper
	axn.nlog = app.Logger
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
