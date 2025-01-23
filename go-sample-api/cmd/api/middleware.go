package main

import (
	"net/http"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS ,PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type Accept X-CSRF-Token Authorization")

			return
		} else {
			next.ServeHTTP(w, r)

		}
	})
}
