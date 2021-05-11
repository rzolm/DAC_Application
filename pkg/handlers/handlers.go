package handlers

import (
	"fmt"
	"net/http"

	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/models"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//this sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//index page handler - method with access to repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor login")
	RenderTemplate(w, "advisor_login.page.gohtml", &models.TemplateData{})
}

//advisor home page - method with access to repository
func (m *Repository) Advisor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor home page")
	RenderTemplate(w, "advisor_home.gohtml", &models.TemplateData)
}

//patient home page - method with access to repository
func (m *Repository) Patient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("patient home page")
	RenderTemplate(w, "patient_home.gohtml", &models.TemplateData)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData)) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return
	}
}
