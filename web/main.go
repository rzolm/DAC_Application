package main

import (
	"fmt"
	"net/http"

	"_github.com/go-sql-driver/mysql/pkg/config"
	//"_github.com/go-sql-driver/mysql/pkg/config"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.createTemplateCache()
	if err err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request))
	http.HandleFunc("/", handlers.advisor_login)
	http.HandleFunc("/", handlers.advisor_home)
	http.HandleFunc("/", handlers.patient_home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(":8080", nil)
}
