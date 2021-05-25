package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/handlers"
	"github.com/rzolm/DAC_Application/pkg/render"
)

const portNumber = ":8080"

//outside of main so can be accessed by middleware
var app config.AppConfig

var session *scs.SessionManager

func main() {
	//change to true when in production
	app.InProduction = false

	//scs session 'MANAGER' type functions:
	//initialise a new session manager
	session = scs.New()
	//configure session lifetime
	session.Lifetime = 12 * time.Hour
	session.Cookie.Persist = true
	//sets the same site attribute on the session cookie (if empty will not be set)
	session.Cookie.SameSite = http.SameSiteLaxMode
	//points to the session field in config (AppConfig) - should be set to true in production
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	//if set to true it will read from the cache not the disk
	app.UseCache = false

	//repo variable receives reference to the app config
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request))
	// http.HandleFunc("/", handlers.advisor_login)
	// http.HandleFunc("/", handlers.advisor_home)
	// http.HandleFunc("/", handlers.patient_home)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
