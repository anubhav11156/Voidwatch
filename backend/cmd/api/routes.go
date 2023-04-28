package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)
	mux.Post("/authenticate", app.authenticate)

	mux.Get("/getAllMovies", app.AllMovies)

	//	will apply a middleware to any routes starting with /admin
	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired) // if this is passed then only go down else return
		mux.Get("/movies", app.MovieCatalog)
	})

	return mux
}
