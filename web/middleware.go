package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//old example code
// func writeToConsole(next http.Handler) http.Handler {
// 	return http.HandleFunc(func(w http.ResponseWriter, t *http.Request) {
// 		fmt.Println("get the page")
// 		next.ServeHTTP(w, r)
// 	}
// }

//use no surf to create a csrf handler token

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		//can this be set in another way?
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
