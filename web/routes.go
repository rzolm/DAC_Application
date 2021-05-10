package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/rzolm/DAC_Application/config"
)

func routes(app *config.AppConfig) http.Handler {
	//build a mux
	mux := pat.New()

	mux.Get("/", http.HandleFunc(handlers.repo.advisor_login))
	mux.Get("/", http.HandleFunc(handlers.repo.advisor_home))

	return mux
}
