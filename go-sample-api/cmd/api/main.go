package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = 8080

// create structs for application config
type application struct {
	DSN          string //database connection string
	Domain       string //domain name
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	APIKey       string
}

func main() {

	// set application config
	var app application
	app.Domain = "localhost"

	// read form command line args

	// set database
	dsn := flag.String("dsn", "host=localhost user=postgres dbname=gosampledb password=123456 port=5432 sslmode=disable timezone=UTC connect_timeout=5", "database connection string")
	flag.Parse()
	app.DSN = *dsn
	// connect to database
	conn, err := app.connectDB()

	if err != nil {
		log.Fatal(err)
	}
	// connect by interface
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	defer app.DB.Connection().Close()

	// start server
	// http.HandleFunc("/", Hello)
	// http.HandleFunc("/about", About)
	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}
	log.Printf("Starting server on port %d", port)

	err = http.ListenAndServe(fmt.Sprintf("%s:%d", app.Domain, port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
