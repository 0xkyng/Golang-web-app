package main

// Serving real web pages.
import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/codekyng/go-web/pkg"
	"github.com/codekyng/go-web/pkg/config"
	"github.com/codekyng/go-web/render"
)

const portNumber = ":8080"

// In other for a function to respond to a request from a web browser;
// It has to handle two parameters;
// A response writer called (w http.ResponseWriter, r *http.Request)
// and a request r *http.Request

// main is the main application function
func main() {
	var app config.AppConfig

	templateCache, err := render.CraeteTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.Usedcache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)


	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err =srv.ListenAndServe()
	log.Fatal(err)
}