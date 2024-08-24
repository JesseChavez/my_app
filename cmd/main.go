package main

import (
	"log"
	"os"
	"my_app/app/actions"
	"my_app/config"
)

func main() {
	initArgs := os.Args[1:]

	log.Println("Init args:", initArgs)

	if len(initArgs) == 0 {
		log.Fatal("Argument missing, requires 'web:start', 'job:start', or 'db:migrate'")
	}

	command := initArgs[0]

	switch command {
	case "web:start":
		runWebApplication()
	case "job:start":
		runJobApplication()
	case "db:migrate":
		runMigration()
	default:
		log.Fatal("Invalid arguments, requires 'web:start', 'job:star', or 'db:migrate'")
	}
}

func runWebApplication() {
	app := config.NewApplication()

	axn := actions.Action{}

	mux := config.InitRoutes(&app, &axn)

	app.InitWebApplication(mux)

	axn.InitAction(&app)

	app.ListenAndServe()
}

func runJobApplication() {
	app := config.NewApplication()

	app.InitJobApplication()

	app.StartAndProcess()
}

func runMigration () {
	app := config.NewApplication()

	app.InitDbMigration()

	app.StartMigration()
}
