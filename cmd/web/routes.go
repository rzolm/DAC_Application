package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	//a mux is a 'http request multiplexer'

	//create a new router
	mux := chi.NewRouter()

	//import the middleware package from chi
	mux.Use(middleware.Recoverer)

	//example of how to load the function from the middleware file in web
	//mux.Use(writeToConsole)

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//get the pages via the handlers repo
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/", handlers.Repo.Advisor)
	mux.Get("/", handlers.Repo.Patient)

	//agent login form
	mux.Post("/", handlers.Repo.Index)

	return mux
}
