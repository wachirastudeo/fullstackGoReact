package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprint(w, "Hello World from ", app.Domain)
	// json data
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go went Gone",
		Version: "1.0.0",
	}
	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
func (app *application) About(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "About")
}

func (app *application) AllMovie(w http.ResponseWriter, r *http.Request) {

	rd, _ := time.Parse("2006-01-02", "2022-01-01")
	var movie []models.Movie

	handler := models.Movie{
		ID:          1,
		Title:       "handler",
		ReleaseDate: rd,
		Runtime:     200,
		MPAARating:  "PG-13",
		Description: "some description",
		Image:       "some image", Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	movie = append(movie, handler)
	rotla := models.Movie{
		ID:          2,
		Title:       "rotla",
		ReleaseDate: rd,
		Runtime:     100,
		MPAARating:  "PG-13",
		Description: "some description rotla",
		Image:       "some image", Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	movie = append(movie, rotla)

	out, err := json.Marshal(movie)
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
