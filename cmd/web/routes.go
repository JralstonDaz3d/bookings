package main

import (
	"github.com/JralstonDaz3d/bookings/pkg/config"
	"github.com/JralstonDaz3d/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	// Router using using pat
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// Router using chi
	mux := chi.NewRouter()

	// some chi middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//some custom middleware
	mux.Use(WriteToConsole)

	// Routes to our pages (update this when a new route or template is added!)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/reservation", http.HandlerFunc(handlers.Repo.Reservation))
	mux.Get("/room", http.HandlerFunc(handlers.Repo.Room))
	mux.Get("/rooms", http.HandlerFunc(handlers.Repo.Rooms))

	// Tell file server where to get static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
