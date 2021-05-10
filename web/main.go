package main

import (
	"fmt"
	"log"
	"net/http"

	"_github.com/go-sql-driver/mysql/pkg/config"
	//"_github.com/go-sql-driver/mysql/pkg/config"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.createTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

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
