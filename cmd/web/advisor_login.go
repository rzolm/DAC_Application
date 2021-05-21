package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var username string = "root"
var password string = ""
var hostname string = "127.0.0.1:3306"
var dbname string = "dac_01"

var agentUsername string    //CFenwick
var agentPassword string    //Password12345
var agentVerificationId int //to verify username and password match

type Agent struct {
	agentId       int
	agentUsername string
	agentPassword string
	login         string
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "username:password@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
	defer db.Close()
	return db
}

func verifyUsername(db *sql.DB, username string) string {

	agent := Agent{}
	//username := "CFenwick"
	rows, err := db.Query("select agentUsername from agentLogin where agentUsername = ?", username)
	if err != nil {
		fmt.Println("error retrieving username")
	}

	for rows.Next() {
		err = rows.Scan(&agent.agentUsername)
	}

	return agent.agentUsername
}

func verifyPassword(db *sql.DB, password string) string {

	agent := Agent{}
	//username := "CFenwick"
	rows, err := db.Query("select agentPassword from agentLogin where agentPassword = ?", password)
	if err != nil {
		fmt.Println("error retrieving password")
	}

	for rows.Next() {
		err = rows.Scan(&agent.agentPassword)
		if err != nil {
			fmt.Println("error retrieving records")
		}
	}

	return agent.agentPassword
}

func verifyCredentials(db *sql.DB, username string, password string) int {

	agent := Agent{}
	//username := "CFenwick"
	rows, err := db.Query("select agentId from agentLogin where agentUsername = ? AND agentPassword = ?", username, password)
	if err != nil {
		fmt.Println("error retrieving credentials")
	}

	for rows.Next() {
		err = rows.Scan(&agent.agentId)
		if err != nil {
			fmt.Println("error retrieving records")
		}
	}

	return agent.agentId
}

func stampLogin(db *sql.DB, agentId int) {
	agent := Agent{}
	fmt.Println("agent id:", agentId)
	update := "2021-04-08 16:31:00"
	_, err := db.Exec("update agentLogin set login = ? where agentId = ?", update, agentId)
	if err != nil {
		fmt.Println("error updating login")
	}

	rows, err := db.Query("select * from agentLogin where agentId = ?", agentId)

	for rows.Next() {
		err := rows.Scan(&agent.agentId, &agent.agentUsername, &agent.agentPassword, &agent.login)
		if err != nil {
			fmt.Println("error retrieving records")
		}
	}

	defer db.Close()
}

func main() {
	fmt.Println("agent login")
	//connectToDatabase()

	fmt.Println("enter your username")
	fmt.Scanln(&agentUsername)

	fmt.Println("enter your password")
	fmt.Scanln(&agentPassword)

	db, err := sql.Open("mysql", "root:@/dac_01") //THIS IS A PROBLEM - INSECURE!!!!
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}

	agentUsername = verifyUsername(db, agentUsername)
	if agentUsername == "" {
		fmt.Println("user does not exist")
	}
	agentPassword = verifyPassword(db, agentPassword)

	fmt.Println("agent username is:", agentUsername)
	agentVerificationId = verifyCredentials(db, agentUsername, agentPassword)
	fmt.Println("agent verified:", agentVerificationId)
	if agentVerificationId == 0 {
		fmt.Println("login unsuccessful")
	} else if agentVerificationId > 0 {
		fmt.Println("login successful")
	}
	stampLogin(db, agentVerificationId)
}
