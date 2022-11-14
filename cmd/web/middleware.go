package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// WriteToConsole writes something to the console when a page is loaded
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf provides CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads our session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
