package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

// create structs for application config
type application struct {
	Domain string
}

func main() {

	// set application config
	var app application
	app.Domain = "localhost"

	// read form command line args

	// connect database

	// start server
	// http.HandleFunc("/", Hello)
	// http.HandleFunc("/about", About)

	log.Printf("Starting server on port %d", port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", app.Domain, port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
