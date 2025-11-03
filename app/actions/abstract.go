package actions

import (
	"github.com/JesseChavez/enki"
)

type Action struct {
	repo enki.Repository
	view enki.IViewSupport
	smgr enki.ISessionManager
	nlog enki.ILogger
	// cjar enki.CookieJar
}

func (axn *Action) InitAction(app *enki.Enki) {
	axn.repo = app.DB
	axn.view = app.ViewSupport
	axn.smgr = app.SessionManager
	axn.nlog = app.Logger
}
