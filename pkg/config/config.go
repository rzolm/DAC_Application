package config

//this is used as an alternative to global variables
//each entry in the struct can be used application wide
//serveing to reduce overhead

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	//allows writing logs to wherever needed
	InfoLog      *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}
