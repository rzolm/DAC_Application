package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rzolm/DAC_Application/internal/forms"
	"github.com/rzolm/DAC_Application/pkg/config"
	"github.com/rzolm/DAC_Application/pkg/models"
	"github.com/rzolm/DAC_Application/pkg/render"
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

//handles the advisor login form
func (m *Repository) PostLogin(w http.ResponseWriter, r http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	//login object
	login := models.AdvisorLogin{
		Username: r.Form.Get("Username"),
		Password: r.Form.Get("Password"),
	}
	form := forms.New(r.PostForm)

	form.Has("Username", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["login"] = login

		render.Template(w, r, "index.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}

//index page handler - method with access to repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor login")
	render.Template(w, "advisor_login.gohtml", &models.TemplateData{})
}

//index page handler - post login
func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor login")
	render.Template(w, "advisor_login.gohtml", &models.TemplateData{})
}

//advisor home page - method with access to repository
func (m *Repository) Advisor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("advisor home page")
	render.Template(w, "advisor_home.gohtml", &models.TemplateData{})
}

//patient home page - method with access to repository
func (m *Repository) Patient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("patient home page")
	render.Template(w, "patient_home.gohtml", &models.TemplateData{})
}

func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "testing, testing, 123..."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	//send data to the template
	render.Template(w, "advisor_login.page.gohtml", &models.TemplateData{
		//StringMap: StringMap,
	})
}

// func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
// 	// parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml)
// 	// err := parsedTemplate.Execute(w, nil)
// 	// if err != nil {
// 	// 	fmt.Println("error parsing templates", err)
// 	// 	return
// 	// }
// 	stringMap: stringMap,
// }
