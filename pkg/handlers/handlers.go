package handlers

import (
	"errors"
	"fmt"
	"github.com/JralstonDaz3d/bookings/pkg/config"
	"github.com/JralstonDaz3d/bookings/pkg/models"
	"github.com/JralstonDaz3d/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets reposity for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	lifeTime := m.App.Session.Lifetime.String()
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	m.App.Session.Put(r.Context(), "life_time", lifeTime)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	lifeTime := m.App.Session.GetString(r.Context(), "life_time")

	stringMap["remote_ip"] = remoteIP
	stringMap["life_time"] = lifeTime

	//send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Divide is a page to test errors
func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 0.0

	f, err := DivideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "bad")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func DivideValues(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("cannpot divide by zero!")
		return 0, err
	}
	result := x / y
	return result, nil
}
