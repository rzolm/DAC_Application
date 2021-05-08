package main

import (
	"fmt"
	"html/template"
	"net/http"

	"_github.com/go-sql-driver/mysql/pkg/config"
)

type TemplateData struct {
	StringMap    map[string]string
	IntMap       map[string]int
	FloatMap     map[string]float32
	Data         map[string]interface{}
	CSRFToken    string
	Notification string
	Warning      string
	Error        string
}

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

//this sets for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//index page handler
func (m *Repository) home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor login")
	RenderTemplate(w, "advisor_login.gohtml", &TemplateData{})
}

//advisor home page
func (m *Repository) advisor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor home page")
	RenderTemplate(w, "advisor_home.gohtml", &TemplateData{})
}

//patient home page
func (m *Repository) patient(e http.ResponseWriter, r *http.Request) {
	fmt.Println("patient home page")
	RenderTemplate(w, "patient_home.gohtml", &TemplateData{})
}

func renderTemplate(w http.ResponseWriter, gohtml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return
	}
}
