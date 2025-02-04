package main

import (
	_ "backend/cmd/api/docs" // เปลี่ยนชื่อไฟล์ให้ตรงกับ go.mod

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() http.Handler {
	//create router
	// Create a router mux
	mux := chi.NewRouter()

	// Add middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// Register Swagger
	mux.Get("/swagger/*", httpSwagger.WrapHandler)

	// Register the API routes under "/api/v1"
	mux.Route("/api/v1", func(r chi.Router) {
		// Public routes
		r.Get("/", app.Home)
		r.Get("/about", app.About)
		r.Get("/allmoviedemo", app.AllMoviedemo)

		// auth routes
		r.Post("/authenticate", app.authenticate)
		r.Get("/refresh", app.refreshToken)
		r.Get("/logout", app.logout)

		// movie routes
		r.Get("/movies", app.AllMovies)
		r.Get("/movies/{id}", app.GetMovie)
		r.Get("/genres", app.AllGenres)

		//group admin
		r.Route("/admin", func(r chi.Router) {
			// protected routes
			r.Use(app.authRequired)

			r.Get("/movies", app.MovieCatalog)
			r.Get("/movies/{id}", app.MovieForEdit)
			r.Post("/movies", app.InsertMovie)
			r.Put("/movies/{id}", app.UpdateMovie)
			r.Delete("/movies/{id}", app.DeleteMovie)

		})

	})

	return mux
}
