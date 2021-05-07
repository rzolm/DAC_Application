package main

import (
	"fmt"
	"html/template"
	"net/http"

	"_github.com/go-sql-driver/mysql/pkg/config"
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

//this sets for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//index page handler
func (m *Repository) home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor login")
	RenderTemplate(w, "advisor_login.gohtml")
}

//advisor home page
func (m *Repository) advisor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor home page")
	RenderTemplate(w, "advisor_home.gohtml")
}

//patient home page
func (m *Repository) patient(e http.ResponseWriter, r *http.Request) {
	fmt.Println("patient home page")
	RenderTemplate(w, "patient_home.gohtml")
}

func renderTemplate(w http.ResponseWriter, gohtml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return
	}
}
