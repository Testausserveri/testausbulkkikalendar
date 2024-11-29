package main

import (
	"log"
	"net/http"
	"testausserveri/testausbulkkikalendar/constants"
	"testausserveri/testausbulkkikalendar/gcal"
	"testausserveri/testausbulkkikalendar/handlers"
)

func main() {
	gcal.Init()

	// Parse all templates in the templates directory
	handlers.Init("./templates")

	// Create a file server handler
	fs := http.FileServer(http.Dir(constants.STATIC_CONTENT))

	// Strip the prefix if needed and serve files
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define the handlers
	http.Handle("/", handlers.AuthCheck(http.HandlerFunc(handlers.IndexHandler)))
	http.Handle("/query", handlers.AuthCheck(http.HandlerFunc(handlers.QueryHandler)))

	// Start the server
	log.Println("Server listening on :" + constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)
}
