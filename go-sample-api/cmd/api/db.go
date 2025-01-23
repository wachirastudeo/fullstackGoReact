package main

import (
	"database/sql"
	"log"
)

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {

		return nil, err
	}
	return db, err

}
func (app *application) connectDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return connection, nil

}
