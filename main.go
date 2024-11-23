package main

import (
	"log"
	"net/http"
	"testausserveri/testausbulkkikalendar/handlers"
)

const PORT = "8080"

func main() {
	// Parse all templates in the templates directory
	handlers.Init("templates/*.html")
	// Define the handlers

	http.HandleFunc("/", handlers.Index)

	// Start the server
	log.Println("Server listening on :" + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
