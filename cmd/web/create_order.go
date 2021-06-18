package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	OrderId              int
	PatientId            int
	AdvisorId            int
	TimePlaced           time.Time
	DeliveryDate         time.Time
	DeliveryCategory     int
	CuttingCompleted     int
	PackedId             int
	PrescriptionReceived int
	DispatchedTime       time.Time
}

var db *sql.DB
var username string = "root"
var password string = "TsP3#86b29"
var hostname string = "127.0.0.1:3306"
var dbname string = "dac_01"

func main() {
	fmt.Println("create order")

	db, err := sql.Open("mysql", "root:TsP3#86b29@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
	CreateOrder(db)
}

func CreateOrder(db *sql.DB) {
	//query to write initial order details
	create :=
		`insert into orders 
	(patientId, agentId, timePlaced, 
	deliveryCategory, cuttingCompleted, packedId, prescriptionReceived, 
	dispatchedTime) 
	values ((?), (?), (?), (?), (?), (?), (?), (?))`

	location, _ := time.LoadLocation("Europe/London")
	t := time.Now().In(location)

	patientId := 10
	agentId := 2
	timePlaced := t
	deliveryCategory := 1
	cuttingCompleted := 0
	packedId := 0
	prescriptionReceived := 0
	dispatchedTime := t //uses current time as a temp placeholder

	_, err := db.Exec(create, patientId, agentId, timePlaced, deliveryCategory,
		cuttingCompleted, packedId, prescriptionReceived, dispatchedTime)
	if err != nil {
		fmt.Print("error writing to database, please check the query")
	}
	fmt.Println(t)
}
