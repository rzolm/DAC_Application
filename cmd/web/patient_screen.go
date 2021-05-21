package main

import (
	"fmt"
	"html/template"
)

//pointer to a template
//tutorial - container into which templates get parsed and are held
var patientTemplate *template.Template

type Patient struct {
	PatientNumber       int64
	PatientStatus       string
	PatientNHS          int64
	PatientGender       string
	PatientFirstName    string
	PatientLastName     string
	PatientDOB          string
	PatientAddress      string
	PatientCity         string
	PatientCounty       string
	PatientPostcode     string
	PatientContactNum01 int64
	PatientContactNum02 int64
	PatientEmail        string
}

type PatientHistory struct {
	PreviousOrderId             int
	PreviousOrderItems          int   //slice of product Id's
	PatientProducts             []int //slice of product Id's
	PatientOutstandingComplaint bool
	PreviousNotes               string
}

//initialise the template
//func init() {
//pointer to a template
//tutorial - container into which templates get parsed and are held
//advisorTemplate = template.Must(template.ParseFiles("patient.gohtml"))
//}

func main() {
	//is a database connection required here or go through a dedicated function

	//need to create a slice to pass into the struct
	previousOrder := getPreviousOrder()
	patientProducts := getPatientProducts()

	//temp hardcoded patient data to test function
	patient01 := Patient{
		//patientId: 8
		PatientNumber:       1234567,
		PatientStatus:       "active",
		PatientNHS:          123456789,
		PatientGender:       "F",
		PatientFirstName:    "Maisie",
		PatientLastName:     "Green",
		PatientDOB:          "1995-6-16",
		PatientAddress:      "106 Albert Square",
		PatientCity:         "Oldham",
		PatientCounty:       "Greater Manchester",
		PatientPostcode:     "OL38 2WW",
		PatientContactNum01: 0702372736,
		PatientContactNum02: 0161254723,
		PatientEmail:        "Maisie_Green@yahoo.com",
	}

	patientHistory01 := PatientHistory{
		PreviousOrderId:             192,
		PreviousOrderItems:          previousOrder,
		PatientProducts:             patientProducts,
		PatientOutstandingComplaint: false,
		PreviousNotes:               "Phil Smith: Order placed",
	}

	displayPatient(patient01, patientHistory01)
}

//send data to the template
func sendToTemplate() {

}

//get the patients previous order
func getPreviousOrder() int {
	products := 23

	return products
}

//get items patient commonly requests
func getPatientProducts() []int {
	products := []int{23, 111, 101, 148}

	return products
}

//receive Patient struct p01
//send p01 to the template
func displayPatient(p01 Patient, p02 PatientHistory) {
	fmt.Println(p01)
	fmt.Println(p02)
	//err := patientTemplate.Execute(os.Stdout, p01) //need to send multiple structs to the template
	//if err != nil {
	//log.Fatalln(err)
	//}
}
