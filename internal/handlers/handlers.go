package handlers

import (
	"encoding/json"
	"github.com/JralstonDaz3d/bookings/internal/config"
	"github.com/JralstonDaz3d/bookings/internal/forms"
	"github.com/JralstonDaz3d/bookings/internal/models"
	"github.com/JralstonDaz3d/bookings/internal/render"
	"log"
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

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["test"] = "Home"

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		data["reservation"] = emptyReservation
		//return
	}
	data["reservation"] = reservation

	stringMap["name"] = reservation.Name
	stringMap["email"] = reservation.Email
	stringMap["phone"] = reservation.Phone

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Form:      forms.New(nil),
		Data:      data,
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

	data := make(map[string]interface{})
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		return
	}
	data["reservation"] = reservation

	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)

	//send data to the template
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Form:      forms.New(nil),
		Data:      data,
	})
}

// RoomsPost is the rooms page post handler
func (m *Repository) RoomsPost(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	var emptyReservation models.Reservation
	data := make(map[string]interface{})

	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")

		data["reservation"] = emptyReservation
		//return
	}
	data["reservation"] = reservation

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	// Add form values
	reservation.Rooms = r.Form.Get("rooms")
	reservation.Guests = r.Form.Get("guests")
	reservation.Roomtype = r.Form.Get("roomtype")
	reservation.Search = r.Form.Get("search")

	data["reservation"] = reservation

	form := forms.New(r.PostForm)

	// form is valid so add the user info to the session
	m.App.Session.Put(r.Context(), "reservation", reservation)

	// render template
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Form:      form,
		Data:      data,
	})
}

// RoomsSearch is the rooms page search post handler
func (m *Repository) RoomsSearch(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	search := r.Form.Get("search")

	// Put the values in a stringMap and insert into template
	stringMap := make(map[string]string)
	stringMap["search"] = search
	stringMap["ok"] = "true"
	stringMap["msg"] = "Successfully sent form data!"

	// add the user info to the session
	m.App.Session.Put(r.Context(), "search", search)

	// render template
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	form := forms.New(r.PostForm)

	if m.App.Session.Get(r.Context(), "reservation") != "" {

		reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
		if !ok {
			log.Println("cannot get item from session")
			//return
			var emptyReservation models.Reservation
			data["reservation"] = emptyReservation
		}
		data["reservation"] = reservation
	}

	//send data to the template
	render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{
		Form: form,
		Data: data,
	})
}

// ReservationPost is the reservation page post handler - shows how to handle and process form post data
func (m *Repository) ReservationPost(w http.ResponseWriter, r *http.Request) {
	//stringMap := make(map[string]string)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Phone:    r.Form.Get("phone"),
		Message:  r.Form.Get("message"),
		Rooms:    r.Form.Get("rooms"),
		Guests:   r.Form.Get("guests"),
		Roomtype: r.Form.Get("roomtype"),
		Zip:      r.Form.Get("zip"),
		City:     r.Form.Get("city"),
		Country:  r.Form.Get("country"),
		Start:    r.Form.Get("start"),
		End:      r.Form.Get("end"),
	}

	form := forms.New(r.PostForm)

	// form validation + GoValidator
	form.Required("name", "email", "phone", "rooms", "guests", "roomtype", "zip", "city", "country", "start", "end")
	form.MinLength("name", 3, r)
	form.MinLength("phone", 10, r)
	form.IsEmail("email")
	form.IsInt("rooms", "guests")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	if !form.Valid() {
		//send data to the template
		render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// form is valid

	// add the user info to the session
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/rooms", http.StatusSeeOther)
	/*
		m.App.Session.Put(r.Context(), "name", reservation.Name)
		m.App.Session.Put(r.Context(), "email", reservation.Email)
		m.App.Session.Put(r.Context(), "phone", reservation.Phone)
		m.App.Session.Put(r.Context(), "message", reservation.Message)
		m.App.Session.Put(r.Context(), "rooms", reservation.Rooms)
		m.App.Session.Put(r.Context(), "guests", reservation.Guests)
		m.App.Session.Put(r.Context(), "roomtype", reservation.Roomtype)
		m.App.Session.Put(r.Context(), "zip", reservation.Zip)
		m.App.Session.Put(r.Context(), "city", reservation.City)
		m.App.Session.Put(r.Context(), "country", reservation.Country)
		m.App.Session.Put(r.Context(), "start", reservation.Start)
		m.App.Session.Put(r.Context(), "end", reservation.End)

		// render template
		render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
			Form:      form,
			Data:      data,
		})
	*/
}

// PrivacyPolicy is the PP page handler
func (m *Repository) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Privacy Policy"

	//send data to the template
	render.RenderTemplate(w, r, "privacy-policy.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// TermsOfUse is the TOU page handler
func (m *Repository) TermsOfUse(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Terms of Use"

	//send data to the template
	render.RenderTemplate(w, r, "terms-of-use.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

// AvailabilityJSON responds with JSON
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	resp := jsonResponse{
		Ok:      true,
		Message: "Available",
		Start:   start,
		End:     end,
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	//log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// AvailabilityPostJSON responds with JSON
func (m *Repository) AvailabilityPostJSON(w http.ResponseWriter, r *http.Request) {

}
