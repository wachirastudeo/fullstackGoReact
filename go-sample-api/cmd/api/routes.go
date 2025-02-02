package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	//create router
	r := chi.NewRouter()
	// middleware
	r.Use(middleware.Recoverer)
	r.Use(app.enableCORS)
	//register routes
	r.Get("/", app.Home)
	r.Get("/about", app.About)
	r.Get("/allmovie", app.AllMovie)
	r.Get("/allmoviedemo", app.AllMoviedemo)
	r.Post("/authenticate", app.authenticate)
	r.Get("/refresh", app.refreshToken)
	r.Get("/logout", app.logout)
	r.Get("/movie/{id}", app.GetMovie)

	//group admin
	r.Route("/admin", func(r chi.Router) {
		r.Get("/movie/{id}", app.MovieEdit)

	})

	return r
}
