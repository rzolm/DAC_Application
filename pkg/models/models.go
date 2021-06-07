package models

import "time"

//make the models singular so that slices that hold multiple data
//can be named as plurals

//login holds login data
type AdvisorLogin struct {
	Username string
	Password string
}

type Advisor struct {
	Name       string
	Username   string
	Password   string
	IsLoggedIn bool
	LastLogin  time.Time
}

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

type Product struct {
}

type Order struct {
}
