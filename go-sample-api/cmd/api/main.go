package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

// create structs for application config
type application struct {
	DSN    string //database connection string
	Domain string //domain name
	DB     *sql.DB
}

func main() {

	// set application config
	var app application
	app.Domain = "localhost"

	// read form command line args

	// connect database
	dsn := flag.String("dsn", "host=localhost user=postgres dbname=gosampledb password=123456 port=5432 sslmode=disable timezone=UTC connect_timeout=5", "database connection string")
	flag.Parse()
	app.DSN = *dsn

	conn, err := app.connectDB()

	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn
	defer conn.Close()
	// start server
	// http.HandleFunc("/", Hello)
	// http.HandleFunc("/about", About)

	log.Printf("Starting server on port %d", port)

	err = http.ListenAndServe(fmt.Sprintf("%s:%d", app.Domain, port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
