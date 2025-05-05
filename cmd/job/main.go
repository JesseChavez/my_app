package main

import (
	"log"
	"os"
	"my_app/config"
)

func main() {
	initArgs := os.Args[1:]

	log.Println("Init args:", initArgs)

	if len(initArgs) == 0 {
		log.Fatal("Argument missing, requires 'start', or 'stop'")
	}

	command := initArgs[0]

	switch command {
	case "start":
		startJobApplication()
	case "stop":
		stopJobApplication()
	default:
		log.Fatal("Invalid arguments, requires 'start' or 'stop'")
	}
}

func startJobApplication() {
	app := config.NewApplication()

	app.InitJobApplication()

	app.StartAndProcess()
}

func stopJobApplication() {
}
