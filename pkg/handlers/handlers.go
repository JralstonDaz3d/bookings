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
	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["test"] = "Home"
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["life_time"] = m.App.Session.GetString(r.Context(), "life_time")

	stringMap["name"] = m.App.Session.GetString(r.Context(), "name")
	stringMap["email"] = m.App.Session.GetString(r.Context(), "email")
	stringMap["phone"] = m.App.Session.GetString(r.Context(), "phone")

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// ContactPost is the home page post handler - shows how to handle and process form post data and display it on a template page.
func (m *Repository) ContactPost(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	message := r.Form.Get("message")

	// Write plain text out to page
	//w.Write([]byte("Posted to Contact"))
	//w.Write([]byte(fmt.Sprintf("%s with email:%s and phone:%s say: %s", name, email, phone, message)))

	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["name"] = name
	stringMap["email"] = email
	stringMap["phone"] = phone
	stringMap["message"] = message

	// also add the user info to the session
	m.App.Session.Put(r.Context(), "name", name)
	m.App.Session.Put(r.Context(), "email", email)
	m.App.Session.Put(r.Context(), "phone", phone)

	// render template
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{
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
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
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
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{
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
	render.RenderTemplate(w, r, "room.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Rooms is the rooms page handler
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["test"] = "Reservation"

	//populate from session data (if exists)
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["life_time"] = m.App.Session.GetString(r.Context(), "life_time")

	stringMap["name"] = m.App.Session.GetString(r.Context(), "name")
	stringMap["email"] = m.App.Session.GetString(r.Context(), "email")
	stringMap["phone"] = m.App.Session.GetString(r.Context(), "phone")
	stringMap["message"] = m.App.Session.GetString(r.Context(), "message")
	stringMap["rooms"] = m.App.Session.GetString(r.Context(), "rooms")
	stringMap["guests"] = m.App.Session.GetString(r.Context(), "guests")
	stringMap["roomtype"] = m.App.Session.GetString(r.Context(), "roomtype")
	stringMap["zip"] = m.App.Session.GetString(r.Context(), "zip")
	stringMap["city"] = m.App.Session.GetString(r.Context(), "city")
	stringMap["country"] = m.App.Session.GetString(r.Context(), "country")
	stringMap["start"] = m.App.Session.GetString(r.Context(), "start")
	stringMap["end"] = m.App.Session.GetString(r.Context(), "end")

	stringMap["search"] = m.App.Session.GetString(r.Context(), "search")

	//send data to the template
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// RoomsPost is the rooms page post handler
func (m *Repository) RoomsPost(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	message := r.Form.Get("message")
	rooms := r.Form.Get("rooms")
	guests := r.Form.Get("guests")
	roomtype := r.Form.Get("roomtype")
	zip := r.Form.Get("zip")
	city := r.Form.Get("city")
	country := r.Form.Get("country")
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	// Write plain text out to page
	//w.Write([]byte("Posted to Contact"))
	//w.Write([]byte(fmt.Sprintf("%s with email:%s and phone:%s say: %s", name, email, phone, message)))

	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["name"] = name
	stringMap["email"] = email
	stringMap["phone"] = phone
	stringMap["message"] = message
	stringMap["rooms"] = rooms
	stringMap["guests"] = guests
	stringMap["roomtype"] = roomtype
	stringMap["zip"] = zip
	stringMap["city"] = city
	stringMap["country"] = country
	stringMap["start"] = start
	stringMap["end"] = end

	stringMap["ok"] = "true"
	stringMap["msg"] = "Successfully sent form data!"

	// add the user info to the session
	m.App.Session.Put(r.Context(), "name", name)
	m.App.Session.Put(r.Context(), "email", email)
	m.App.Session.Put(r.Context(), "phone", phone)
	m.App.Session.Put(r.Context(), "message", message)

	m.App.Session.Put(r.Context(), "rooms", rooms)
	m.App.Session.Put(r.Context(), "guests", guests)
	m.App.Session.Put(r.Context(), "roomtype", roomtype)

	m.App.Session.Put(r.Context(), "zip", zip)
	m.App.Session.Put(r.Context(), "city", city)
	m.App.Session.Put(r.Context(), "country", country)
	m.App.Session.Put(r.Context(), "start", start)
	m.App.Session.Put(r.Context(), "end", end)

	// render template
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// RoomsSearch is the rooms page search post handler
func (m *Repository) RoomsSearch(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	search := r.Form.Get("search")
	/*
		rooms := r.Form.Get("rooms")
		guests := r.Form.Get("guests")
		roomtype := r.Form.Get("roomtype")
		zip := r.Form.Get("zip")
		city := r.Form.Get("city")
		country := r.Form.Get("country")
		start := r.Form.Get("start")
		end := r.Form.Get("end")
	*/
	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["search"] = search
	/*stringMap["rooms"] = rooms
	stringMap["guests"] = guests
	stringMap["roomtype"] = roomtype
	stringMap["zip"] = zip
	stringMap["city"] = city
	stringMap["country"] = country
	stringMap["start"] = start
	stringMap["end"] = end*/

	stringMap["ok"] = "true"
	stringMap["msg"] = "Successfully sent form data!"

	// add the user info to the session
	m.App.Session.Put(r.Context(), "search", search)
	/*m.App.Session.Put(r.Context(), "rooms", rooms)
	m.App.Session.Put(r.Context(), "guests", guests)
	m.App.Session.Put(r.Context(), "roomtype", roomtype)

	m.App.Session.Put(r.Context(), "zip", zip)
	m.App.Session.Put(r.Context(), "city", city)
	m.App.Session.Put(r.Context(), "country", country)
	m.App.Session.Put(r.Context(), "start", start)
	m.App.Session.Put(r.Context(), "end", end)*/

	// Write plain text out to page
	//w.Write([]byte("Posted to Contact"))
	//w.Write([]byte(fmt.Sprintf("Searching: %s with rooms::%s and guests:%s roomtype: %s", search, rooms, guests, roomtype)))

	// render template
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["test"] = "Reservation"

	//populate from session data (if exists)
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["life_time"] = m.App.Session.GetString(r.Context(), "life_time")

	stringMap["name"] = m.App.Session.GetString(r.Context(), "name")
	stringMap["email"] = m.App.Session.GetString(r.Context(), "email")
	stringMap["phone"] = m.App.Session.GetString(r.Context(), "phone")
	stringMap["message"] = m.App.Session.GetString(r.Context(), "message")
	stringMap["rooms"] = m.App.Session.GetString(r.Context(), "rooms")
	stringMap["guests"] = m.App.Session.GetString(r.Context(), "guests")
	stringMap["roomtype"] = m.App.Session.GetString(r.Context(), "roomtype")
	stringMap["zip"] = m.App.Session.GetString(r.Context(), "zip")
	stringMap["city"] = m.App.Session.GetString(r.Context(), "city")
	stringMap["country"] = m.App.Session.GetString(r.Context(), "country")
	stringMap["start"] = m.App.Session.GetString(r.Context(), "start")
	stringMap["end"] = m.App.Session.GetString(r.Context(), "end")

	//send data to the template
	render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// ReservationPost is the reservation page post handler - shows how to handle and process form post data
func (m *Repository) ReservationPost(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	message := r.Form.Get("message")
	rooms := r.Form.Get("rooms")
	guests := r.Form.Get("guests")
	roomtype := r.Form.Get("roomtype")
	zip := r.Form.Get("zip")
	city := r.Form.Get("city")
	country := r.Form.Get("country")
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	// Write plain text out to page
	//w.Write([]byte("Posted to Contact"))
	//w.Write([]byte(fmt.Sprintf("%s with email:%s and phone:%s say: %s", name, email, phone, message)))

	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["name"] = name
	stringMap["email"] = email
	stringMap["phone"] = phone
	stringMap["message"] = message
	stringMap["rooms"] = rooms
	stringMap["guests"] = guests
	stringMap["roomtype"] = roomtype
	stringMap["zip"] = zip
	stringMap["city"] = city
	stringMap["country"] = country
	stringMap["start"] = start
	stringMap["end"] = end

	stringMap["ok"] = "true"
	stringMap["msg"] = "Successfully sent form data!"

	// add the user info to the session
	m.App.Session.Put(r.Context(), "name", name)
	m.App.Session.Put(r.Context(), "email", email)
	m.App.Session.Put(r.Context(), "phone", phone)
	m.App.Session.Put(r.Context(), "message", message)

	m.App.Session.Put(r.Context(), "rooms", rooms)
	m.App.Session.Put(r.Context(), "guests", guests)
	m.App.Session.Put(r.Context(), "roomtype", roomtype)

	m.App.Session.Put(r.Context(), "zip", zip)
	m.App.Session.Put(r.Context(), "city", city)
	m.App.Session.Put(r.Context(), "country", country)
	m.App.Session.Put(r.Context(), "start", start)
	m.App.Session.Put(r.Context(), "end", end)

	// render template
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// PrivacyPolicy is the about page handler
func (m *Repository) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Privacy Policy"

	//send data to the template
	render.RenderTemplate(w, r, "privacy-policy.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// TermsOfUse is the about page handler
func (m *Repository) TermsOfUse(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Terms of Use"

	//send data to the template
	render.RenderTemplate(w, r, "terms-of-use.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
