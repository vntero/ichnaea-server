package main

import (
	"ichnaea-server/routes"
	"log"
	"net/http"
)

func main() {
	// initialize routes
	routes.GetHello()

	log.Println("Listening on port :5005")
	http.ListenAndServe(":5005", nil)
}