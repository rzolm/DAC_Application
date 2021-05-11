//setting up a template cache

package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rzolm/DAC_Application/config"
	"github.com/rzolm/DAC_Application/pkg/handlers"
)

var functions = template.FuncMap{}

//this sets the  config fopr the template package
func newTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *handlers.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		//get the template cache from the app config
		tc = app.Templatecache()
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get the template cahe from the appconfig
	//tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get tempplate from cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
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

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("/templates/*.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("page is currently", page)
		ts, err := template.New(name).Funcs(functions).parseFiles(pages)
		if err != nil {
			return myCache, err
		}
	}
}
