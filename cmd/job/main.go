package main

import (
	base "my_app/app"
	"my_app/app/models"
	"my_app/config"
	"log"
	"os"
)

func main() {
	initArgs := os.Args[1:]

	log.Println("Init args:", initArgs)

	if len(initArgs) == 0 {
		log.Fatal("Argument missing, requires 'start', 'stop', or 'work'")
	}

	command := initArgs[0]

	switch command {
	case "start":
		startJobApplication()
	case "stop":
		stopJobApplication()
	case "work":
		workJobApplication()
	default:
		log.Fatal("Invalid arguments, requires 'start', 'stop', or 'work'")
	}
}

func startJobApplication() {
	app := config.NewApplication()

	app.InitJobApplication()

	models.InitModel(&app)

	config.InitQueues(&app)

	base.InitUnits(&app)
	base.InitProps(&config.Props)

	app.StartAndProcess()
}

func stopJobApplication() {
}

func workJobApplication() {
	app := config.NewApplication()

	app.InitJobApplication()

	models.InitModel(&app)

	config.InitQueues(&app)

	base.InitUnits(&app)
	base.InitProps(&config.Props)

	app.StartAndWork()
}
