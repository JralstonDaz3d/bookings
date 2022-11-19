package main

import (
	"encoding/gob"
	"fmt"
	"github.com/JralstonDaz3d/bookings/internal/config"
	"github.com/JralstonDaz3d/bookings/internal/handlers"
	"github.com/JralstonDaz3d/bookings/internal/helpers"
	"github.com/JralstonDaz3d/bookings/internal/models"
	"github.com/JralstonDaz3d/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const PortNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// AddIntValues adds two int values
func AddIntValues(x, y int) int {
	return x + y
}

// main is the entry point to this app
func main() {

	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// set dev or prod mode (dev=false)
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// create session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	// serve the web page
	fmt.Println(fmt.Sprintf("Starting app on port %s", PortNumber))

	srv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
