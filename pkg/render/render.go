package render

import (
	"bytes"
	"fmt"
	"github.com/JralstonDaz3d/bookings/pkg/config"
	"github.com/JralstonDaz3d/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template file
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		// create template cache
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	//use buffer
	buf := new(bytes.Buffer)

	//check for and add template data
	td = AddDefaultData(td)

	//execute
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all .tmpl files
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//look for layout file
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}
		if err != nil {
			return myCache, err
		}

		myCache[name] = ts
	}

	return myCache, nil

}

// First Way
var tc = make(map[string]*template.Template)

// RenderTemplatesCache renders a template file from cache
func RenderTemplatesCache(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//Check if we already have the template in cache
	_, inMap := tc[t]
	if !inMap {
		//need to create template
		log.Print("creating template and adding to cache")
		err = CreateTemplatesCache(t)
		if err != nil {
			log.Print(err)
		}
	} else {
		//we have template in cache
		log.Print("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
}

func CreateTemplatesCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err)
	}

	//add templates to cache (map)
	tc[t] = tmpl

	return nil
}
