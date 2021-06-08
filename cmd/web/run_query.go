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

type AgentTest struct {
	agentId       int
	agentUsername string
	agentPassword string
	agentLogin    string
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "username:password@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
	return db
}

func queryPatients(db *sql.DB) {
	rows, err := db.Query("select * from agentLogin")
	if err != nil {
		fmt.Println("error retrieving records")
	}

	var agent = AgentTest{}
	//all the columns in the db must be requested otherwise
	//the results are returned as 0 - boooooo!!
	for rows.Next() {
		err = rows.Scan(&agent.agentId, &agent.agentUsername, &agent.agentPassword, &agent.agentLogin)
		fmt.Println(agent)
	}
	defer db.Close()
}

func main() {

	db, err := sql.Open("mysql", "root:TsP3#86b29@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
	queryPatients(db)

}
