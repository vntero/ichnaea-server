package routes

import (
	"ichnaea-server/handlers"

	"log"
	"net/http"
)

func GetHello() {
	http.HandleFunc("/", handlers.ListHello)
	log.Println("/ is now available")
}