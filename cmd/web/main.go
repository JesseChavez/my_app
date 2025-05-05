package main

import (
	"log"
	"net/http"
	"os"
	"my_app/app/actions"
	"my_app/app/models"
	"my_app/app/views"
	"my_app/config"
	"time"

	"github.com/JesseChavez/enki"
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
		startWebApplication()
	case "stop":
		stopWebApplication()
	default:
		log.Fatal("Invalid arguments, requires 'start' or 'stop'")
	}
}

func startWebApplication() {
	app := config.NewApplication()

	axn := actions.Action{}

	mux := config.InitRoutes(&app, &axn)

	app.InitWebApplication(mux)

	axn.InitAction(&app)

	views.InitView(&app)

	models.InitModel(&app)

	app.ListenAndServe()
}

func stopWebApplication() {
	_ = config.NewApplication()

	host := "http://localhost"

	baseUrl := host + ":" + enki.WebPort

	stopUrl := baseUrl + "/shutdown"

	http.DefaultClient.Timeout = 10 * time.Second

	resp, err := http.Get(stopUrl)

	if err != nil {
		log.Println("Unable to stop sever running on:", baseUrl)
		log.Println(err.Error())
	}

	log.Println("Server is stopping with response")
	log.Println(resp)
}
