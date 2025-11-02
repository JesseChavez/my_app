package main

import (
	"my_app/app"
	"my_app/config"
	"log"
	"os"
)

func main() {
	initArgs := os.Args[1:]

	app.Log.Info("Init args: ", initArgs)

	if len(initArgs) == 0 {
		log.Fatal("Argument missing, requires 'db:migrate', 'db:create', 'db:drop', etc.")
	}

	runCommand(initArgs)
}

func runCommand(command []string) {
	app := config.NewApplication()

	app.ExecuteCommand(command)
}
