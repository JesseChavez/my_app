package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JesseChavez/spt"
)


func RunApplication() {
	routes := IntializeRoutes()

	port := spt.FetchEnv("PORT", "3000")

	log.Println("Web Application is starting")
	log.Println("*  Using Port:", port)
	log.Println("Use Ctrl-C to stop")

	webPort := fmt.Sprintf(":%v", port)

	log.Fatal(http.ListenAndServe(webPort, routes))
}
