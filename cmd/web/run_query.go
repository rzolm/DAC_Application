package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var username string = "root"
var password string = "TsP3#86b29"
var hostname string = "127.0.0.1:3306"
var dbname string = "dac_01"

type Patients struct {
	// id           int
	// status       int
	// nhs          int64
	// gender       string
	PatientFirstName    string
	PatientLastName     string
	// dob          string
	// address      string
	// city         string
	// county       string
	// postcode     string
	// telephpone01 int64
	// telephone02  int64
	// email        string
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:TsP3#86b29@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
	return db
}

func runQuery(db *sql.DB) {
	rows, err := db.Query("select * from patients where patientID = 8")
	if err != nil {
		fmt.Println("error retrieving records")
	}

	//all the columns in the db must be requested otherwise
	//the results are returned as 0 - boooooo!!
	var patient = Patient{}

		err = rows.Scan(&patient.PatientFirstname, &patientLastName)
		fmt.Println(patient)
	}
	defer db.Close()
}

func main() {
	db, err := sql.Open("mysql", "root:TsP3#86b29@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
	runQuery(db)

}
