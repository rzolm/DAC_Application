package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

//pointer to a template
//tutorial - container into which templates get parsed and are held
var advisorTemplate *template.Template

type Advisor struct {
	AdvisorFirstName  string
	AdvisorLastName   string
	AdvisorActive     bool
	AdvisorLastLogin  string
	AdvisorLastLogout string
}

type AdvisorSchedule struct {
	Patient          int
	LastDiallAttempt string
	LastContact      string
	LastOrder        string
	PatientPriority  int
}

type AdvisorActions struct {
	search        string
	advisorStatus int
}

type Patient struct {
	patientRef int
}

//initialise the template
func init() {
	//pointer to a template
	//tutorial - container into which templates get parsed and are held
	advisorTemplate = template.Must(template.ParseFiles("/Users/Magnesium/DAC_Application/templates/advisor_home.page.gohtml"))
}

func displayAdvisor(a01 Advisor) {
	fmt.Println(a01.AdvisorFirstName, a01.AdvisorLastName)
	err := advisorTemplate.Execute(os.Stdout, a01) //need to send multiple structs to the template
	if err != nil {
		log.Fatalln(err)
	}
}

//function to search for a patient
func searchPatient(patientFirstName string, patientLastName string) {
	patientName := patientFirstName + patientLastName

	fmt.Println(patientName)
}

func main() {
	fmt.Println("Agent Home Screen")

	advisor01 := Advisor{
		AdvisorFirstName:  "Phil",
		AdvisorLastName:   "Smith",
		AdvisorActive:     false,
		AdvisorLastLogin:  "2021-04-13 09:00:03",
		AdvisorLastLogout: "2021-04-13 17:03:02",
	}
	displayAdvisor(advisor01)
	searchPatient("Maisie", "Green")
}
