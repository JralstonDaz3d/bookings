package main

import (
	"github.com/JralstonDaz3d/bookings/internal/config"
	"github.com/JralstonDaz3d/bookings/internal/handlers"
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

	// Get requests
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/reservation", http.HandlerFunc(handlers.Repo.Reservation))
	mux.Get("/room", http.HandlerFunc(handlers.Repo.Room))
	mux.Get("/rooms", http.HandlerFunc(handlers.Repo.Rooms))
	mux.Get("/privacy-policy", http.HandlerFunc(handlers.Repo.PrivacyPolicy))
	mux.Get("/terms-of-use", http.HandlerFunc(handlers.Repo.TermsOfUse))

	// Post requests
	mux.Post("/contact-form", http.HandlerFunc(handlers.Repo.ContactPost))
	mux.Post("/reservation-form", http.HandlerFunc(handlers.Repo.ReservationPost))
	mux.Post("/rooms-form", http.HandlerFunc(handlers.Repo.RoomsSearch))
	mux.Post("/rooms", http.HandlerFunc(handlers.Repo.RoomsPost))

	// Post AJAX
	mux.Post("/available", http.HandlerFunc(handlers.Repo.AvailabilityJSON))

	// Tell file server where to get static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
