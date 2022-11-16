package render

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"net/http"
	"text/template"

	"github.com/codekyng/go-web/pkg/config"
	"github.com/codekyng/go-web/pkg/models"
)


var functions = template.FuncMap{

}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}


func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	 return td
}
// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template
	if app.Usedcache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CraeteTemplateCache()
	}
	// get the template cache from the app config

	templateCache = app.TemplateCache


	// get requested from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_= template.Execute(buf, td)

	
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)

	}

}

// create a template cache
func CraeteTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil

}

