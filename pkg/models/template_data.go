package models

import "github.com/rzolm/DAC_Application/internal/forms"

//TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string //crossed site request forgery token
	Notification    string
	Warning         string
	Error           string
	Form            *forms.Form
	Flash           string
	IsAuthenticated int
}
