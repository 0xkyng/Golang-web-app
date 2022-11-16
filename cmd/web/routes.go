package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	handlers "github.com/codekyng/go-web/pkg"
	"github.com/codekyng/go-web/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}