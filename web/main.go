package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/handlers"
	"github.com/rzolm/DAC_Application/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	//if set to true iyt will read from the cache not the disk
	app.UseCache = false

	//repo variable receives reference to the app config
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request))
	// http.HandleFunc("/", handlers.advisor_login)
	// http.HandleFunc("/", handlers.advisor_home)
	// http.HandleFunc("/", handlers.patient_home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
