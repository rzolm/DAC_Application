package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/models"
)

//holds data sent from handlers to data types
// type TemplaData struct {
// 	StringMap map[string]string
// 	IntMap    map[string]int
// 	FloatMap  map[string]float32
// 	Data      map[string]interface{}
// 	CSRFToken string
// 	Message   string
// 	Warning   string
// 	Error     string
// }

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
func (m *Repository) home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor login")
	RenderTemplate(w, "advisor_login.page.gohtml", &models.TemplateData)
}

//advisor home page - method with access to repository
func (m *Repository) advisor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor home page")
	RenderTemplate(w, "advisor_home.gohtml", &models.TemplateData)
}

//patient home page - method with access to repository
func (m *Repository) patient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("patient home page")
	RenderTemplate(w, "patient_home.gohtml", &models.TemplateData)
}

func RenderTemplate(w http.ResponseWriter, gohtml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return
	}
}
