package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	handlers "github.com/codekyng/go-web/pkg"
	"github.com/codekyng/go-web/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// using chi to route
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/", handlers.Repo.About)


	return mux
}