package app

import (
	"github.com/JesseChavez/enki"
)

var DB  enki.Repository
var Log enki.ILogger
var Job enki.IJobSupport

func InitUnits(app *enki.Enki) {
	DB  = app.DB
	Log = app.Logger
	Job = app.JobSupport
}
