package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	//create router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	//register routes
	r.Get("/", app.Home)
	r.Get("/about", app.About)

	return r
}
