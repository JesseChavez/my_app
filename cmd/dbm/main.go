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
		log.Fatal("Argument missing, requires 'db:migrate', 'db:create', or 'db:drop'")
	}

	command := initArgs[0]

	switch command {
	case "db:migrate":
		runMigrate()
	case "db:create":
		runCreate()
	case "db:drop":
		runDrop()
	default:
		log.Fatal("Invalid arguments, requires 'db:migrate', 'db:create', or 'db:drop'")
	}
}

func runMigrate() {
	app := config.NewApplication()

	app.InitDbMigration()

	app.StartMigration()
}

func runCreate() {
}

func runDrop() {
}
