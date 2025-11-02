package config

import (
	"my_app/app/jobs"

	"github.com/JesseChavez/enki"
)

func InitQueues(app *enki.Enki) {
	queues := app.InitQueueing()

	queues.Register("default", app.TypeOf(jobs.AddUserJob{}))
}
