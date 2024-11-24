package main

import (
	"log"
	"net/http"
	"testausserveri/testausbulkkikalendar/constants"
	"testausserveri/testausbulkkikalendar/handlers"
	"testausserveri/testausbulkkikalendar/oauth"
)

func main() {
	oauth.Init()

	// Parse all templates in the templates directory
	handlers.Init("templates/*.html")
	// Define the handlers
	http.HandleFunc("/", handlers.Index)

	// Start the server
	log.Println("Server listening on :" + constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)
}
