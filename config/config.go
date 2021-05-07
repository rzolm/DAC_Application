package config

//this is used as an alternative to global variables
//each entry in the struct can be used application wide

import (
	"html/template"
	"log"
)

//AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger //allows writing logs to wherever
}
