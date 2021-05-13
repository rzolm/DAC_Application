package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	//build a mux (multiplexer)
	//mux := pat.New()

	//mux.Get("/", http.HandleFunc(handlers.Repo.advisor_login))
	//mux.Get("/", http.HandleFunc(handlers.Repo.advisor_home))

	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.advisor_login.page)
	mux.Get("/", handlers.Repo.advisor_home.page)
	mux.Get("/", handlers.Repo.patient_home.page)

	return mux
}
