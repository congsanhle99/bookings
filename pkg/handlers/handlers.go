package handlers

import (
	"fmt"
	"net/http"

	"github.com/lcsanh/bookings/pkg/config"
	"github.com/lcsanh/bookings/pkg/models"
	"github.com/lcsanh/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println("remoteIP", remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again!"

	// remoteIP will be empty string if is nothing in the session name 'remote_ip'
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	// send data to the template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation render the makes the Reservation page and display form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.html", &models.TemplateData{})
}

// Generals render the Room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.html", &models.TemplateData{})
}

// Majors render the Room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.html", &models.TemplateData{})
}

// Availability render the search Availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.html", &models.TemplateData{})
}

// contact render the search contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.html", &models.TemplateData{})
}
