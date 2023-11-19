package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"net/http"
	"path/filepath"

	"github.com/Adityabhaumik/project1/pkg/config"
	"github.com/Adityabhaumik/project1/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefault(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	fmt.Println("render template accessed")
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplatecache()
	}
	// tc := app.TemplateCache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("cant get template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefault(td)
	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplatecache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all files named .pages.tmpl from templates folder

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//loop through loopes

	for _, page := range pages {
		//page is the global path
		name := filepath.Base(page)                    // name is the name of the file only
		ts, err := template.New(name).ParseFiles(page) //ts -> template set
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}

	return myCache, nil
}
