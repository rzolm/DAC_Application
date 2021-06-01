package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/models"
)

//map of functions to be used in the template
//FuncMap is a type of map defining the mapping of names to functions
var functions = template.FuncMap{}

var app *config.AppConfig

//this sets the  config fopr the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//renders templates using html/template
func Template(w http.ResponseWriter, gohtml string, td *models.TemplateData) {
	//get the template cache from the app config
	var tc map[string]*template.Template

	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get the template cahe from the appconfig
	//tc := app.TemplateCache

	//if the template exists 't' will return a value
	t, ok := tc[gohtml]
	if !ok {
		fmt.Println("error")
	}

	//this buffer will hold bytes to test the template
	buf := new(bytes.Buffer)

	//take the template execute it and store in the above buf
	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmp)
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template", err)
	// 	return
	//}
}

//create a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	//parse all the templates and store them in the map myCache
	myCache := map[string]*template.Template{}

	//search the templates folder for templates that have '.page.gohtml' in the name
	pages, err := filepath.Glob("/templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	//for loop to run through files in pkg/templates
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("page is currently", page)

		//
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// matches, err := filepath.Glob("./templates/*layout.gohtml")
		// if err != nil {
		// 	return myCache, err
		// }

		// if len(matches) > 0 {
		// 	ts, err := ts.ParseGlob("./templates/*layout.gohtml")
		// 	if err != nil {
		// 		return myCache, err
		// 	}
		// }

		myCache[name] = ts
	}
	return myCache, nil
}
