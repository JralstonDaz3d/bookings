package handlers

import (
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
	stringMap := make(map[string]string)
	stringMap["test"] = "Home"

	remoteIP := r.RemoteAddr
	lifeTime := m.App.Session.Lifetime.String()
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	m.App.Session.Put(r.Context(), "life_time", lifeTime)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "About"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	lifeTime := m.App.Session.GetString(r.Context(), "life_time")

	stringMap["remote_ip"] = remoteIP
	stringMap["life_time"] = lifeTime

	//send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Contact"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	lifeTime := m.App.Session.GetString(r.Context(), "life_time")

	stringMap["remote_ip"] = remoteIP
	stringMap["life_time"] = lifeTime

	//send data to the template
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Room is the room page handler
func (m *Repository) Room(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Room"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	lifeTime := m.App.Session.GetString(r.Context(), "life_time")

	stringMap["remote_ip"] = remoteIP
	stringMap["life_time"] = lifeTime

	//send data to the template
	render.RenderTemplate(w, "room.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Rooms is the rooms page handler
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Rooms"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	lifeTime := m.App.Session.GetString(r.Context(), "life_time")

	stringMap["remote_ip"] = remoteIP
	stringMap["life_time"] = lifeTime

	//send data to the template
	render.RenderTemplate(w, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Reservation"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	lifeTime := m.App.Session.GetString(r.Context(), "life_time")

	stringMap["remote_ip"] = remoteIP
	stringMap["life_time"] = lifeTime

	//send data to the template
	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
