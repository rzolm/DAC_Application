package config

//this is used as an alternative to global variables
//each entry in the struct can be used application wide
//this also serves to reduce overhead
import (
	"html/template"
	"log"
)

//AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	//allows writing logs to wherever needed
	InfoLog *log.Logger
}
